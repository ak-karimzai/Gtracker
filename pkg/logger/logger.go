package logger

type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(args ...interface{})
}
