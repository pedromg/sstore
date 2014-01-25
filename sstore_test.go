package sstore

import (
	"testing"
	"github.com/pedromg/sstore"
)

const	fn = "test_file_sstore.sstore"
const	dsc = "This defs file for test purposes."
var	data = map[string]string{
		"aaa":"111",
		"bbb":"222",
		"ccc":"333",
	}

func TestCreateFile(t *testing.T) {
	// create test values
	var the_store sstore.SStore
	the_store.Name = fn
	the_store.Descr = dsc
	the_store.Data = data
	
	// create the file
	err := sstore.CreateFile(the_store);
	
	if err != nil {
		t.Error("Error: TestCreateFile test failed!")
	} else {
		t.Log("TestCreateFile test passed!")
	}

}

func TestReadFile(t *testing.T) {
	var the_store sstore.SStore
	var err error
	
	the_store, err = sstore.ReadFile(fn)
	
	if err!= nil {
		t.Error("Error: TestReadFile test failed!")
	} else {
		if (the_store.Name == fn) && (the_store.Descr == dsc) {
			t.Log("TestReadFile test passed!")
		} else {
			t.Error("Error: TestReadFile test conditions failed!")
		}
	}
}

func TestValues( t *testing.T) {
	var the_store sstore.SStore
	var err error
	
	the_store, err = sstore.ReadFile(fn)
	if err == nil {
		if the_store.Data["aaa"] != "111" {
			t.Error("Value for key aaa differs than expected!")
		}
		if the_store.Data["bbb"] != "222" {
			t.Error("Value for key bbb differs than expected!")
		}
		if the_store.Data["ccc"] != "333" {
			t.Error("Value for key ccc differs than expected!")
		}
		//if the_store.Data["ddd"] != nil {
		//	t.Error("Value for non existant key ddd should be nil!")
		//}
	}
}
