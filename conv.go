package conv

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

var errType = errors.New("conv: type of data source not supported")

type File string

// ToJSON convert data source to json struct or map
func ToJSON(src, obj interface{}) (err error) {
	var data []byte
	switch a := src.(type) {
	case []byte:
		data = a
	case string:
		data = []byte(a)
	case io.Reader:
		data, err = ioutil.ReadAll(a)
		if err != nil {
			return
		}
		if c, ok := a.(io.Closer); ok {
			c.Close()
		}
	case File:
		var file *os.File
		file, err = os.Open(string(a))
		if err != nil {
			return
		}
		data, err = ioutil.ReadAll(file)
		if err != nil {
			return
		}
		file.Close()
	default:
		return errType
	}
	return json.Unmarshal(data, obj)
}

// ToXML convert data source to xml struct or map
func ToXML(src, obj interface{}) (err error) {
	var data []byte
	switch a := src.(type) {
	case []byte:
		data = a
	case string:
		data = []byte(a)
	case io.Reader:
		data, err = ioutil.ReadAll(a)
		if err != nil {
			return
		}
		if c, ok := a.(io.Closer); ok {
			c.Close()
		}
	case File:
		var file *os.File
		file, err = os.Open(string(a))
		if err != nil {
			return
		}
		data, err = ioutil.ReadAll(file)
		if err != nil {
			return
		}
		file.Close()
	default:
		return errType
	}
	return xml.Unmarshal(data, obj)
}
