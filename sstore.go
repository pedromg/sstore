// Package sstore creates or opens a simple definitions key:value file.
// The first line is used to describe the file, so contains no usable data.
package sstore

import (
	"io/ioutil"
	"strings"
)

type SStore struct {
	Name  string
	Descr string
	Data  map[string]string
}

// CreateFile creates a file with name, inserts the firstline descriptive text
// on the first line, and then the data key:value map of strings.
func CreateFile(s SStore) error {
	// set the map
	var d []byte

	d = append(d, s.Descr...)
	d = append(d, "\n"...)
	d = append(d, map2bytes(s.Data)...)

	err := ioutil.WriteFile(s.Name, d, 0777)

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

func ReadFile(f string) (SStore, error) {
	var theStore SStore
	var lines, liness []string
	var l int

	theStore.Data = make(map[string]string)
	theStore.Name = f

	content, error := ioutil.ReadFile(theStore.Name)

	if error != nil {
		return theStore, error
	}

	// Populate the sstore.SStore structure.
	lines = strings.Split(string(content), "\n")
	theStore.Descr = lines[0]
	for l = 1; l < len(lines)-1; l += 1 {
		liness = strings.Split(lines[l], ":")
		theStore.Data[liness[0]] = liness[1]
	}

	// return
	return theStore, nil
}
