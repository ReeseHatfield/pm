package pm

import (
	"fmt"
	"os"
	"os/exec"
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
