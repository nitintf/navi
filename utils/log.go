package utils

import "fmt"

func LogSuccess(message string) {
	fmt.Printf("\033[1;32m%s\033[0m\n", message)
}

func LogError(message string) {
	fmt.Printf("\033[1;31m%s\033[0m\n", message)
}

func LogInfo(message string) {
	fmt.Printf("\n\033[1;36m%s\033[0m%s\n", ">>  ", message)
}

func LogExplanation(message string) {
	fmt.Printf("\n\033[1;33m%s\033[0m\n\n%s\n", "Explanation >>  ", message)
}
