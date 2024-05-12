package dir

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// Обработка файлов: удаляем знаки препинания, приводим найденные слова к нижнему регистру и вносим данные в мапу
func ProcessingFiles() error {
	FileDataMap = make(map[string]*FileData)
	dir := "./examples"

	re := regexp.MustCompile(`[^\p{L}\s]`)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Пропускаем директории
		if info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error of reading file %s: %v", path, err)
		}

		fileData := &FileData{
			FileName: info.Name(),
			Words:    make(map[string]struct{}),
		}

		processedContent := re.ReplaceAllString(string(content), "")
		processedContent = strings.ToLower(processedContent)
		words := strings.Fields(processedContent)

		for _, word := range words {
			fileData.Words[word] = struct{}{}
		}

		FileDataMap[info.Name()] = fileData

		return nil
	})
	if err != nil {
		return fmt.Errorf("error of processing _files: %v", err)
	}
	return nil
}

// Поиск файла, содержащего слово
func SearchFile(word string) ([]string, error) {
	//Определяем количество горутин в зависимости от количества файлов в директории
	numFiles := len(FileDataMap)
	numWorkers := runtime.NumCPU()
	if numFiles < numWorkers {
		numWorkers = numFiles
	}
	//Создаем два канала для передачи работы горутинам и передачи имен файлов из горутин
	work := make(chan string, numFiles)
	results := make(chan string, numFiles)

	//Функция для поиска в файле
	search := func() {
		for fileName := range work {
			fileData, ok := FileDataMap[fileName]
			if !ok {
				continue
			}
			if _, ok := fileData.Words[word]; ok {
				results <- fileName
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			search()
		}()
	}

	//Передача имен файлов в канал
	for fileName := range FileDataMap {
		work <- fileName
	}
	close(work)

	//Запускаем горутину для имен файлов
	go func() {
		wg.Wait()
		close(results)
	}()

	//Читаем имена файлов из канала и добавляем в слайс
	var foundFiles []string
	for fileName := range results {
		foundFiles = append(foundFiles, fileName)
	}

	if len(foundFiles) == 0 {
		return nil, fmt.Errorf("слово '%s' не найдено в файлах", word)
	}

	return foundFiles, nil
}
