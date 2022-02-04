package main

import (
	"fmt"
	spa "github.com/roberthodgen/spa-server"
	"math/rand"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

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
	rand.Seed(time.Now().UTC().UnixNano())
	angka := strconv.Itoa(rand.Int())
	angka = strings.Replace(angka, "0", "1", -1)
	port := angka[0:4]

	fmt.Println("Sedang membuka localhost:" + port + "....")
	open("http://localhost:" + port)
	http.ListenAndServe(":"+port, spa.SpaHandler(".", "index.html"))
}
