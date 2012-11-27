// bitops.go
// provide utility functions for locating specific bits or fields of bits within a []byte
// and for generating appropriate bitmasks to read or update them from the appropriate slice

package GoBitFields

import (
    "errors"
    "math"
    "math/big"
    "strings"
)

/* locateField
 * Returns the start byte (sbyt) and the end byte (ebyt)
 * of the field requested by pos and length.
 */
func locateField(pos int, length int) (sbyt int, ebyt int) {
    // length is in bits, so we'll find the byte that contains the start of our field
    sbyt := math.Floor(ufloat(pos)/8.0)
    // now find the byte that contains the end of the field
    ebyt := math.Floor(ufloat(length)/8.0) + sbit

    return
}

/* mask
 * generate a bit mask 
 */
func mask(mlen int, offset int) mask []int8 {
    mask := make([]int8, mlen.Ciel(ufloat(mlen)/8.0), mlen.Ciel(ufloat(mlen)/8.0))
    for i := 0; i < len(mask); i++ {
        switch {
        case i == 0:
            // this will produce the start bit of our mask, based on the offset.
            // Mask will start at bit 'offset'
            // example: offset 1 is (2^7)-1 or 0111 1111
            mask[i] = int8(math.Pow(2, 8-offset) - 1)

        default:
            // 1111 1111
            // this will appear on middle elements that should always be all filled in
            mask[i] = int8(255)

        case i == len(mask) - 1 && len(mask) > 1:
            // how many bits are left?
            // this will stop our mask at position 'bitsleft'
            // example: 5 bitsleft is 256 - (2^(8-5)) = 256 - (2^3) = 248 or 1111 1000
            bitsleft := mlen - ((len(mask) - 1) * 8)
            mask[i] = int8(256 - math.Pow(2, 8 - bitsleft))
        }
    }

    return
}

func byteArrToBigInt(data []byte) num *big.Int {
    tempInt = big.NewInt(0)
    for i := 0; i < len(data); i++ {
        temp := data[i] & mask[i]
        // ignore the return value because this sets tempInt to the desired result *and* returns
        _ = tempInt.Add(tempInt, big.NewInt(int64(temp)))

        if i < len(data) - 1 {
            // shift everything to the left 8 bits on all but the final byte
            // as above, ignore the return value
            _ = tempInt = tempInt.Mul(tempInt, big.NewInt(256))
        }
    }
    return
}

func getOffset(lastByte byte) offset int {
    // TODO: there's probably a better way to do this
    // but it's probably fast enough for now, as we'll be iterating 
    // no more than 8 times per field and these should be fast ops
    offset := 0

    for i:= 0; i < 8; i++ {
        if lastByte & 1 == 0 {
            offset += 1
            lastByte = lastByte >> 1
        } else {
            break
        }
    }

    return
}

/* maskdata
 * Mask off the desired data, and return it in the format requested.
 * In this case, offset is how many bits after the field there are to the next byte boundary
 */
func maskdata(data []byte, mask []byte, offset int, type string) (data interface{}, err error) {
    err := nil

    // make sure data and mask are the same length
    if len(data) != len(byte) {
        return _, errors.New("'data' and 'mask' length mismatch")
    }

    // let's do our actual mask operation to isolate the data we actually want
    tempInt := byteArrToBigInt(data)

    // shift everything to the right by 'offset' bits
    // as above, ignore the return value
    _ = tempInt.Quo(tempInt, big.NewInt( math.Pow(2, offset) ))

    // Now try to convert to the given type and return
    switch {
        // return type int
        case strings.Contains(type, "int"):
            data := tempInt.Int64()

        case type == "string":
            data := string(tempInt.Bytes())

        // TODO: add more of these as needed, but this should be good for now
    }

    return
}

/* setdata
 * does the opposite of maskdata - write some data into a block of blob []byte
 * and then return the full blob.
 * In this case, as with maskdata, offset is the number of bits after the field there are to the next byte boundary
 */
func setdata(blob []byte, data interface{}, mask []byte, sbyt int, ebyt int, offset int, type string) (blob []byte, err error) {
    err := nil

    // first convert data to []byte
    switch {
    // integers
    // TODO: there might be a better way of doing this
    case type == "int4": // this is not a real type
        bdata := byte( int(data & 15) ) // truncates to 4 bits

    case type == "int8":
        bdata := byte( int(data & 255))

    case type == "int12": // this is not a real type
        bdata := []byte{ (data >> 8) & 15, data & 255 } // int4 and int8

    case type == "int16":
        bdata := []byte{ (data >> 8) & 255, data & 255 }

    case type == "int32":
        bdata := make([]byte, 4, 4)
        for i := 0; i < 4; i++ {
            bdata[i] = (data >> (24-(i*8))) & 255
        }

    case type == "int64":
        bdata := make([]byte, 8, 8)
        for i := 0; i < 8; i++ {
            bdata[i] = (data >> (56-(i*8))) & 255
        }

    // strings
    case type == "string":
        bdata := []byte(data)

        // TODO: add more of these as needed, but this should be good for now
    }

    // now convert bdata to Int as in maskdata
    tempInt := byteArrToBigInt(bdata)
    // add our offset by shifting left by 'offset' bits (or multiply by 2^offset, in this case)
    _ = tempInt.Mul(tempInt, big.NewInt( math.Pow(2, offset) ))

    // now convert back to []byte
    bdata = tempInt.Bytes()

    // isolate the piece of the blob that we need
    bpart := blob[sbyt:ebyt]

    // is our field long enough? if not, set error and continue. Just notify that data has been truncated to fit.
    if len(bdata) > len(bpart) {
        err := errors.New("Data too large to fit into field. Data has been truncated to fit.")
    }

    // blank the data and then write in the new stuff
    for i := 0; i < len(bpart); i++ {
        // invert the mask so we can blank the existing data (if any) with bmask & bpart
        bmask := 255 ^ mask[i]
        bpart[i] = bpart[i] & bmask
        // now set the relevant data with |
        bpart[i] = bpart[i] | bdata[i]
    }

    // now copy the data back into blob and we're done!
    for i := sbyt; i <= ebyt; i++ {
        blob[i] = bpart[i]
    }

    return
}