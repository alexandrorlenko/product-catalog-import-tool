package fileHandler

type FileHandlerInterface interface {
	InitReader(path string)
	ReadLine() ([]string, error)
}
