// GoBitFields interface
package GoBitFields

/* BitFielder interface
 * Provide functionality similar to .NET BitVector32 for addressing structs with bit fields.
 * Field order should be guaranteed.
 */
type BitFielder interface {
    CreateField(string, uint, string) error
                        // Create a field of type field (int) bits long of a given type
    SetField(string, interface{}, string) error
    Field(string)       // Read a field by name
    ReadData([]byte)    // Read a byte array into our BitFielder
    Data() []byte, int  // Output our binary blob and its intended length
    Write([]byte) (int, error) 
                        // Forced to also fulfill io.Writer
}

type BitField struct {
    data    []byte              // hold the actual data blob here
    length  uint                // length in bits (in case any must be trimmed when outputting, or buffered when inputting)
    fields  map[string] [2]int  // a map of fields with int addresses {pos, len}
    ftype   map[string] string  // a map of field types
}