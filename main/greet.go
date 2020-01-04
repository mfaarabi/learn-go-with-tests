package main

import (
	"net/http"
	"os"

	"github.com/mfaarabi/learn-go-with-tests/greet"
)

const theString = "Wendy\n"
const port = ":5000"

func HTTPGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, theString)
}

func main() {
	greet.Greet(os.Stdout, theString)
	http.ListenAndServe(port, http.HandlerFunc(HTTPGreeterHandler))
}
