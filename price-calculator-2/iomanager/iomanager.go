package iomanager

type IOManager interface {
	WriteResult(data interface{}) error
	ReadLines() ([]string, error)
}
