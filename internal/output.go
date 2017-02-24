// Module to write output to a file

package internal

import (
	// "bufio"
	// "fmt"
	"io/ioutil"
	// "os"
)

func WriteResult(filename, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
