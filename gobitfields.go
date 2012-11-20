// GoBitFields package
package GoBitFields

import (
    "errors"
    "math"
)

/* CreateField
 * Create a field with name "fname" of length "bitlen" within BitField.
 * Store the type "typeof" within our field map.
 */
func (BitField b) CreateField(fname string, bitlen uint, typeof string) {
    // Add the field to the field map
    if _, present := b.fields[fname]; !present {
        b.fields[fname] = []int { b.length, bitlen }
        b.ftype[fname] = typeof
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
    ftype = b.ftype[fname]
    // TODO: this

    // completed successfully
    return nil
}

/* SetFieldUInt4
 * Same as SetField, but setting a nibble uint (uint4) field
 */
func (BitField b) SetFieldUInt4(fname string, fdata uint8) {
    fpos = b.field[fname][0]
    flen = b.field[fname][1]
    ftype = b.ftype[fname]
    // TODO: this

    // completed successfully
    return nil
}

/* SetFieldBit
 * Same as SetField, but setting a single bit field
 */
 func (BitField b) SetFieldBit(fname string, fdata uint8) {
    fpos = b.field[fname][0]
    flen = b.field[fname][1]
    ftype = b.ftype[fname]
    // TODO: this

    // completed successfully
    return nil
 }