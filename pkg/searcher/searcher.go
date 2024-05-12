package searcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"searcher/pkg/internal/dir"
	"strings"
)

// хэндлер для запуска поиска слова в файлах
func Search(w http.ResponseWriter, r *http.Request) {
	//Получаем слово из адресной строки и приводим к нижнему регистру (в таком виде они хранятся в мапе)
	word := strings.ToLower(r.URL.Path[len("/files/search/"):])
	if word == "" {
		http.Error(w, "Слово для поиска не указано в URL", http.StatusBadRequest)
		return
	}

	result, err := dir.SearchFile(word)
	if err != nil {
		fmt.Errorf("слово '%s' не найдено в файлах", word)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ProcessFile() error {
	dir.ProcessingFiles()
	return nil
}
