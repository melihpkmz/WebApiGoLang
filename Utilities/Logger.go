package Utilities

import (
	"os"
	"time"
)

type LogType int

const(
	UNDEFINED LogType = iota + 1
	WARNING
	ERROR
	INFO
)

func (e LogType) String() string{
	return [...]string{"","UNDEFINED","WARNING","ERROR","INFO"}[e]
}


func Logger(logType string, params ...string) {
	logFileName := time.Now().Format(time.RFC3339) + ".txt"
	description := ""
	for _, val := range params{
		description = description + " " + val + " "
	}
	tempString := time.Now().Format(time.RFC3339) +"  "+ logType+ "  " + description
	f, err := os.Create(logFileName)
	if err != nil{
		panic(err)
	}
	_, err = f.WriteString(tempString)

	if err != nil{
		panic(err)
	}


}