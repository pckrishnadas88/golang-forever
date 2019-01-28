package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	flag.Parse()
	lenOfCmdArgs := len(flag.Args())
	if lenOfCmdArgs == 0 {
		log.Fatal("Please use start or stop with commands.")
	}
	/**
	* For starting a command pass `start` as the first cmd argument.
	 */
start:
	if strings.Join(flag.Args()[:1], " ") == "start" {
		//arg1, arg2 := "node", "../web-test/index.js" // now accept as cmd args.
		arg1, arg2 := strings.Join(flag.Args()[1:2], " "), strings.Join(flag.Args()[2:], " ")
		cmd := exec.Command(arg1, arg2)
		stdout, err := cmd.StdoutPipe() // For getting the log of the external command.
		scanner := bufio.NewScanner(stdout)
		go func() {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()
		if err != nil {
			log.Fatal(err)
		}
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Waiting for command to finish...")
		log.Printf("Process id is %v", cmd.Process.Pid)
		err = cmd.Wait()
		log.Printf("Command finished with error, now restarting: %v", err)
		goto start
	}
	// TODO : Stop a run.
	// if strings.Join(flag.Args()[:1], " ") == "stop" {
	// 	syscall.Kill(-pid, syscall.SIGKILL)
	// }
}
