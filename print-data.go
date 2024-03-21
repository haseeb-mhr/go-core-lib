package gocorelib

import (
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
)

func PrintString(str string) {
	fmt.Printf("Its is String value => %v", str)
}

func GetFile(fileUrl string) (https.FileWrapper, error){
	return https.GetFile(fileUrl)
}