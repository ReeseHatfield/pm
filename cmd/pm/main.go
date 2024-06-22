package main

import (
	"fmt"
	"pm/pm"
)

func main() {

	key, err := pm.GetKey()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("SHA-256 hash of serial number: %x\n", key)
}
