// Package sstore creates or opens a simple definitions key:value file.
// The first line is used to describe the file, so contains no usable data.
package sstore

import (
	"io/ioutil"
)

type sStore struct {
	Name	string
	Descr	string
	Data	map[string]string
}

// CreateFile creates a file with name, inserts the firstline descriptive text 
// on the first line, and then the data key:value map of strings.
func CreateFile(s sStore) error {
	// set the map 
	var d []byte
	d = append(d, s.Descr...)
	d = append(d, "\n"...)
	d = append(d, map2bytes(s.Data)...)
	err := ioutil.WriteFile(s.Name, d, 1)
	if err != nil {
		return err
	}
	return nil
}

func map2bytes(m map[string]string) []byte {
	var b []byte
	for k, v := range m {
		b = append(b, k+":"+v...)
		b = append(b, "\n"...)
	}
	return b
}
