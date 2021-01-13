package main

import (
	"fmt"

	au "github.com/logrusorgru/aurora"
)

// Info displays an info message
func Info(args ...interface{}) {
	args = append([]interface{}{au.Cyan("INFO:")}, args...)
	fmt.Println(args...)
}

// Warning displays a warning message
func Warning(args ...interface{}) {
	args = append([]interface{}{au.Yellow("WARNING:")}, args...)
	fmt.Println(args...)
}

// Error displays an error message
func Error(args ...interface{}) {
	args = append([]interface{}{au.Red("ERROR:")}, args...)
	fmt.Println(args...)
}
