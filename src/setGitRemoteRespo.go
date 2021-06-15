//https://github.com/turf96/GoLearning.git
//https://gitee.com/turfele/go-learning.git

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var option string
	var out *exec.Cmd
	var repositoryAddr string
	if len(os.Args) > 1 {
		option = os.Args[1]
	}

	switch option {
	case "github":
		repositoryAddr = "\"https://github.com/turf96/GoLearning.git\""

	case "gitee":
		repositoryAddr = "\"https://gitee.com/turfele/go-learning.git\""
	default:
		fmt.Println("ERROR, ./setGitRemoteRespo [ github | gitee ]")
		os.Exit(0)
	}

	command := "git remote set-url origin " + reposigittoryAddr

	fmt.Println(command)

	out = exec.Command("/bin/bash", "-c",command)
	_, err := out.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("ok")
	}

	gitConfig, err := exec.Command("/bin/bash", "-c", "git config --list").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(gitConfig))
	}

}