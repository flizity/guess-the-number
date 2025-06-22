package main

import (
	"bufio"
	"strings"

	"github.com/fatih/color"
)

func AskForReplay(reader *bufio.Reader) bool {
	c := color.New(color.FgYellow)

	for {
		c.Print("\nХотите сыграть еще раз? (да/нет): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "да", "д":
			return true
		case "нет", "н":
			return false
		default:
			color.Red("Пожалуйста, введите 'да' или 'нет'")
		}
	}
}
