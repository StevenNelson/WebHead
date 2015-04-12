/*
 Begin by opening a browser, and point it at localhost.
 Then, very quickly, start an http listener.
*/

package webhead

import (
	"net/http"
	"os/exec"
)

func Start() error {

	// Start opening the browser, pointed at the application
	err := openUI()
	if err != nil {
		return err
	}

	// The webserver has to run on the main go thread, or the program will
	// reach the end and quit. Nothing else will happen until this func returns
	err = startServer()
	return err

}

func startServer() error {
	err := http.ListenAndServe(":12415", nil)
	return err
}

func openUI() error {
	var cmd = exec.Command("open", "http://localhost:12415/")
	err := cmd.Run()
	return err
}
