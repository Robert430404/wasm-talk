package main

import (  
  "fmt"
  "syscall/js"
	"math/rand"
)

// Actual function for generating our numbers
func generateNumbers(threshold int) (int) {  
  var result int

	for i := 0; i < threshold; i++ {
		result = rand.Int();
	}

	return result
}

// Function for wrapping our core function so we can use it
func generateNumbersWrapper() js.Func {  
	// Generate a js function for bindig to the browser
	wrappedNumbersFunction := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// We have a required argument, make sure it's present
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		// Make sure we can resolve the document
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return "Unable to get document object"
		}

		// Make sure our output field exists
		output := jsDoc.Call("querySelector", ".Demo__webassemblyOutput")
		if !output.Truthy() {
			return "Unable to get the output area"
		}

		// Get the passed threshold
		inputThreshold := args[0].Int()

		// Actually generate the number
		generatedNumber := generateNumbers(inputThreshold)

		// Set the value on the output
		output.Set("value", generatedNumber)

		return nil
	})

	// Return our bound wrapped function
	return wrappedNumbersFunction
}

func main() {
	// Output our bootstrap message
  fmt.Println("We're in.")

	// Bind it to JS
  js.Global().Set("generateNumber", generateNumbersWrapper())

	// Make sure we don't die
	select {}
}