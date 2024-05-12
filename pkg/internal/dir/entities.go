package dir

// Структура для хранения данных о файлах
type FileData struct {
	FileName string
	Words    map[string]struct{}
}

var FileDataMap map[string]*FileData
