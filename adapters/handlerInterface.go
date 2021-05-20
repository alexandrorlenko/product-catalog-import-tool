package adapters

type HandlerInterface interface {
	Init(t FileType)
	Parse(file string) []map[string]interface{}
	Write(file string, data [][]string)
}
