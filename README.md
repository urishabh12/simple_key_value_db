# simple key value db

Example
```
d := db.NewDB("test")

d.Put("john", "Accountant")
d.Put("lisa", "Athlete")

value, err := d.Get("john")
if err != nil {
	panic(err)
}
```
