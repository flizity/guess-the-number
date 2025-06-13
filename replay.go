package main

import (
	"bufio"
	"strings"

	"github.com/fatih/color"
)

func AskForReplay(reader *bufio.Reader) bool {
	c := color.New(color.FgYellow)
	c.Print("\nХотите сыграть еще раз? (да/нет): ")

	answer, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(answer)) == "да"
}
