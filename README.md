GoBitFields
===========

golang bit field package

WARNING: This might work with some commenting, but is untested and will probably error.

## TODO ##
Lots

## IMPLEMENTED ##

### BitFielder Interface ###
* [x] CreateField(string, uint, string) error
* [x] SetField(string, interface{}) error
* [x] Field(string)
* [ ] ReadData([]byte)
* [ ] Data() []byte, int
* [ ] Write([]byte) (int, error) 