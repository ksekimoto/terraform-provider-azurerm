package azs

import (
	"fmt"
	"log"
)

const DEBUG = true
const ERROR = true
const TRACE = false

func Printf(format string, v ...interface{}) {
	if DEBUG {
		log.Printf(fmt.Sprintf(format, v...))
	}
}

func Debug(format string, v ...interface{}) {
	if DEBUG {
		log.Printf(fmt.Sprintf(format, v...))
	}
}

func Trace(format string, v ...interface{}) {
	if TRACE {
		log.Printf(fmt.Sprintf(format, v...))
	}
}

func Error(format string, v ...interface{}) {
	if ERROR {
		log.Printf(fmt.Sprintf(format, v...))
	}
}
