package kerneldmi

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
)

// Type - platform vendor, name, serial
type Type struct {
	Vendor string
	Name   string
	Serial string
}

// GetBoardType - get board type using kernel DMI information
func GetBoardType() (Type, error) {

	var board Type
	var err error

	board.Vendor, err = ReadDMIStringParameter("/sys/class/dmi/id/board_vendor")
	if err != nil {
		return Type{}, err
	}

	board.Name, err = ReadDMIStringParameter("/sys/class/dmi/id/board_name")
	if err != nil {
		return Type{}, err
	}

	board.Serial, err = ReadDMIStringParameter("/sys/class/dmi/id/board_serial")
	if err != nil {
		return BoardType{}, err
	}

	return board, nil
}

// GetChassisType - get chassis type using kernel DMI information
func GetChassisType() (Type, error) {

	var chassis Type
	var err error

	chassis.Vendor, err = ReadDMIStringParameter("/sys/class/dmi/id/chassis_vendor")
	if err != nil {
		return Type{}, err
	}

	chassis.Name, err = ReadDMIStringParameter("/sys/class/dmi/id/chassis_name")
	if err != nil {
		return Type{}, err
	}

	chassis.Serial, err = ReadDMIStringParameter("/sys/class/dmi/id/chassis_serial")
	if err != nil {
		return Type{}, err
	}

	return chassis, nil
}

// GetProductType - get product type using kernel DMI information
func GetProductType() (Type, error) {

	var product Type
	var err error

	product.Vendor, err = ReadDMIStringParameter("/sys/class/dmi/id/product_vendor")
	if err != nil {
		return Type{}, err
	}

	product.Name, err = ReadDMIStringParameter("/sys/class/dmi/id/product_name")
	if err != nil {
		return Type{}, err
	}

	product.Serial, err = ReadDMIStringParameter("/sys/class/dmi/id/product_serial")
	if err != nil {
		return Type{}, err
	}

	return product, nil
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
