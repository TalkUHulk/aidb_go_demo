package aidb

import (
	"io/ioutil"
)

func readBinaryFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}
