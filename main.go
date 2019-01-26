package main

import (
	"flag"
	"log"
	"os/exec"
	"strings"
)

func main() {
start:
	flag.Parse()
	//arg1, arg2 := "node", "../web-test/index.js" // now accept as cmd args.
	arg1, arg2 := strings.Join(flag.Args()[:1], " "), strings.Join(flag.Args()[1:], " ")
	cmd := exec.Command(arg1, arg2)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	log.Printf("Process id is %v", cmd.Process.Pid)
	err = cmd.Wait()
	log.Printf("Command finished with error, now restarting: %v", err)
	goto start
}
