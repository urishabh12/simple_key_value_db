# Simple Key Value DB

## Example

```
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
```
