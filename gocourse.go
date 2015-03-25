package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// prepare flag parsers
	var example = flag.Int("ex", 0, "example name to run")
	flag.Parse()

	examples := []string{"c01calculator", "c02calculator", "c03calculator",
		"c04calculator", "c05interview", "c06interview"}

	var err error
	if *example > 0 && *example <= len(examples) {
		err = compileAndRunExample(examples[*example-1])
	} else {
		flag.PrintDefaults()
	}

	if err != nil {
		fmt.Printf("An error occured - %s", err)
	}
}

func compileAndRunExample(packageName string) (err error) {
	cmd := exec.Command("go", "install", "github.com/erezak/gocourse/"+packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return
	}
	cmd = exec.Command(packageName)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return
}
