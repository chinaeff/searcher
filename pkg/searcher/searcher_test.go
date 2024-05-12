package searcher

import (
	"searcher/pkg/internal/dir"
	"testing"
)

func TestSearchFileByWord_NotFound(t *testing.T) {
	// Инициализация карты файлов с пустым содержимым
	dir.FileDataMap = make(map[string]*dir.FileData)

	// Выполнение поиска слова, которого нет в файлах
	result, err := dir.SearchFile("test")

	// Проверка, что результат равен nil
	if result != nil {
		t.Errorf("ожидался результат nil, получено: %v", result)
	}

	// Проверка, что ошибка не равна nil
	if err == nil {
		t.Error("ожидалась ошибка, получено nil")
	}
}

func TestPreprocessFiles_ErrorReadingFile(t *testing.T) {
	// Попытка обработки несуществующего файла
	err := dir.ProcessingFiles()

	// Проверка, что ошибка не равна nil
	if err == nil {
		t.Error("ожидалась ошибка чтения файла, получено nil")
	}
}
