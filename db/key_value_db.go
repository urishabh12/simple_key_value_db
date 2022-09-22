package db

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	datastore_folder_name = "datastore"
)

type DB struct {
	name     string
	basePath string
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewDB(name string) *DB {
	path := fmt.Sprintf("./%s/%s", datastore_folder_name, name)
	os.MkdirAll(path, 0755)
	return &DB{
		name:     name,
		basePath: path,
	}
}

func (db *DB) Get(key string) (string, error) {
	path := fmt.Sprintf("%s/%s", db.basePath, key)

	data, err := db.readFromFile(path)
	if err != nil {
		return "", err
	}

	//Convert json string to Key value struct
	kv, err := convertToKeyValue(data)
	if err != nil {
		return "", err
	}

	return kv.Value, nil
}

func (db *DB) Put(key string, value string) error {
	path := fmt.Sprintf("%s/%s", db.basePath, key)

	//data := fmt.Sprintf(`{"key":"%s","value":"%s"}`, key, value)
	kv := KeyValue{
		Key:   key,
		Value: value,
	}
	data, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	return db.writeToFile(path, data)
}

func (db *DB) createFile(path string) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}

func (db *DB) writeToFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return f.Sync()
}

func (db *DB) readFromFile(path string) (string, error) {
	fInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	text := make([]byte, fInfo.Size())
	_, err = f.Read(text)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

func convertToKeyValue(data string) (KeyValue, error) {
	var kv KeyValue
	err := json.Unmarshal([]byte(data), &kv)
	if err != nil {
		return kv, err
	}

	return kv, nil
}
