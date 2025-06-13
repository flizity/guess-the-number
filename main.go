package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	reader := bufio.NewReader(os.Stdin)

	for {
		PlayGame(r, 1, 100, reader)

		if !AskForReplay(reader) {
			color.Blue("Спасибо за игру! До свидания!")
			break
		}
	}
}

func PlayGame(r *rand.Rand, min, max int, reader *bufio.Reader) {
	number := r.Intn(max-min+1) + min
	attempts := make([]int, 0, 8)
	cBlue := color.New(color.FgBlue)
	cRed := color.New(color.FgRed)

	fmt.Print("Хотите начать? (да/нет): ")
	if answer, _ := reader.ReadString('\n'); strings.TrimSpace(answer) != "да" {
		cRed.Println("До свидания!")
		return
	}

	fmt.Printf("Загадываю число от %d до %d. У тебя 8 попыток!\n", min, max)

	for i := 1; i <= 8; i++ {
		fmt.Printf("Попытка %d/8: ", i)

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Введите число!")
			i--
			continue
		}

		if choice < min || choice > max {
			fmt.Printf("Число должно быть между %d и %d!\n", min, max)
			i--
			continue
		}

		attempts = append(attempts, choice)

		switch {
		case choice == number:
			cBlue.Printf("\nПоздравляю! Ты угадал число %d за %d попыток!\n", number, len(attempts))
			fmt.Println("Твои попытки:", attempts)
			return
		case choice < number:
			fmt.Println("Загаданное число больше")
		default:
			fmt.Println("Загаданное число меньше")
		}
	}

	cRed.Printf("\nУвы, число было %d. Твои попытки: %v\n", number, attempts)
}
