package variable

// Basic types

// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte alias for uint8
// rune alias for int32 represents a Unicode code point
// float32 float64
// complex64 complex128

// =============================================================================

// Zero values

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

// Example:
// var i int
// var f float64
// var b bool
// var s string

// =============================================================================

// Variable declaration

// These variables are public, they can be used across packages
var R rune = 'ข'
var S string = "กขคง"
var B bool = true
var I int = 100
var Byte byte = byte('a')

// These variable are private, they can be used only in this package
var pi float64 = 3.1415
var g complex128 = 0.867 + 0.5i

// These can be also used to declare variables
// var i, j, k int
// var b, f, s = true, 2.3, "four"
// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// Short variable declaration

// Example:
// i := 42
// f := 3.1415
// g := 0.867 + 0.5i
// a, b := 1, "b"

// =============================================================================

// Type conversions

// Example:
// var i int = 65
// var f float64 = float64(i)
// var u uint = uint(f)
// var s string = string(i)

// =============================================================================

// Constants

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.

// Example:
// const Pi = 3.14
