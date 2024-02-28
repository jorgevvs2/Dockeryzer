package utils

import "github.com/fatih/color"

func GetSuccessOutput() *color.Color {
	return color.New(color.FgGreen)
}

func GetInfoOutput() *color.Color {
	return color.New(color.FgHiBlue)
}

func GetErrorOutput() *color.Color {
	return color.New(color.FgRed)
}

func GetWarningOutput() *color.Color {
	return color.New(color.FgYellow)
}

func GetBoldOutput() *color.Color {
	return color.New(color.Bold)
}
