package main

import ( 
	"testing"
	"os/exec"
	"fmt"
	"bytes"
)

var testCases = []struct {
	input    string
	key		 int
	expected string
}{
	{
		input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		key:	  23,
		expected: "XYZABCDEFGHIJKLMNOPQRSTUVW",
	},
	{
		input:    "ATTACKATONCE",
		key:	  4, 
		expected: "EXXEGOEXSRGI",
	},
	{
		input:    "cipher",
		key:	  6, 
		expected: "iovnkx",
	},
	{
		input:    "caesar",
		key:	  15, 
		expected: "rpthpg",
	},
	{
		input:    "TESTlolKek",
		key:	  19, 
		expected: "MXLMeheDxd",
	},
}

func TestCaesar(t *testing.T) {
	var actual bytes.Buffer
	for _, tt := range testCases {
		command := fmt.Sprintf("caesar -e --key %d %s", tt.key, tt.input)	
		caesar := exec.Command(command)
		caesar.Stdout = &actual
		err := caesar.Run()

		if err != nil {
			t.Fatalf("Caesar(%q) returned error %q.  Error not expected.", tt.input, err)
		}
		if actual.String() != tt.expected {
			t.Fatalf("Caesar(%q) was expected to return %v but returned %v.",
				tt.input, tt.expected, actual)
		}
	}
}
