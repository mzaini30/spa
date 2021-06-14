package main
import "net/http"
import "fmt"
import spa "github.com/roberthodgen/spa-server"

func main() {
	fmt.Print("Silahkan buka localhost:2020")
    http.ListenAndServe(":2020", spa.SpaHandler(".", "index.html"))
}
