package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"sync"
)

// example of using

//import "ginTemplate/pkg/logging"
//
//func main() {
//	logging.DefaultLogger.Info().Msg("Hello")
//}

type Logger struct {
	*zerolog.Logger
}

var DefaultLogger Logger
var once sync.Once

func CustomConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006/01/02 - 15:04:05",
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s", i)
		},
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("%s=", i)
		},
		FormatFieldValue: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
		FormatErrFieldName: func(i interface{}) string {
			return fmt.Sprintf("%s=", i)
		},
		FormatErrFieldValue: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
	}
}

func initLogger() Logger {
	var logger Logger
	once.Do(func() {
		output := CustomConsoleWriter()
		zerologInst := zerolog.New(output).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Logger()
		logger = Logger{&zerologInst}
	})
	return logger
}

func init() {
	DefaultLogger = initLogger()
}
