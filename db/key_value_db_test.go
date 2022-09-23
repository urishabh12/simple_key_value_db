package db

import (
	"fmt"
	"os"
	"testing"
)

const (
	dbName = "test"
)

func Test_ForCaseSensitive(t *testing.T) {
	d, err := NewDB(dbName)
	if err != nil {
		t.Error(err)
	}

	keyUpper := "KEY"
	keyLower := "key"

	d.Put(keyUpper, "value upper")
	d.Put(keyLower, "value lower")

	vLower, err := d.Get(keyLower)
	if err != nil {
		t.Fatalf(err.Error())
	}

	vUpper, err := d.Get(keyUpper)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if vLower == vUpper {
		t.Errorf("both the value are same")
	}

	err = deleteDB(d.basePath)
	if err != nil {
		t.Error(err)
	}
}

func Test_GetGivesMostRecentPut(t *testing.T) {
	d, err := NewDB(dbName)
	if err != nil {
		t.Error(err)
	}

	key := "recent"

	d.Put(key, "r")
	d.Put(key, "ran")
	d.Put(key, "random")

	v, err := d.Get(key)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if v != "random" {
		t.Errorf("value is not the most recent put")
	}

	err = deleteDB(d.basePath)
	if err != nil {
		t.Error(err)
	}
}

func Test_ConvertToKeyValue(t *testing.T) {
	data := `{"key":"key","value":"value"}`
	kv, err := convertToKeyValue(data)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if kv.Key != "key" || kv.Value != "value" {
		t.Fatalf(fmt.Sprintf("key %s and value %s not same as in data", kv.Key, kv.Value))
	}
}

func Test_PutEqualGet(t *testing.T) {
	d, err := NewDB(dbName)
	if err != nil {
		t.Error(err)
	}

	key := "abc"
	value := "xyz"

	err = d.Put(key, value)
	if err != nil {
		t.Fatalf(err.Error() + " on Put")
	}

	v, err := d.Get(key)
	if err != nil {
		t.Fatalf(err.Error() + " on Get")
	}

	if v != value {
		e := fmt.Sprintf("Get value %s not equal to put value %s", v, value)
		t.Errorf(e)
	}

	err = deleteDB(d.basePath)
	if err != nil {
		t.Error(err)
	}
}

func Test_ReadEqualToWrite(t *testing.T) {
	d, err := NewDB(dbName)
	if err != nil {
		t.Error(err)
	}

	writeData := "hello world! how are you doing"
	file := "path"
	path := fmt.Sprintf("%s/%s", d.basePath, file)
	d.createFile(path)
	d.writeToFile(path, []byte(writeData))
	readData, err := d.readFromFile(path)
	if err != nil {
		t.Errorf(err.Error())
	}

	if readData != writeData {
		e := fmt.Sprintf("length not equal of read %s and write %s data", readData, writeData)
		t.Errorf(e)
	}

	err = deleteDB(d.basePath)
	if err != nil {
		t.Error(err)
	}
}

func Test_LengthOfReadEqualToWrite(t *testing.T) {
	d, err := NewDB(dbName)
	if err != nil {
		t.Error(err)
	}

	writeData := "hello world! how are you doing"
	file := "path"
	path := fmt.Sprintf("%s/%s", d.basePath, file)
	d.createFile(path)
	d.writeToFile(path, []byte(writeData))
	readData, err := d.readFromFile(path)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(readData) != len(writeData) {
		e := fmt.Sprintf("length not equal of read %d and write %d data", len(readData), len(writeData))
		t.Errorf(e)
	}

	err = deleteDB(d.basePath)
	if err != nil {
		t.Error(err)
	}
}

func deleteDB(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	err = os.RemoveAll("./datastore")
	return err
}
