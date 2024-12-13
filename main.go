package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(CharName, CharClass string) string {
	if CharClass == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", CharName, 5+randint(3, 5))
	}

	if CharClass == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", CharName, 5+randint(5, 10))
	}

	if CharClass == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", CharName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(CharName, CharClass string) string {
	if CharClass == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", CharName, 10+randint(5, 10))
	} else if CharClass == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", CharName, 10+randint(-2, 2))
	} else if CharClass == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", CharName, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"

}

// обратите внимание на "if else" и на "else"
func special(CharName, CharClass string) string {
	if CharClass == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", CharName, 80+25)
	} else if CharClass == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", CharName, 5+40)
	} else if CharClass == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", CharName, 10+30)
	}
	return "неизвестный класс персонажа"

}

// здесь обратите внимание на имена параметров
func StartTraining(CharName, CharClass string) string {
	if CharClass == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", CharName)
	}

	if CharClass == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", CharName)
	}

	if CharClass == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", CharName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(CharName, CharClass))
		}

		if cmd == "defence" {
			fmt.Println(defence(CharName, CharClass))
		}

		if cmd == "special" {
			fmt.Println(special(CharName, CharClass))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func ChoiseCharClass() string {
	var ApproveChoice string
	var CharClass string

	for ApproveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &CharClass)
		if CharClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if CharClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if CharClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &ApproveChoice)
		ApproveChoice = strings.ToLower(ApproveChoice)
	}
	return CharClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var CharName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &CharName)

	fmt.Printf("Здравствуй, %s\n", CharName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	CharClass := ChoiseCharClass()

	fmt.Println(StartTraining(CharName, CharClass))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
