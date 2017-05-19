package sstore

import (
	"testing"
)

const fn = "test_file_sstore.sstore"
const dsc = "This defs file for test purposes."

var data = map[string]string{
	"aaa": "111",
	"bbb": "222",
	"ccc": "333",
}

func TestCreateFile(t *testing.T) {
	// create test values
	var theStore SStore
	theStore.Name = fn
	theStore.Descr = dsc
	theStore.Data = data

	// create the file
	err := CreateFile(theStore)

	if err != nil {
		t.Error("Error: TestCreateFile test failed!")
	} else {
		t.Log("TestCreateFile test passed!")
	}

}

func TestReadFile(t *testing.T) {
	var theStore SStore
	var err error

	theStore, err = ReadFile(fn)

	if err != nil {
		t.Error("Error: TestReadFile test failed!")
	} else {
		if (theStore.Name == fn) && (theStore.Descr == dsc) {
			t.Log("TestReadFile test passed!")
		} else {
			t.Error("Error: TestReadFile test conditions failed!")
		}
	}
}

func TestValues(t *testing.T) {
	var theStore SStore
	var err error

	theStore, err = ReadFile(fn)
	if err == nil {
		if theStore.Data["aaa"] != "111" {
			t.Error("Value for key aaa differs than expected!")
		}
		if theStore.Data["bbb"] != "222" {
			t.Error("Value for key bbb differs than expected!")
		}
		if theStore.Data["ccc"] != "333" {
			t.Error("Value for key ccc differs than expected!")
		}
		//if theStore.Data["ddd"] != nil {
		//	t.Error("Value for non existant key ddd should be nil!")
		//}
	}
}
