package main

import "fmt"

type LogLevel int

const (
	LogLevel_Error LogLevel = iota
	LogLevel_Warning
	LogLevel_Info
	LogLevel_Debug
	LogLevel_Trace
)

var LogLevelName = map[LogLevel]string{
	LogLevel_Error:   "Error",
	LogLevel_Warning: "Warning",
	LogLevel_Info:    "Info",
	LogLevel_Debug:   "Debug",
	LogLevel_Trace:   "Trace",
}

var GlobalLogLevel = LogLevel_Info

func Logf(level LogLevel, format string, a ...any) (n int, e error) {
	if GlobalLogLevel <= level {
		return fmt.Printf(fmt.Sprintf("[%s] %s", LogLevelName[level], format), a...)
	}

	return 0, nil
}

func Logln(level LogLevel, a ...any) (n int, e error) {
	if GlobalLogLevel <= level {
		return fmt.Println(fmt.Sprintf("[%s] %v", LogLevelName[level], fmt.Sprint(a...)))
	}

	return 0, nil
}

func Log(level LogLevel, a ...any) (n int, e error) {
	if GlobalLogLevel <= level {
		return fmt.Print(fmt.Sprintf("[%s] %v", LogLevelName[level], fmt.Sprint(a...)))
	}

	return 0, nil
}
