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

	// scanner := bufio.NewScanner(os.Stdin)

	// pmDotDat := files.GetPMPath()
	// if !files.FileExists(pmDotDat) {
	// 	fmt.Println(pmDotDat, "was not found")
	// 	// walk thru creation process with encryption
	// 	fmt.Println("Would you like to set up pm?")

	// 	var line string
	// 	for scanner.Scan() {

	// 		line = scanner.Text()

	// 		if utils.IsValidYesNoString(line) {
	// 			break
	// 		}
	// 	}

	// 	if rune(line[0]) == 'y' {
	// 		// create and encrypt file
	// 		fmt.Println("generating file")
	// 	} else {
	// 		os.Exit(1)
	// 	}
	// }

	pm.RunPmShell(key)

	// // status := files.SavePmFile(key, mockData)

	// // fmt.Printf("status: %v\n", status)

	// dict, err := files.LoadPmFile(key)

	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Println(dict)

}
