package pm

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetKey() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	// Run the lsusb command and print its output
	fmt.Println(RunCommand("lsusb"))

	// Read the bus number from the user
	fmt.Print("Enter bus number: ")
	bus, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading bus number: %v\n", err)
		return nil, err
	}
	bus = strings.TrimSpace(bus)

	// Read the device number from the user
	fmt.Print("Enter device number: ")
	dev, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading device number: %v\n", err)
		return nil, err
	}
	dev = strings.TrimSpace(dev)

	fmt.Printf("Device: %s, Bus: %s\n", dev, bus)

	// Construct the device path and run udevadm command
	devicePath := fmt.Sprintf("/dev/bus/usb/%s/%s", bus, dev)
	stdout := RunCommand(fmt.Sprintf("udevadm info --name=%s", devicePath))

	// Filter the output using grep to find the SERIAL
	cmd := exec.Command("grep", "SERIAL")
	cmd.Stdin = strings.NewReader(stdout)
	serialNo, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Hash the serial number and print the hash
	hash := sha256.Sum256(serialNo)

	return hash[:], nil
}
