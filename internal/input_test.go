// Module to test input module

package internal

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestReadInputFile(t *testing.T) {
	sampleValidInput := `5
5
1:1,2:0,2:3,3:4,4:3
0:1,2:3,3:0,3:4,4:1
5
0:1,4:3,2:3,3:1,4:1
0:1,0:0,1:2,2:3,4:3
`
	sampleValidConfig := Config{
		GridSize:        5,
		TotalShips:      5,
		TotalMissiles:   5,
		P1ShipPositions: "1:1,2:0,2:3,3:4,4:3",
		P2ShipPositions: "0:1,2:3,3:0,3:4,4:1",
		P1Moves:         "0:1,4:3,2:3,3:1,4:1",
		P2Moves:         "0:1,0:0,1:2,2:3,4:3",
	}

	var testCases = []struct {
		input       string
		expected    interface{}
		expectError bool
	}{
		{sampleValidInput, &sampleValidConfig, false},
		{"Invalid data", nil, true},
		{"again invalid", nil, true},
	}
	for _, tc := range testCases {
		f, err := ioutil.TempFile("/tmp", "TestReadInputFile")
		if err != nil {
			t.Errorf("Test setup error: %v", err)
		}
		f.WriteString(tc.input)
		f.Close()
		defer os.Remove(f.Name())
		actual, err := ReadInputFile(f.Name())
		if (actual == nil) && (tc.expected != nil) {
			t.Errorf("Parsed config is nil")
		} else if (actual != nil) && (!reflect.DeepEqual(actual, tc.expected)) {
			t.Errorf("Parsed config does not match")
		}
		if (err != nil) && (!tc.expectError) {
			t.Errorf("Input config parser returned error")
		}
		if (err == nil) && (tc.expectError) {
			t.Errorf("Input config parser should return error")
		}
	}
	_, err := ReadInputFile("/tmp/invalid_path_non_existent_2017")
	if err == nil {
		t.Errorf("Expected error for invalid file path, returned success")
	}
}
