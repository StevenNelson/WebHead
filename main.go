/*
 Begin by
*/

package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {

	// Start opening the browser, pointed at the application
	openUI()

	// The webserver has to run on the main go thread, or the program will
	// reach the end and quit. Nothing else will happen until this func returns
	startServer()

}

func startServer() {
	http.HandleFunc("/", uiCall)
	err := http.ListenAndServe(":12415", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func uiCall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from your favorite application")
}

func openUI() {
	time.Sleep(1 * time.Second)
	var cmd = exec.Command("open", "http://localhost:12415/")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Growing a head")
	}
}
