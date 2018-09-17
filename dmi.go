package kerneldmi

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
)

// BoardType - platform vendor, name, serial
type BoardType struct {
	BoardVendor string
	BoardName   string
	BoardSerial string
}

// GetBoardType - get board type using kernel DMI information
func GetBoardType() (BoardType, error) {

	var board BoardType
	var err error

	board.BoardVendor, err = ReadDMIStringParameter("/sys/class/dmi/id/board_vendor")
	if err != nil {
		return BoardType{}, err
	}

	board.BoardName, err = ReadDMIStringParameter("/sys/class/dmi/id/board_name")
	if err != nil {
		return BoardType{}, err
	}

	board.BoardSerial, err = ReadDMIStringParameter("/sys/class/dmi/id/board_serial")
	if err != nil {
		return BoardType{}, err
	}

	return board, nil
}

// ReadDMIStringParameter - reads simple DMI kernel key,value(STRING) parameter; where key - filename, value - content on first line
func ReadDMIStringParameter(path string) (string, error) {

	// read file to memory
	data, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return "", err
	}

	// split by newlines
	line := bytes.Split(data, []byte("\n"))

	// fallback
	if len(line) == 0 {
		return "", nil
	}

	// get value from first line
	value := string(bytes.TrimSpace(line[0]))

	return value, nil
}
