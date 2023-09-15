package main

import (
	"fmt"
	"groupie-tracker-geolocalization/pkg/funcs"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	// Parsing JSON
	funcs.ParseJson()

	mux := http.NewServeMux()

	// Serving static files
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))

	// Handlers
	mux.HandleFunc("/", funcs.HomeHandler)
	mux.HandleFunc("/group/", funcs.GroupHandler)
	mux.HandleFunc("/search", funcs.SearchHandler)


	// Initiating port
	launchserver("http://localhost:8000")
	log.Println("Server is launched on " + "http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func launchserver(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
