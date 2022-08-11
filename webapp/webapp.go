package main

// net/http allows us to respond and to make
// server requests
import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

func errorCheck(err error) {
	// Handle errors
	if err != nil {
		log.Fatal(err)
	}
}

// The writer allows us to write to the browser
// Create a message and add it to the
// response that displays in the browser
func write(writer http.ResponseWriter, msg string) {
	// Perform type conversion to bytes
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

// request is the request from the browser
func englishHandler(writer http.ResponseWriter,
	request *http.Request) {
	write(writer, "Hello Internet")
}

func spanishHandler(writer http.ResponseWriter,
	request *http.Request) {
	write(writer, "Hola Internet")
}

func frenchHandler(writer http.ResponseWriter,
	request *http.Request) {
	write(writer, "Bonjour Internet")
}

func interactHandler(writer http.ResponseWriter,
	request *http.Request) {

	// Get our text from the file
	todoVals := getStrings("todos.txt")

	// Print to the terminal
	fmt.Printf("%#v\n", todoVals)
	// Create a template using the html
	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)

	// Create a todo list with the number
	todos := ToDoList{
		ToDoCount: len(todoVals),
		ToDos:     todoVals,
	}

	// Write the template to the ResponseWriter
	// Pass the todo struct data
	err = tmpl.Execute(writer, todos)
}

// Retreives lines of text from a file
func getStrings(fileName string) []string {
	var lines []string

	// Try to open the file (It must exist)
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	// Close file when the function ends
	defer file.Close()

	// Read lines of text and save to lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())

	// Return the text
	return lines
}

func newHandler(writer http.ResponseWriter,
	request *http.Request) {

	// Create a template using the html
	tmpl, err := template.ParseFiles("new.html")
	errorCheck(err)

	err = tmpl.Execute(writer, nil)
}

func createHandler(writer http.ResponseWriter,
	request *http.Request) {
	todo := request.FormValue("todo")
	// Define options for working with the file
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// Open file with options and permissions
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	// Append new text to file
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	// Close file
	err = file.Close()
	errorCheck(err)
	// Redirect to defined page while passing
	// ResponseWriter, original request,
	// and a successful request message
	http.Redirect(writer, request, "/interact", http.StatusFound)
}

func main() {
	// Our app is available at directory
	// hello for the localhost port 8080
	// When it receives a request it calls
	// the correct Handler
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hola", spanishHandler)
	http.HandleFunc("/bonjour", frenchHandler)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)

	// Listens for browser requests and responds
	// Only receives a value if there is an error
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
