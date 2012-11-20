GoBitFields
===========

golang bit field package

WARNING: This doesn't work yet

## TODO ##
Lots

## IMPLEMENTED ##

### BitFielder Interface ###
* [x] CreateField(string, uint, string) error
* [-] SetField(string, interface{}) error
* [ ] Field(string)
* [ ] ReadData([]byte)
* [ ] Data() []byte, int
* [ ] Write([]byte) (int, error) 

### Other useful functions ###
* [-] SetFieldUInt4(string, uint8) error
* [-] SetFieldBit(string, uint8) error