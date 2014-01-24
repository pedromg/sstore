// Package sstore creates or opens a simple definitions key:value file.
// The first line is used to describe the file, so contains no usable data.
package sstore

import (
	"ioutil"
)

type sStore struct {
	Name	string
	Descr	string
	Data	map[string]string
}

// CreateFile creates a file with name, inserts the firstline descriptive text 
// on the first line, and then the data key:value map of strings.
func CreateFile(name, descr string, data map[string]string) error {
	// set the map 
	d := []byte(descr+"\n") + []byte(data)
	err := ioutil.WriteFile(name, d, nil)
	if err != nil {
		return err
	}
}
