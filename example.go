package main

import (
	"fmt"

	"github.com/urishabh12/simple_key_value_db/db"
)

func main() {
	d := db.NewDB("test")
	d.Put("john", "Accountant")
	d.Put("lisa", "Athlete")

	value, err := d.Get("john")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
