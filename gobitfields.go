// GoBitFields package
package GoBitFields

import (
    "errors"
    "math"
    "reflect"
)

/* CreateField
 * Create a field with name "fname" of length "bitlen" within BitField.
 * Store the type "typeof" within our field map.
 */
func (BitField b) CreateField(fname string, bitlen uint, typeof reflect.Type) {
    // Add the field to the field map
    if _, present := b.fields[fname]; !present {
        b.fields[fname] = []interface{} { b.length, bitlen, typeof }
    } else {
        return errors.New("Field already exists")
    }

    // increase the size of the data blob
    b.length += bitlen
    // new binary blob length, rounded up to the nearest byte for padding
    // because Go doesn't specify any types smaller than a byte
    newlen = math.Ciel(ufloat(b.length)/8.0)
    // Grow the byte slice accordingly
    temp = make([]byte, len(b.data), newlen)
    copy(temp, b.data)
    b.data = temp

    // Completed successfully
    return nil
}

/* SetField
 * Set field name "fname" to arbitrary data "fdata." Truncate fdata to fit within the field if necessary.
 */
func (BitField b) SetField(fname string, fdata interface{}) err error {
    fpos = b.field[fname][0]
    flen = b.field[fname][1]
    ftype = b.field[fname][2]

    if reflect.TypeOf(fdata).Name() != ftype {
        return errors.New("Data Type Mismatch")
    }
}

/* SetFieldUInt4
 * Same as SetField, but setting a nibble uint (uint4) field
 */
func (BitField b) SetFieldUInt4(fname string, fdata uint8) {
    fpos = b.field[fname][0]
    flen = b.field[fname][1]
    ftype = b.field[fname][2]

    if ftype != "uint4" {
        return errors.New("Field Data Type Not UInt4")
    }
}