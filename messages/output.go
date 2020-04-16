package messages

import "github.com/fatih/color"

// ErrorMessage takes in a string and prints red output to the terminal
func ErrorMessage(err string) {
	// custom print function
	errorMessage := color.New(color.FgRed).PrintfFunc()
	// error message
	errorMessage("Error: %s", err)
}

// SuccessMessage takes in a string and prints green output to the terminal
func SuccessMessage(success string) {
	// custom print ffunction
	successMessage := color.New(color.FgGreen).PrintFunc()
	// success message
	successMessage(success)
}

// WarningMessage takes in a string and print yellow output to the terminal
func WarningMessage(warning string) {
	// custom print function
	warningMessage := color.New(color.FgHiYellow).PrintFunc()
	// warning message
	warningMessage(warning)
}
