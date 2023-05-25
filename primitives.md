# Summary

## Primitive Types: there are a number of different primitive types in Golang

### Boolean Type
*   **bool** the value is either: true or false; on or off; 0 or 1
*   the default value of **bool** in Golang is *false* , or off or 0
*   **Not** an alias for other types of data (eg: int)

### Numeric Types
*   Integers
    *   Signed integers
        *   int type has varying size, but minimum of 32 bits
        *   8 bit(int8) through to 64 bit (int64)
    *   Unsigned integers
        *   8 bit(byte and uint8) through to 32 bit (uint32)
        *   these values can only be positive
    *   Arithmetic operations
        *   Addition (+), Subtraction (-), Multiplication (*), Division (/), Remainder (%) modulus
    *   Bitwise operations
        *   And (&), Or (|), Xor (^), Not (&^)
    *   Zero value IS 0
    *   Can not mix types of data in same family! (uint16 + uint32 = error)
    *   Floating Point numbers
        *   Follow IEEE-754 standard
        *   Zero value IS 0
        *   32 and 64 bit versions
        *   Literal styles
            *   Decimal (3.142)
            *   Exponential (13e18 or 2E10)
            *   Mixed (15.4e12)
        *   Arithmetic operations
            *   Addition, Subtraction, Multiplication, Division. Not *remainder*
*   Complex numbers
    *   Zero value is (0+0i)
    *   64 and 128 bit versions
    *   Built-in functions
        *   complex - make complex number from two floats
        *   real - get real part as float
        *   imag - get imaginary part as float
    *   Arithmetic operations
        *   Addition, Subtraction, Multiplication, Division.

### Text Types
*   Strings
    *   UTF-8
    *   Immutable
    *   Can be concatenated with plus (+) operator
    *   Can be converted to []byte using the strconv function
*   Rune
    *   UTF-32
    *   Alias for int32
    *   Special methods normally required to process
        *   eg: strings.Reader#ReadRune
