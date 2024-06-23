package pm

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"pm/data"
	"pm/files"
	"strings"
)

// RunCommand runs a shell command and returns its output as a string
func RunCommand(cmdString string) string {
	cmd := exec.Command("sh", "-c", cmdString)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing command '%s': %v\n", cmdString, err)
		os.Exit(1)
	}
	return string(stdout)
}

func Usage() {

	usage := `
1. get [website thingy] Gets credentials for [website thing]
2. add [website thingy] Add credentials for [website thingy]
3. del [website thingy] Delete credentials for [website thingy]
4. upd [website thingy] Update credentials for [website thingy]
	`
	fmt.Println(usage)
}

func RunPmShell(key []byte) {

	dict, err := files.LoadPmFile(key)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Println("Could not load pm, Is your key wrong?")
		return
	}

	fmt.Println("Welcome to the pm shell")
	Usage()

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print(">> ")
		cmd, arg := GetCommand(*reader)

		switch cmd {
		case "help", "-help", "--help":
			Usage()
		case "get":
			dict = get(arg, dict)
		case "add":
			dict = add(arg, dict)
		case "del":
			dict = del(arg, dict)
		case "upd":
			dict = upd(arg, dict)
		case "print-all":
			// debug, delete in prod
			fmt.Println(dict)
		case "exit", "quit", "q":
			return
		default:
			fmt.Println("Unknown command:", cmd, arg)
			fmt.Println("Try -help for usage")
		}

		files.SavePmFile(key, dict)
	}
}

func get(service string, dict data.PMDictionary) data.PMDictionary {
	return dict
}

func upd(service string, dict data.PMDictionary) data.PMDictionary {
	return dict
}

func add(service string, dict data.PMDictionary) data.PMDictionary {
	return dict
}

func del(service string, dict data.PMDictionary) data.PMDictionary {
	return dict
}

func GetCommand(reader bufio.Reader) (cmd, arg string) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")

	// there is probably a better way of doing this
	cmd = parts[0]
	if len(parts) > 1 {
		arg = parts[1]
	} else {
		arg = ""
	}

	return cmd, arg

}
