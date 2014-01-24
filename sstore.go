// Package sstore creates or opens a simple definitions key:value file.
// The first line is used to describe the file, so contains no usable data.
package sstore

import (
	"io/ioutil"
	"strings"
)

type SStore struct {
	Name	string
	Descr	string
	Data	map[string]string
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
	var the_store SStore
	the_store.Name = f
	content, error := ioutil.ReadFile(the_store.Name)
	if error != nil {
		return nil, error
	}
	lines := strings.Split(string(content), "\n")
	the_store.Descr = lines[0]
	for l := range lines {
		the_l := strings.Split(l,":")
		the_store.Data[l][0] = the_l[0]
		the_store.Data[l][1] = the_l[1]
	}
	return the_store, nil
}
