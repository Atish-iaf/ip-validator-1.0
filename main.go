package main

import (
	"fmt"
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

func main() {

	var ip string
	fmt.Scanf("%s", &ip)

	checkIPv4, errIPv4 := validateIPv4(ip)
	checkIPv6, errIPv6 := validateIPv6(ip)

	if errIPv4 != nil {
		fmt.Println(errIPv4)
	} else if checkIPv4 {
		fmt.Println("validIPv4")
	} else if errIPv6 != nil {
		fmt.Println(errIPv6)
	} else if checkIPv6 {
		fmt.Println("validIPv6")
	} else {
		fmt.Println("invalidIP")
	}
}
