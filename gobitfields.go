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
    return
}

/* SetField
 * Set field name "fname" to arbitrary data "fdata." Truncate fdata to fit within the field if necessary.
 */
func (BitField b) SetField(fname string, fdata interface{}, ftype string) err error {
    fpos = b.field[fname][0]
    flen = b.field[fname][1]
    ftype = b.ftype[fname]
    // startbyte and endbyte
    sbyt ebyt := locateField(fpos, flen)
    // left-side offset (for generating the mask)
    loff := fpos - (sbyt * 8)
    // generate the mask
    mask := mask(flen, loff)
    // calculate the offset using the last byte of the mask
    offset := getOffset(mask[len(mask) - 1])
    // set the data
    b.data = setdata(b.data, fdata, mask, sbyt, ebyt, offset, ftype)

    // completed successfully
    return
}