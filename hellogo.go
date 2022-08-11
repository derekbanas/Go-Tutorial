// A package is a collection of code
// We can define what package we want our code to belong to
// We use main when we want our code to run in the terminal
package main

// Import multiple packages
// You could use an alias like f "fmt"
import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// Create alias to long function names
var pl = fmt.Println

/*
I'm a block comment
*/

// ----- FUNCTIONS -----
func sayHello() {
	pl("Hello")
}

// Returns sum of values
func getSum(x int, y int) int {
	return x + y
}

// Return multiple values
func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

// Return potential error
func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		// Define error message returned with dummy value
		// for ans
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		// If no error return nil
		return x / y, nil
	}
}

// Variadic function
func getSum2(nums ...int) int {
	sum := 0
	// nums gets converted into a slice which is
	// iterated by range (More on slices later)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeVal2(myPtr *int) {
	*myPtr = 12
}

// Receives array by reference and doubles values
func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

// ----- FUNCTION THAT EXCEPTS GENERICS -----
// This generic type parameter is capital, between
// square brackets and has a rule for what data
// it will except called a constraint
// any : anything
// comparable : Anything that supports ==
// More Constraints : pkg.go.dev/golang.org/x/exp/constraints

// You can also define what is excepted like this
// Define that my generic must be an int or float64
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

// ----- STRUCTS -----
type customer struct {
	name    string
	address string
	bal     float64
}

// This struct has a function associated
type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Customer passed as values
func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// Struct composition : Putting a struct in another
type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

// ----- DEFINED TYPES -----
// I'll define different cooking measurement types
// so we can do conversions
type Tsp float64
type TBs float64
type ML float64

// Convert with functions (Bad Way)
func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

// Associate method with types
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

// ----- INTERFACES -----
type Animal interface {
	AngrySound()
	HappySound()
}

// Define type with interface methods and its
// own method
type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks its Prey")
}

// Return the cats name with a type conversion
func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrr")
}

// ----- CONCURRENCY -----
func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1 :", i)
	}
}
func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2 :", i)
	}
}

// These functions will print in order using
// channels
// Func receives a channel and then sends values
// over channels once each time it is called
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

// ----- BANK ACCOUNT EXAMPLE -----
// Here I'll simulate customers accessing a
// bank account and lock out customers to
// allow for individual access
type Account struct {
	balance int
	lock    sync.Mutex // Mutual exclusion
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n",
			v, a.balance)
		a.balance -= v
	}
}

// ----- CLOSURES -----
// Pass a function to a function
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

