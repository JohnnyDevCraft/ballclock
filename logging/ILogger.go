package logging

//ILogger Interface
type ILogger interface {
	Info(data string)
	Warn(data string)
	Err(data string)
}
