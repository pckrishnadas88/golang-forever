package main

import (
	"log"
	"os/exec"
)

func main() {
start:
	arg1, arg2 := "node", "../web-test/index.js" // TODO : accept as cmd args.
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
