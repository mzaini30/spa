package main

import "net/http"

// import "fmt"
import spa "github.com/roberthodgen/spa-server"
import "runtime"
import "os/exec"

// import "github.com/pkg/browser"

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	open("http://localhost:2020")
	// fmt.Print("Silahkan buka localhost:2020")
	// browser.OpenUrl('http://localhost:2020')
	http.ListenAndServe(":2020", spa.SpaHandler(".", "index.html"))
}
