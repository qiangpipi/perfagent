package logs

import (
	. "fmt"
)

var D bool

func Debug(debug ...interface{}) {
	if D {
		Println(debug)
	}
}

func Error(err ...interface{}) {
	Println(err)
}

func Info(info ...interface{}) {
	Println(info)
}
