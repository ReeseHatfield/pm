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
    Command		|	 		Description
====================================================
get service		|	 Gets credentials for service
add service		|	 Add credentials for service
del service		|	 Delete credentials for service
upd service		|	 Update credentials for service
ls 			|	 List all services
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

		fmt.Printf("arg: %v\n", arg)

		switch cmd {
		case "help", "-help", "--help":
			Usage()
		case "ls":
			ls(dict)
		case "get":
			get(arg, dict)
		case "add":
			dict = add(arg, dict, reader)
		case "del":
			dict = del(arg, dict, reader)
		case "upd":
			dict = upd(arg, dict, reader)
		case "find":
			//TODO fuzzy find thru services
		case "exit", "quit", "q":
			return
		default:
			fmt.Println("Unknown command:", cmd, arg)
			fmt.Println("Try -help for usage")
		}

		files.SavePmFile(key, dict)
	}
}

func ls(dict data.PMDictionary) {
	for k := range dict {
		fmt.Println(k)
	}
}

func get(service string, dict data.PMDictionary) {
	cred, ok := dict[service]

	if !ok {
		fmt.Println("Could not find service: ", service)
		return
	}

	fmt.Println(cred.String())
}

func upd(service string, dict data.PMDictionary, reader *bufio.Reader) data.PMDictionary {
	_, ok := dict[service]
	if !ok {
		fmt.Println("Err, could not find serbice: ", service)
		return dict
	}

	fmt.Print("What would you like to change? (username or password)? ")
	changeMe, err := reader.ReadString('\n')
	changeMe = strings.TrimSpace(changeMe)

	if err != nil {
		fmt.Println("Invalid option!")
		return dict
	}

	if !strings.EqualFold(changeMe, "password") && !strings.EqualFold(changeMe, "username") {
		fmt.Println("Invalid option!")
		return dict
	}

	fmt.Printf("What would you like to change the %s to? ", changeMe)
	updatedField, err := reader.ReadString('\n')
	updatedField = strings.TrimSpace(updatedField)

	if err != nil {
		fmt.Println("Err, could not update field")
		return dict
	}

	if strings.EqualFold(changeMe, "password") {
		dict[service] = data.Credentials{
			Username: dict[service].Username,
			Password: updatedField,
		}
	} else {
		dict[service] = data.Credentials{
			Username: updatedField,
			Password: dict[service].Password,
		}
	}

	return dict

}

func add(service string, dict data.PMDictionary, reader *bufio.Reader) data.PMDictionary {

	if service == "" {
		fmt.Println("Error: service must have name")
		return dict
	}

	fmt.Print("Username/Email? ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return dict
	}
	username = strings.TrimSpace(username)

	fmt.Print("Password?")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return dict
	}
	password = strings.TrimSpace(password)

	dict[service] = data.Credentials{
		Username: username,
		Password: password,
	}

	return dict

}

func del(service string, dict data.PMDictionary, reader *bufio.Reader) data.PMDictionary {
	_, ok := dict[service]
	if !ok {
		fmt.Println("Err, could not find service: ", service)
		return dict
	}

	fmt.Printf("Are you sure you want to permanently delete credentials for %s [y/n] ", service)
	yn, err := reader.ReadString('\n')
	yn = strings.TrimSpace(yn)

	if err != nil {
		fmt.Println("Err could not delete credentials")
		return dict
	}

	if yn == "n" {
		return dict
	}

	fmt.Printf("Deleting service %s ...\n", service)
	delete(dict, service)

	return dict
}

func GetCommand(reader bufio.Reader) (cmd string, arg string) {

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")

	cmd = parts[0]
	remaining := parts[1:]

	arg = ""
	for _, s := range remaining {
		arg += (s + "-")
	}

	arg = strings.TrimSuffix(arg, "-")

	return cmd, arg

}