// ----- RECURSION -----
func factorial(num int) int {
	// This condition ends calling functions
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// When a Go program executes it executes a function named main
// Go statements don't require semicolons
func main() {
	// Prints text and a newline
	// List package name followed by a period and the function name
	pl("Hello Go")

	// Get user input (To run this in the terminal go run hellogo.go)
	pl("What is your name?")
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	// Copy text up to the newline
	// The blank identifier _ will get err and ignore it (Bad Practice)
	// name, _ := reader.ReadString('\n')
	// It is better to handle it
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		// Log this error
		log.Fatal(err)
	}

	// ----- VARIABLES -----
	// var name type
	// Name must begin with letter and then letters or numbers
	// If a variable, function or type starts with a capital letter
	// it is considered exported and can be accessed outside the
	// package and otherwise is available only in the current package
	// Camal case is the default naming convention

	// var vName string = "Derek"
	// var v1, v2 = 1.2, 3.4

	// Short variable declaration (Type defined by data)
	// var v3 = "Hello"

	// Variables are mutable by default (Value can change as long
	// as the data type is the same)
	// v1 := 2.4

	// After declaring variables to assign values to them always use
	// = there after. If you use := you'll create a new variable

	// ----- DATA TYPES -----
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('🦍'))

	// ----- CASTING -----
	// To cast type the type to convert to with the variable to
	// convert in parentheses
	// Doesn't work with bools or strings
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	// Convert string to int (ASCII to Integer)
	// Returns the result with an error if any
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// Convert int to string (Integer to ASCII)
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	// Convert string to float
	cV7 := "3.14"
	// Handling potential errors (Prints if err == nil)
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	// Use Sprintf to convert from float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)

	// ----- IF CONDITIONAL -----
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not and Important Birthday")
	}

	// ! turns bools into their opposite value
	pl("!true =", !true)

	// ----- STRINGS -----
	// Strings are arrays of bytes []byte
	// Escape Sequences : \n \t \" \\
	sV1 := "A word"

	// Replacer that can be used on multiple strings
	// to replace one string with another
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	// Get length
	pl("Length : ", len(sV2))

	// Contains string
	pl("Contains Another :", strings.Contains(sV2, "Another"))

	// Get first index match
	pl("o index :", strings.Index(sV2, "o"))

	// Replace all matches with 0
	// If -1 was 2 it would replace the 1st 2 matches
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	// Remove whitespace characters from beginning and end of string
	sV3 := "\nSome words\n"
	sV3 = strings.TrimSpace(sV3)

	// Split at delimiter
	pl("Split :", strings.Split("a-b-c-d", "-"))

	// Upper and lowercase string
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))

	// Prefix or suffix
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

	// ----- RUNES -----
	// In Go characters are called Runes
	// Runes are unicodes that represent characters
	rStr := "abcdefg"

	// Runes in string
	pl("Rune Count :", utf8.RuneCountInString(rStr))

	// Print runes in string
	for i, runeVal := range rStr {
		// Get index, Rune unicode and character
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	// ----- TIME -----
	// Get day, month, year and time data
	// Get current time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	// ----- MATH -----
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)

	// Shorthand increment
	// Instead of mInt = mInt + 1 (mInt += 1)
	// -= *= /=
	mInt := 1
	mInt += 1
	// Also increment or decrement with ++ and --
	mInt++

	// Float precision increases with the size of your values
	pl("Float Precision =", 0.11111111111111111111+
		0.11111111111111111111)

	// Create a random value between 0 and 50
	// Get a seed value for our random number generator based on
	// seconds since 1/1/70 to make our random number more random
	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	// There are many math functions
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))

	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot

	// ----- FORMATTED PRINT -----
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A',
		3.14, true, 1, 1)

	// Float formatting
	fmt.Printf("%9f\n", 3.14)      // Width 9
	fmt.Printf("%.2f\n", 3.141592) // Decimal precision 2
	fmt.Printf("%9.f\n", 3.141592) // Width 9 no precision

	// Sprintf returns a formatted string instead of printing
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)

	// ----- FOR LOOPS -----
	// for initialization; condition; postStatement {BODY}
	// Print numbers 1 through 5
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	// Do the opposite
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	// x is out of the scope of the for loop so it doesn't exist
	// pl("x :", x)

	// For is used to create while loops as well
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// While true loop (Infinite Loop) will be used for a guessing
	// game
	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 : ")
		pl("Random Number is :", randNum)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Lower")
		} else if iGuess < randNum {
			pl("Higher")
		} else {
			pl("You Guessed it")
			break
		}

		// Cycle through an array with range
		// More on arrays later
		// We don't need the index so we ignore it
		// with the blank identifier _
		aNums := []int{1, 2, 3}
		for _, num := range aNums {
			pl(num)
		}
	}

	// ----- ARRAYS -----
	// Collection of values with the same data type
	// and the size can't be changed
	// Default values are 0, 0.0, false or ""

	// Declare integer array with 5 elements
	var arr1 [5]int

	// Assign value to index
	arr1[0] = 1

	// Declare and initialize
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Get by index
	pl("Index 0 :", arr2[0])

	// Length
	pl("Arr Length :", len(arr2))

	// Iterate with index
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	// Iterate with range
	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v)
	}

	// Multidimensional Array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// Print multidimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	// String into slice of runes
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)

	// ----- SLICES -----
	// Slices are like arrays but they can grow
	// var name []dataType
	// Create a slice with make
	sl1 := make([]string, 6)

	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	// Size of slice
	pl("Slice Size :", len(sl1))

	// Cycle with for
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	// Cycle with range
	for _, x := range sl1 {
		pl(x)
	}

	// Create a slice literal
	sl2 := []int{12, 21, 1974}
	pl(sl2)

	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array
	sArr := [5]int{1, 2, 3, 4, 5}
	// Start at 0 index up to but not including the 2nd index
	sl3 := sArr[0:2]
	pl(sl3)

	// Get slice from beginning
	pl("1st 3 :", sArr[:3])

	// Get slice to the end
	pl("Last 3 :", sArr[2:])

	// If you change the array the slice also changes
	sArr[0] = 10
	pl("sl3 :", sl3)

	// Changing the slice also changes the array
	sl3[0] = 1
	pl("sArr :", sArr)

	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	// Printing empty slices will return nils which show
	// as empty slices
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])

	// ----- FUNCTIONS -----
	// func funcName(parameters) returnType {BODY}
	// If you only need a function in the current package
	// start with lowercase letter
	// Letters and numbers in camelcase
	sayHello()
	pl(getSum(5, 4))
	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)

	// Function that can return an error
	ans, err := getQuotient(5, 0)
	if err == nil {
		pl("5/4 =", ans)
	} else {
		pl(err)
		// End program
		// log.Fatal(err)
	}

	// Function receives unknown number of parameters
	// Variadic Function
	pl("Unknown Sum :", getSum2(1, 2, 3, 4))

	// Pass an array to a function by value
	vArr := []int{1, 2, 3, 4}
	pl("Array Sum :", getArraySum(vArr))

	// Go passes the value to functions so it isn't changed
	// even if the same variable name is used
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)

	// ----- POINTERS -----

	// You can pass by reference with the &
	// (Address of Operator)
	// Print amount and address for amount in memory
	f4 := 10
	pl("f4 :", f4)
	pl("f4 Address :", &f4)

	// Store a pointer (Pointer to int)
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)

	// Print value at pointer
	pl("f4 Value :", *f4Ptr)

	// Assign value using pointer
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	// Change value in function
	pl("f4 before function :", f4)
	changeVal2(&f4)
	pl("f4 after function :", f4)

	// Pass an array by reference
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	// Passing a slice to a function works just
	// like when using variadic functions
	// Just add ... after the slice when passing
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))

	// ----- FILE IO -----
	// We can create, write and read from files

	// Create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Says to close the file after program ends or when
	// there is a closing curly bracket
	defer f.Close()

	// Create list of primes
	iPrimeArr := []int{2, 3, 5, 7, 11}
	// Create string array from int array
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// Cycle through strings and write to file
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the created file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read from file and print once per line
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// Append to file
	/*
		 Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

		O_RDONLY : open the file read-only
		O_WRONLY : open the file write-only
		O_RDWR   : open the file read-write

		These can be or'ed

		O_APPEND : append data to the file when writing
		O_CREATE : create a new file if none exists
		O_EXCL   : used with O_CREATE, file must not exist
		O_SYNC   : open for synchronous I/O
		O_TRUNC  : truncate regular writable file when opened
	*/

	// Check if file exists
	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err = os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

	// ----- COMMAND LINE ARGUMENTS -----
	// You can pass values to your program
	// from the command line
	// Create cltest.go
	// go build cltest.go
	// .\cltest 24 43 12 9 10
	// Returns an array with everything
	// passed with the name of the app
	// in the first index
	// Outputs the max number passed in

	// ----- PACKAGES -----
	// Packages allow you to keep related code together
	// Go looks for package code in a directory

	// If you are using VSC and have multiple
	// modules you get this error
	// gopls requires a module at the root of
	// your workspace
	// 1. Settings
	// 2. In search type gopls
	// 3. Paste "gopls": { "experimentalWorkspaceModule": true, }
	// 4. Restart VSC

	// cd /D D:\Tutorials\GoTutorial

	// Create a go directory : mkdir app
	// cd app
	// Choose a module path and create a go.mod file
	// go mod init example/project

	// Go modules allow you to manage libraries
	// They contain one project or library and a
	// collection of Go packages
	// go.mod : contains the name of the module and versions
	// of other modules your module depends on

	// Create a main.go file at the same level as go.mod

	// You can have many packages and sub packages
	// create a directory called mypackage in the project
	// directory mkdir mypackage
	// cd mypackage

	// Create file mypackage.go in it

	// Package names should be all lowercase

	// ----- MAPS -----
	// Maps are collections of key/value pairs
	// Keys can be any data type that can be compared
	// using == (They can be a different type than
	// the value)
	// var myMap map [keyType]valueType

	// Declare a map variable
	var heroes map[string]string
	// Create the map
	heroes = make(map[string]string)

	// You can do it in one step
	villians := make(map[string]string)

	// Add keys and values
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	// Define with map literal
	superPets := map[int]string{1: "Krypto",
		2: "Bat Hound"}

	// Get value with key (Use %v with Printf)
	fmt.Printf("Batman is %v\n", heroes["Batman"])

	// If you access a key that doesn't exist
	// you get nil
	pl("Chip :", superPets[3])

	// You can check if there is a value or nil
	_, ok := superPets[3]
	pl("Is there a 3rd pet :", ok)

	// Cycle through map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// Delete a key value
	delete(heroes, "The Flash")

	// ----- GENERICS -----
	// We can specify the data type to be used at a
	// later time with generics
	// It is mainly used when we want to create
	// functions that can work with
	// multiple data types
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))

	// This causes an error
	// pl("5.6 + 4.7 =", getSumGen("5.6", "4.7"))

	// ----- STRUCTS -----
	// Structs allow you to store values with many
	// data types

	// Add values
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.bal = 234.56

	// Pass to function as values
	getCustInfo(tS)
	// or as reference
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)

	// Create a struct literal
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	// Structs with functions
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())

	// Go doesn't support inheritance, but it does
	// support composition by embedding a struct
	// in another
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}

	bus1.info()

	// ----- DEFINED TYPES -----
	// We used a defined type previously with structs
	// You can use them also to enhance the quality
	// of other data types
	// We'll create them for different measurements

	// Convert from tsp to mL
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f mL\n", ml1)

	// Convert from TBs to mL
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f mL\n", ml2)

	// You can use arithmetic and comparison
	// operators
	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	// We can convert with functions
	// Bad Way
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 TBs = %.2f mL\n", TBToML(3))

	// We can solve this by using methods which
	// are functions associated with a type
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())

	// ----- PROTECTING DATA -----
	// We want to protect our data from receiving
	// bad values by moving our date struct
	// to another package using encapsulation
	// We'll use mypackage like before

	// ----- INTERFACES -----
	// Interfaces allow you to create contracts
	// that say if anything inherits it that
	// they will implement defined methods

	// If we had animals and wanted to define that
	// they all perform certain actions, but in their
	// specific way we could use an interface

	// With Go you don't have to say a type uses
	// an interface. When your type implements
	// the required methods it is automatic
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	// We can only call methods defined in the
	// interface for Cats because of the contract
	// unless you convert Cat back into a concrete
	// Cat type using a type assertion
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())

	// ----- CONCURRENCY -----
	// Concurrency allows us to have multiple
	// blocks of code share execution time by
	// pausing their execution. We can also
	// run blocks of codes in parallel at the same
	// time. (Concurrent tasks in Go are called
	// goroutines)

	// To execute multiple functions in new
	// goroutines add the word go in front of
	// the function calls (Those functions can't
	// have return values)

	// We can't control when functions execute
	// so we may get different results
	go printTo15()
	go printTo10()

	// We have to pause the main function because
	// if main ends so will the goroutines
	time.Sleep(2 * time.Second) // Pause 2 seconds

	// You can have goroutines communicate by
	// using channels. The sending goroutine
	// also makes sure the receiving goroutine
	// receives the value before it attempts
	// to use it

	// Create a channel : Only carries values of
	// 1 type
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)

	// Using locks to protect data from being
	// accessed by more than one user at a time
	// Locks are another option when you don't
	// have to pass data
	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

	// ----- CLOSURES -----
	// Closures are functions that don't have to be
	// associated with an identifier (Anonymous)

	// Create a closure that sums values
	intSum := func(x, y int) int { return x + y }
	pl("5 + 4 =", intSum(5, 4))

	// Closures can change values outside the function
	samp1 := 1
	changeVar := func() { samp1 += 1 }
	changeVar()
	pl("samp1 =", samp1)

	// Pass a function to a function
	useFunc(sumVals, 5, 8)

	// ----- RECURSION -----
	// Recursion occurs when a function calls itself
	// There must be a condition that ends this
	// Finding a factorial is commonly used
	pl("Factorial 4 =", factorial(4))
	// 1st : result = 4 * factorial(3) = 4 * 6 = 24
	// 2nd : result = 3 * factorial(2) = 3 * 2 = 6
	// 3rd : result = 2 * factorial(1) = 2 * 1 = 2

	// ----- REGULAR EXPRESSIONS -----
	// You can use regular expressions to test
	// if a string matches a pattern

	// Search for ape followed by not a space
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	pl(match)

	// You can compile them
	// Find multiple words ending with at
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")

	// Did you find any matches?
	pl("MatchString :", r.MatchString(reStr2))

	// Return first match
	pl("FindString :", r.FindString(reStr2))

	// Starting and ending index for 1st match
	pl("Index :", r.FindStringIndex(reStr2))

	// Return all matches
	pl("All String :", r.FindAllString(reStr2, -1))

	// Get 1st 2 matches
	pl("All String :", r.FindAllString(reStr2, 2))

	// Get indexes for all matches
	pl("All Submatch Index :", r.FindAllStringSubmatchIndex(reStr2, -1))

	// Replace all matches with Dog
	pl(r.ReplaceAllString(reStr2, "Dog"))

	// ----- AUTOMATED TESTING -----
	// Automated tests make sure your program still
	// works while you change the code
	// Create app2 directory with testemail.go
	// cd app2
	// Create testemail_test.go
	// go mod init app2
	// Run Tests : go test -v

	// ----- GO ON THE WEB -----
	// Our Go apps can run in the browser
	// Create directory webapp with webapp.go
	// Run our app with go run webapp.go
	// Ctrl + c to stop server
}
