package utils

import "fmt"

func PrintLog(functionPath string, message string) {
	fmt.Printf("(%s):: %s\n", functionPath, message)
}
