package main

import (
	"fmt"
	"os"
	"pm/files"
	pm "pm/manage"
)

func main() {

	path := files.GetPMPath()
	_, err := os.Open(path)
	if os.IsNotExist(err) {
		fmt.Println("This appears to be your first time using pm")
		fmt.Println("Please choose a device to act as your hardward key")
	}

	key, err := pm.GetKey()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SHA-256 hash of serial number: %x\n", key)

	pm.RunPmShell(key)
}
