package gocorelib

import (
	"fmt"
	"strings"

	"github.com/apimatic/go-core-runtime/https"
)

func PrintString(str string) {
	fmt.Printf("Its is String value => %v", str)
}

func GetFile(fileUrl string) (https.FileWrapper, error){
	return https.GetFile(fileUrl)
}


func MergePath(left, right string) string {
	if right == "" {
		return left
	}
	if left[len(left)-1] == '/' && right[0] == '/' {
		return left + strings.Replace(right, "/", "", -1)
	} else if left[len(left)-1] == '/' || right[0] == '/' {
		return left + right
	} else {
		return left + "/" + right
	}
}