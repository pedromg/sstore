package sstore

import (
	"fmt"
	"testing"
	"github.com/pedromg/sstore"
)

func TestCreateFile(t *testing.T) {
	// create test values
	var the_store sstore.sStore
	the_store.Name = "test_file_sstore.ss"
	the_store.Descr = "This defs file for test purposes."
	the_store.Data = map[string]string{
		"aaa":"111",
		"bbb":"222",
		"ccc":"333",
	}
	// create the file
	if err := sstore.CreateFile(fn, descr, data), err != nil {
		t.Error("Error creating the file.")
	}



