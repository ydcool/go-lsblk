package main

import (
	"fmt"
	"strings"
)

func main() {
	// Test the parsing logic directly
	line := `NAME="nvme1n1" PATH="/dev/nvme1n1" TYPE="disk" MODEL="SAMSUNG MZQL23T8HCLS-00B7C"`
	result := parseKeyValuePairs(line)
	fmt.Println("Test parseKeyValuePairs:")
	fmt.Printf("MODEL: %s\n", result["MODEL"])
	fmt.Printf("Expected: SAMSUNG MZQL23T8HCLS-00B7C\n")
	fmt.Printf("Test passed: %v\n", result["MODEL"] == "SAMSUNG MZQL23T8HCLS-00B7C")
	fmt.Println()

	// Test with multiple spaces in value
	line2 := `NAME="sda" MODEL="Western Digital HDD 1TB" TYPE="disk"`
	result2 := parseKeyValuePairs(line2)
	fmt.Println("Test with multiple spaces:")
	fmt.Printf("MODEL: %s\n", result2["MODEL"])
	fmt.Printf("Expected: Western Digital HDD 1TB\n")
	fmt.Printf("Test passed: %v\n", result2["MODEL"] == "Western Digital HDD 1TB")
	fmt.Println()

	// Test with unquoted values
	line3 := `NAME=sda TYPE=disk SIZE=1000G`
	result3 := parseKeyValuePairs(line3)
	fmt.Println("Test with unquoted values:")
	fmt.Printf("NAME: %s\n", result3["NAME"])
	fmt.Printf("TYPE: %s\n", result3["TYPE"])
	fmt.Printf("SIZE: %s\n", result3["SIZE"])
	fmt.Printf("Test passed: %v\n", result3["NAME"] == "sda" && result3["TYPE"] == "disk" && result3["SIZE"] == "1000G")
}

// Copy of the parseKeyValuePairs function from lsblk_linux.go
func parseKeyValuePairs(line string) map[string]string {
	result := make(map[string]string)
	line = strings.TrimSpace(line)
	if line == "" {
		return result
	}

	var key, value string
	var inQuotes bool
	var current strings.Builder
	var state int // 0: looking for key, 1: in key, 2: looking for =, 3: looking for value, 4: in value

	for _, r := range line {
		switch state {
		case 0:
			if r != ' ' {
				current.WriteRune(r)
				state = 1
			}
		case 1:
			if r == '=' {
				key = current.String()
				current.Reset()
				state = 3
			} else {
				current.WriteRune(r)
			}
		case 3:
			if r == '"' {
				inQuotes = true
				state = 4
			} else if r != ' ' {
				current.WriteRune(r)
				state = 4
			}
		case 4:
			if inQuotes {
				if r == '"' {
					inQuotes = false
					value = current.String()
					current.Reset()
					result[key] = value
					state = 0
				} else {
					current.WriteRune(r)
				}
			} else {
				if r == ' ' {
					value = current.String()
					current.Reset()
					result[key] = value
					state = 0
				} else {
					current.WriteRune(r)
				}
			}
		}
	}

	// Handle last key-value pair if line ends without space
	if key != "" && (state == 4 || current.Len() > 0) {
		if current.Len() > 0 {
			value = current.String()
		}
		result[key] = value
	}

	return result
}
