package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	initialStamina = 80
	initialAttack  = 5
	initialDefence = 10
)

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику, равный %d.", charName, randint(8, 10))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику, равный %d.", charName, randint(10, 15))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику, равный %d.", charName, randint(2, 4))
	default:
		return "неизвестный класс персонажа"
	}
}

func defence(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", charName, randint(15, 20))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", charName, randint(8, 12))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", charName, randint(12, 15))
	default:
		return "неизвестный класс персонажа"
	}
}

func special(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, initialStamina+25)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, initialAttack+40)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, initialDefence+30)
	default:
		return "неизвестный класс персонажа"
	}
}

func startTraining(charName, charClass string) string {
	switch charClass {
	case "warrior":
		fmt.Printf("%s, ты Воитель — отличный боец ближнего боя.\n", charName)
	case "mage":
		fmt.Printf("%s, ты Маг — превосходный укротитель стихий.\n", charName)
	case "healer":
		fmt.Printf("%s, ты Лекарь — чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)
		cmd = strings.ToLower(cmd)
		switch cmd {
		case "attack":
			fmt.Println(attack(charName, charClass))
		case "defence":
			fmt.Println(defence(charName, charClass))
		case "special":
			fmt.Println(special(charName, charClass))
		case "skip":
			return "тренировка окончена"
		default:
			fmt.Println("неизвестная команда")
		}
	}
}

func choiceCharClass() string {
	var approveChoice, charClass string

	for strings.ToLower(approveChoice) != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		switch charClass {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			fmt.Println("неизвестный класс персонажа")
			continue
		}
		fmt.Print("Введи (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
	}
	return charClass
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Printf("Сейчас твоя выносливость — %d, атака — %d и защита — %d.\n", initialStamina, initialAttack, initialDefence)
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := choiceCharClass()

	fmt.Println(startTraining(charName, charClass))
}
