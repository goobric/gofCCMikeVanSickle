package main

// import the build-in Golang format library 'fmt'
import "fmt"

// variables can be declared at the top level Package level, outside of the function.
// the ':=' symbol can not be used outside of the function block.
// variables have to be explicitly declared.
// ALL VARIABLES THAT ARE DECLARED MUST BE USED, OTHERWISE A COMPILE RUN-TIME ERROR WILL OCCUR.

var m int = 71
var n float32 = 0.52
var p float64 = 3.14159265359

// variables can be declared within a 'var block' using parentheses ()
var (
	actorName       string = "Hawkeye"
	companion       string = "Kate Bishop"
	talesOfSuspense int    = 57
	year            int    = 1964
)

func main() {
	// declare a variable (var) called (i) of type integer (int)
	var i int
	i = 42
	// the above declaration is on 2 lines; this isn't efficient, a better solution is to declare on 1 line
	var j int = 27
	// to be even more concise, this declaration can be shorten to 'k := 60' the ':=' operator means: Declare AND assign the variable to the value. Go will decide the Data Type.
	k := 60
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	// Printf - is used to print a formatting string with percentage % symbol, with the variable name: %v is value, %T is data type; name of variable
	var l int = 99
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", p, p)

	fmt.Printf("%v, %T\n", actorName, actorName)
	fmt.Printf("%v, %T\n", companion, companion)
	fmt.Printf("%v, %T\n", talesOfSuspense, talesOfSuspense)
	fmt.Printf("%v, %T\n", year, year)

	var b bool
	fmt.Printf("%v, %T\n", b, b)

	fmt.Println()
	a := 10 // Go assigns the variable a the value 10 of type int
	c := 3  // Go assigns the variable c the value 3 of type int
	fmt.Println(a + c)
	fmt.Println(a - c)
	fmt.Println(a * c)
	fmt.Println(a / c)
	// the int division sum, only produces an integer value, NOT a float value
	fmt.Println(a % c)

	var d int = 11
	var e int8 = 4
	fmt.Println(d + int(e))
	// as these are two different int types, Golang will NOT perforn the operation.
	// one of the variables will HAVE TO BE converted. In this example 'var e' is converted to an integer

	// Bit operators
	// 11 represented as a binary number is 1011
	// 4 represented as a binary number is 0100
	fmt.Println(a & c)
	fmt.Println(a | c)
	fmt.Println(a ^ c)
	fmt.Println(a &^ c)

	f := 8
	// 8 is 2^3 : 2 to the power of 3
	fmt.Println(f << 3)
	// 2^3 * 2^3 = 2^6 = 2x2x2x2x2x2 = 64
	fmt.Println(f >> 3)
	// bit shift to the right. 2^3 / 2^3 = 2^0 = 1
	fmt.Println()
	q := 2.1e14
	fmt.Printf("%v, %T", q, q)

	fmt.Println()
	// Go has the ability to use Complex numbers with the 'i' imaginary notation
	var comNum complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", comNum, comNum)

	// the String Data type is UTF-8 encoded
	// strings in Golang are aliases for Bytes
	s := "This is a string literal"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Println("The [2] index is the lowercase letter 'i', which has a UTF-8 value of 105")

	// a string can be converted to a slice of bytes
	byt := []byte(s)
	fmt.Printf("%v, %T\n", byt, byt)

	// 'runes' are type aliases for UTF-32 encoded
	r := 'a'         //implicitly declared
	var R rune = 'b' //explicitly declared
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", R, R)
}
