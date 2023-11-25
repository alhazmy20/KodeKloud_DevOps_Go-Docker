package utils

import (
	"log"
	"runtime"
)

func HandleError(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Fatalf("Error in file %s at line %d: %v", file, line, err)
}
