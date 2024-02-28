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
