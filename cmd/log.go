package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/mgutz/ansi"
	log "github.com/sirupsen/logrus"
)

// Formatter is used for logrus.
type Formatter struct {
}

// Format is used to custom format the logs.
func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	serialized := fmt.Sprintf(ansi.Color("["+entry.Time.Format("2006-01-02 15:04:05")+"] ", "yellow+b"))
	serialized += fmt.Sprintf(ansi.Color("["+strPad(strings.ToUpper(entry.Level.String()), 7, " ")+"] ", getColor(entry.Level)))
	serialized += entry.Message
	serialized += "\n"

	return []byte(serialized), nil
}

func strPad(input string, padLength int, padString string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}
	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))
	length := (float64(padLength - inputLength)) / float64(2)
	repeat = math.Ceil(length / float64(padStringLength))
	output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]

	return output
}

func getColor(level log.Level) string {
	switch level {
	case log.TraceLevel:
		return "green+b"
	case log.DebugLevel:
		return "green+b"
	case log.InfoLevel:
		return "blue+b"
	case log.WarnLevel:
		return "red+b"
	case log.ErrorLevel:
		return "red+b"
	case log.FatalLevel:
		return "red+b"
	case log.PanicLevel:
		return "red+b"
	}

	return "white"
}
