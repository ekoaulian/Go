package main
import "fmt"
import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
var message = "Welcome halaman index"
w.Write([]byte(message))
}
func handlerHello(w http.ResponseWriter, r *http.Request) {
var message = "Hello world!"
w.Write([]byte(message))
}
func handlerEko(w http.ResponseWriter, r *http.Request) {
var message = "Ini Halaman Eko!"
w.Write([]byte(message))
}

func main() {
http.HandleFunc("/", handlerIndex)
http.HandleFunc("/index", handlerIndex)
http.HandleFunc("/hello", handlerHello)
http.HandleFunc("/eko", handlerEko)
var address = "localhost:9000"
fmt.Printf("server started at %s\n", address)
err := http.ListenAndServe(address, nil)
if err != nil {
fmt.Println(err.Error())
}
}