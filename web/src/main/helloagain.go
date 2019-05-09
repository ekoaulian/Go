package main
import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("hello"))
}
}

func main() {
// ...
http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("hello again"))
})
}

func main() {
// ...
http.HandleFunc("/", handlerIndex)
http.HandleFunc("/index", handlerIndex)	
fmt.Println("server started at localhost:9000")
http.ListenAndServe(":9000", nil)
}