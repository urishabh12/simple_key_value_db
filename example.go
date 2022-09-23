package main

import (
	"fmt"

	"github.com/urishabh12/simple_key_value_db/db"
)

func main() {
	d, err := db.NewDB("test")
	if err != nil {
		panic(err)
	}

	d.Put("john", "Accountant")
	d.Put("lisa", "Athlete")

	value, err := d.Get("john")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
