package main

import "testing"

type addTest struct {
	ip             string
	expectedOutput bool
}

var testCasesIPv4 = []addTest{
	{"128.0.0.1", true},
	{"125.512.100.abc", false},
	{"256.256.256.256", false},
	{"192.168.01.1", false},
	{"000.12.234.23.23", false},
	{"192.168.1.0", true},
	{"2001:0db8:85a3:0:0:8A2E:0370:7334", false},
}

func TestValidateIPv4(t *testing.T) {

	for _, test := range testCasesIPv4 {
		output, err := validateIPv4(test.ip)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		} else if output != test.expectedOutput {
			t.Errorf("expected %v , got %v\n", test.expectedOutput, output)
		}
	}
}

var testCasesIPv6 = []addTest{
	{"2001:db8:85a3:0:0:8A2E:0370:7334", true},
	{"2001:0db8:85a3:0:0:8A2E:0370:7334", true},
	{"2001:db8:85a3:0000:0000:8a2e:0370:7334", true},
	{"2001:0db8:85a3::8A2E:037j:7334", false},
	{"2F33:12a0:3Ea0:0302", false},
	{"I.Am.not.an.ip", false},
	{"192.168.1.0", false},
}

func TestValidateIPv6(t *testing.T) {

	for _, test := range testCasesIPv6 {
		output, err := validateIPv6(test.ip)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		} else if output != test.expectedOutput {
			t.Errorf("expected %v , got %v\n", test.expectedOutput, output)
		}
	}
}
