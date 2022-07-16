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
	http.ListenAndServe("0.0.0.0:4001", nil)
}

func (l *Listen) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	var out string
	out, _ = Cmd("git", []string{"pull", "origin", "main"})
	fmt.Println(out)
	out, _ = Cmd("docker", []string{"cp", "/root/source/_data", "blog:/home/hexo/blog/source"})
	fmt.Println(out)
	out, _ = Cmd("docker", []string{"cp", "/root/source/_posts", "blog:/home/hexo/blog/source"})
	fmt.Println(out)
	out, _ = Cmd("docker", []string{"cp", "/root/source/link", "blog:/home/hexo/blog/source"})
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
