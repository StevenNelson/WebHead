/*
 Begin by opening a browser, and point it at localhost.
 Then, very quickly, start an http listener.
*/

package webhead

import (
	"net/http"
	"os/exec"
)

func Start(port string) error {

	// Start opening the browser, pointed at the application
	err := openUI(port)
	if err != nil {
		return err
	}

	// The webserver has to run on the main go thread, or the program will
	// reach the end and quit. Nothing else will happen until this func returns
	err = startServer(port)
	return err

}

func startServer(port string) error {
	err := http.ListenAndServe(":"+port, nil)
	return err
}

func openUI(port string) error {
	var cmd = exec.Command("open", "http://localhost:"+port+"/")
	err := cmd.Run()
	return err
}
