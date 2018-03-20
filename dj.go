// The dj command line application allows you to run django manage.py commands
// from any subfolder of the django project.

package main

import (
    "fmt"
	"os"
	"os/exec"
	"path/filepath"
	"log"
)

// isExistingFile checks if the file with the given filename exists.
func isExistingFile(filename string) bool {
    _, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

// findManagePy finds the manage.py file for the given directory.
// Walks from the given dir up to the root directory. If a manage.py
// file is found managePyFilename will contain the path to the
// manage.py file and ok will be true. If a manage.py file is not
// found managePyFilename will contain the empty string and ok will
// be false.
func findManagePy(dir string) (managePyFilename string, ok bool) {
	for dir != "/" {
		filename := filepath.Join(dir, "manage.py")

		if isExistingFile(filename) {
			return filename, true
		} else {
			dir = filepath.Dir(dir)
		}
	}
	return "", false
}

// Creates the arguments that should be passed to the python
// executable.
func createArgs(managePyFilename string) []string{
	args := make([]string, len(os.Args))
	copy(args, os.Args)
	args[0] = managePyFilename

	return args
}

// Executes manage.py with the arguments that were passed
// in to this program on the command line.
func executeManagePy(managePyFilename string) {
	args := createArgs(managePyFilename)
	cmd := exec.Command("python", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		log.Fatal("Unable to start python")
	}
}

// If manage.py is found execute it with the arguments passed
// in to this program on the command line.
func main() {
    dir, err := os.Getwd()
	if err != nil {
		os.Exit(128)
	}
	managePyFilename, ok := findManagePy(dir)
	if ok {
		fmt.Println(managePyFilename)
		executeManagePy(managePyFilename)
	} else {
		fmt.Fprintln(os.Stderr, "Not in a Django project.")
		os.Exit(1)
	}
}
