package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type Listen struct {
}

func main() {
	http.Handle("/", &Listen{})
	http.ListenAndServe("127.0.0.1:8001", nil)
}

func (l *Listen) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	out, err := Cmd("git", []string{"pull", "origin", "main"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func Cmd(commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	cmd.Dir, _ = os.Getwd()
	cmd.Dir += "/source"
	fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}
