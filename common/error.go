package common

import (
	"fmt"
	"log"
)

var (
	RecordNotFound = fmt.Errorf("record not found")
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
