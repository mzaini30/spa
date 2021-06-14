package main
import "net/http"
import spa "github.com/roberthodgen/spa-server"

func main() {
    http.ListenAndServe(":8080", spa.SpaHandler(".", "index.html"))
}
