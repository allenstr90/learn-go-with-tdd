package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world!!!")
}

func main() {
	Greet(os.Stdout, "Allen")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler)))
}
