package logging

import (
	"io"
	"log"
)

var (
	// Trace - trace logging interface
	Trace *log.Logger
	// Info - info logging interface
	Info *log.Logger
	// Warning - warning logging interface
	Warning *log.Logger
	// Error - error logging interface
	Error *log.Logger
)

// Init - initialize a logging controler
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
