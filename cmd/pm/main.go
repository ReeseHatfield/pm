package main

import (
	"fmt"
	"pm/files"
	"pm/pm"
)

func main() {

	key, err := pm.GetKey()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SHA-256 hash of serial number: %x\n", key)

	pmDotDat := "~/.pm/pm.dat"
	if !files.FileExists(pmDotDat) {
		fmt.Println(pmDotDat, "was not found")
		// walk thru creation process with encryption
	}

}
