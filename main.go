package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	min, max := 1, 100

	number := r.Intn((max - min + 1) + min)
	count := 8

	attempts := make([]int, 0, count)

	fmt.Println("хотите начать?(да/нет)")

	var answer string
	fmt.Scan(&answer)

	if answer == "да" {
		fmt.Printf("загадываю число от %d дот %d, угадай верное за %d попыток \n", min, max, count)
	}
	if answer != "да" {
		fmt.Println("иди нахуй тогда")
		return
	}

	for i := 1; i <= count; i++ {
		fmt.Printf("попытка %d/%d: ", i, count)

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("выбери число")
			i--
			continue
		}

		if choice < min {
			fmt.Printf("число не может быть меньше %d! ", min)
			i--
			continue
		}
		if choice > max {
			fmt.Printf("число не может быть больше %d! ", max)
			i--
			continue
		}

		attempts = append(attempts, choice)

		if choice == number {
			fmt.Printf("\nкрасава ебать, ты угадал число %d за %d попыток!\n", number, len(attempts))
			fmt.Println("Твои попытки:", attempts)
			return
		} else if choice < number {
			fmt.Println("это число больше")
		} else {
			fmt.Println("это число меньше")
		}
	}

	fmt.Printf("\nне угадал ебать, число было %d.\n", number)
	fmt.Printf("твои попытки: %v\n", attempts)
}
