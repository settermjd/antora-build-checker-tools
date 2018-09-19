package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	fmt.Println(containsRSTFormatting(getFileContents(getFilename())))
}

// Checks if a string contains reStructuredText formatting.
// Currently this is limited to headers
func containsRSTFormatting(dat []byte) (string, error) {
	if len(string(dat)) == 0 {
		return "", errors.New("No data provided")
	}

	var re = regexp.MustCompile(`(?m)^([\^\~\|=-]{5,})$`)
	if len(re.FindStringIndex(string(dat))) > 0 {
		return fmt.Sprint("Contains reStructuredText formatting"), nil
	}

	return "No reStructuredText formatting detected", nil
}

// Retrieve the contents of a file
func getFileContents(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return dat
}

// Get the file name supplied to the script/code/app
func getFilename() string {
	scanner := bufio.NewScanner(os.Stdin)
	filename := ""
	for scanner.Scan() {
		filename = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	return filename
}
