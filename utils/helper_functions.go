package utils

import (
	"bufio"
	"os"
	"runtime"
	"os/exec"
	"strconv"
	"strings"
	"fmt"
	
)

func Readline() string {
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	return text
}


func Readint() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	number, err := strconv.Atoi(text)
	if err != nil {
		// Handle the error as needed, for now we return 0
		return 0
	}
	return number
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Pause() {
	fmt.Println("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}