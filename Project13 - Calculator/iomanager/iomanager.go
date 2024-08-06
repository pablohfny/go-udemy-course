package iomanager

type IOManager interface {
	ReadData() ([]string, error)
	WriteData(data any) error
}
