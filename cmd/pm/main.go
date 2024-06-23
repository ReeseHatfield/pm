package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"pm/files"
	"pm/pm"
	"pm/utils"
)

func main() {

	key, err := pm.GetKey()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SHA-256 hash of serial number: %x\n", key)

	scanner := bufio.NewScanner(os.Stdin)

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	pmDotDat := usr.HomeDir + "/.pm/pm.dat"
	if !files.FileExists(pmDotDat) {
		fmt.Println(pmDotDat, "was not found")
		// walk thru creation process with encryption
		fmt.Println("Would you like to set up pm?")

		var line string
		for scanner.Scan() {

			line = scanner.Text()

			if utils.IsValidYesNoString(line) {
				break
			}
		}

		if rune(line[0]) == 'y' {
			// create and encrypt file
			fmt.Println("generating file")
		} else {
			os.Exit(1)
		}
	} else {
		fmt.Println("pm.dat WAS found")
	}

}
