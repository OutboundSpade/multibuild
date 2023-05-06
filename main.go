package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var command string
	flag.StringVar(&command, "command", "ls", "command to run")
	flag.Parse()
	//for each directory in the current directory
	//run the command
	var channels []chan bool
	for index, directory := range getDirectories() {
		channels = append(channels, make(chan bool))
		go runCommand(command, directory, channels[index])
	}

	//for each channel in channels
	for _, channel := range channels {
		<-channel
	}
}

func getDirectories() []string {
	//get the current directory
	//get all directories in the current directory
	//return the directories
	wd, err := os.Getwd()
	must(err)
	entries, err := os.ReadDir(wd)
	must(err)
	var directories []string
	for _, entry := range entries {
		if entry.IsDir() {
			//join the current directory with the directory name
			//append the directory to the directories
			if strings.HasPrefix(wd, "/") {
				directories = append(directories, entry.Name())
				continue
			}
			directories = append(directories, wd+"/"+entry.Name())
		}
	}
	return directories
}

func runCommand(command string, directory string, isDone chan<- bool) {
	//change to the directory
	//run the command
	//return true
	log(directory, "executing command: "+command)

	//split the command into the command and the arguments
	//run the command with the arguments
	cmd := exec.Command("bash", "-c", "cd "+directory+"&&"+command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log(directory, "error: "+err.Error())
	}
	isDone <- true
	close(isDone)
}
func log(dir string, message string) {
	//log the message with the directory
	fmt.Printf("%s: %s\n", dir, message)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
