package interfaces

type Logger interface {
	Info(message string)
	Error(message string)
}
