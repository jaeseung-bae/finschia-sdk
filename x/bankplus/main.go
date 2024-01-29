package main

import (
	"fmt"

	"github.com/Finschia/finschia-sdk/x/collection-token/class"
)

func main() {
	err := class.ValidateID("Aa")
	if err != nil {
		fmt.Println(err)
	}
}
