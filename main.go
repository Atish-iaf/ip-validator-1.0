package main

import (
	//"bufio"
	//"log"
	//"os"
	"regexp"
)

// validateIPv4 function to check a given string is valid ipv4 or not.
func validateIPv4(line string) (bool, error) {
	ipv4, err := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`)
	if err != nil {
		return false, err
	}
	if ipv4.MatchString(line) {
		return true, err
	} else {
		return false, err
	}
}

// validateIPv6 function to check a given string is valid ipv6 or not.
func validateIPv6(line string) (bool, error) {
	ipv6, err := regexp.Compile(`^((([0-9a-fA-F]){1,4})\:){7}([0-9a-fA-F]){1,4}$`)
	if err != nil {
		return false, err
	}
	if ipv6.MatchString(line) {
		return true, err
	} else {
		return false, err
	}
}

/*
func main() {
	// opening of a log file
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	log.Println("############### IP VALIDATION LINE BY LINE ###############")

	// opening of file to be read
	fileToRead, ferr := os.Open("ipTestcases.txt")
	if ferr != nil {
		log.Fatal(ferr)
	}
	defer fileToRead.Close()

	scanner := bufio.NewScanner(fileToRead)
	var lineNum, countIPv4, countIPv6 int = 0, 0, 0

	for scanner.Scan() {
		// reading file line by line
		line := scanner.Text()
		lineNum++

		if validateIPv4(line) {
			countIPv4++
			log.Printf("Line : %d - valid IPv4\n", lineNum)
		} else if validateIPv6(line) {
			countIPv6++
			log.Printf("Line : %d - valid IPv6\n", lineNum)
		} else {
			log.Printf("Line : %d - invalid IP\n", lineNum)
		}
	}

	log.Println("------------ RESULTS ------------")
	log.Printf("Number of lines checked : %d\n", lineNum)
	log.Printf("Number of valid IPv4 : %d\n", countIPv4)
	log.Printf("Number of valid IPv6 : %d\n", countIPv6)
	log.Printf("Number of invalid IP : %d\n", lineNum-countIPv4-countIPv6)

}*/
