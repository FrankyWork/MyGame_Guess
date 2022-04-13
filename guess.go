// Игра угадай рандомно сгенерированное число
package main

//Пакеты
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Рандомайзер
	second := time.Now().Unix()
	rand.Seed(second)
	target := rand.Intn(100) + 1
	fmt.Println("Угадай мое число от 0 до 100")
	fmt.Println("Ты готов угадывать??")

	reader := bufio.NewReader(os.Stdin)

	//Переменная при проигрыше
	success := false

	//Цикл попыток ввода
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У тебя", 10-guesses, "попыток угадать")

		//Ввод вариантов ответа
		fmt.Print("Введи число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		//Условия
		if guess < target {
			fmt.Println("Ойй. Твое число меньше загаданного.")
		} else if guess > target {
			fmt.Println("Ойй. Твое число больше загаданного.")
		} else {
			fmt.Println("Good Job! Ты угадал мое число!")
			break
		}
	}
	//Вывод при проигрыше
	if !success {
		fmt.Println("Извини, ты не угадал. Это было число: ", target)
	}
}
