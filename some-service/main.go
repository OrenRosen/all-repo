package main

import (
	"fmt"
	
	"github.com/OrenRosen/shared/mmsql"
)

func main() {
	fmt.Println("------------------ Hello, World!")
	
	mmsql.NewClient()
}
