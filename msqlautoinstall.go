package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

const ShellToUse = "sh"                               
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func main() {

  //fmt.Println("mysql_secure_installation <<EOF\n1\ny\nPASSWORD\nPASSWORD\ny\nn\ny\ny\nEOF\n")
	err, outsec1, errout := Shellout("/usr/local/bin/mysql_secure_installation <<EOF\n\ny\nPASSWORD\nPASSWORD\ny\nn\ny\ny\nEOF\n")
	fmt.Println(outsec1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Println(errout)
	}

	fmt.Println(outsec1)
	fmt.Println(">>>> debug4 mysq oto install bitti>>>")
  }
