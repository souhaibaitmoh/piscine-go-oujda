package main

import (
	"fmt"
	"os"
)

func main() {
	v := os.Args
	for _, c := range v {
		if c == "01" || c == "galaxy" || c == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
