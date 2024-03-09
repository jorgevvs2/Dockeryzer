package utils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func safePrintf(output *color.Color, format string, args ...interface{}) {
	_, err := output.Printf(format, args...)
	if err != nil {
		fmt.Println("Failed to call colored Printf()")
		fmt.Println(err)
		os.Exit(0)
	}
}

// SUCCESS

func getSuccessOutput() *color.Color {
	return color.New(color.FgGreen)
}

func SuccessPrintf(format string, args ...interface{}) {
	safePrintf(getSuccessOutput(), format, args...)
}

func SuccessSprintf(format string, args ...interface{}) string {
	return getSuccessOutput().Sprintf(format, args...)
}

// INFO

func getInfoOutput() *color.Color {
	return color.New(color.FgHiBlue)
}

func InfoPrintf(format string, args ...interface{}) {
	safePrintf(getInfoOutput(), format, args...)
}

// ERROR

func getErrorOutput() *color.Color {
	return color.New(color.FgRed)
}

func ErrorPrintf(format string, args ...interface{}) {
	safePrintf(getErrorOutput(), format, args...)
}

func ErrorSprintf(format string, args ...interface{}) string {
	return getErrorOutput().Sprintf(format, args...)
}

// WARNING

func getWarningOutput() *color.Color {
	return color.New(color.FgYellow)
}

func WarningSprintf(format string, args ...interface{}) string {
	return getWarningOutput().Sprintf(format, args...)
}

// BOLD

func getBoldOutput() *color.Color {
	return color.New(color.Bold)
}

func BoldPrintf(format string, args ...interface{}) {
	safePrintf(getBoldOutput(), format, args...)
}
