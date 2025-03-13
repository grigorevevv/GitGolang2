package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name           string
	Class          string
	Health         int
	AttackPower    int
	DefencePower   int
	SpecialAbility int
}

func GeneratePlayer(name string, class string) Player {
	var health, attack, defence, special int

	switch {
	case class == "knight":
		health, attack, defence, special = 100, 15, 20, 30
	case class == "archer":
		health, attack, defence, special = 75, 25, 10, 50
	case class == "sorcerer":
		health, attack, defence, special = 60, 35, 5, 70
	default:
		fmt.Println("Unknown class, setting default values")
		health, attack, defence, special = 80, 10, 10, 20
	}
	return Player{name, class, health, attack, defence, special}
}

func AttackEnemy(p Player) string {
	damage := p.AttackPower + rand.Intn(10)
	return fmt.Sprintf("%s атакует врага и наносит %d урона!", p.Name, damage)
}

func Defend(p Player) string {
	blocked := p.DefencePower / 2
	return fmt.Sprintf("%s защищается, блокируя %d урона.", p.Name, blocked)
}

func SpecialAbility(p Player) string {
	if p.Class == "knight" {
		return fmt.Sprintf("%s использует 'Рыцарский рывок'! Его защита увеличена на %d.", p.Name, p.SpecialAbility)
	}
	if p.Class == "archer" {
		return fmt.Sprintf("%s использует 'Меткий выстрел'! Атака увеличена на %d.", p.Name, p.SpecialAbility)
	}
	if p.Class == "sorcerer" {
		return fmt.Sprintf("%s использует 'Огненный шар'! Магическая атака на %d.", p.Name, p.SpecialAbility)
	}
	return "Способность неизвестна"
}

func Training(p Player) {
	fmt.Println("Добро пожаловать на тренировку,", p.Name)
	fmt.Println("Ты выбрал класс:", p.Class)
	fmt.Println("Твои характеристики: Здоровье:", p.Health, "Атака:", p.AttackPower, "Защита:", p.DefencePower)
	fmt.Println("Доступные команды: attack, defend, special, exit")

	var command string
	for command != "exit" {
		fmt.Print("Введите команду: ")
		_, err := fmt.Scanln(&command)
		if err != nil {
			fmt.Println(err)
		}
		switch {
		case command == "attack":
			fmt.Println(AttackEnemy(p))
		case command == "defend":
			fmt.Println(Defend(p))
		case command == "special":
			fmt.Println(SpecialAbility(p))
		case command == "exit":
			break
		default:
			fmt.Println("Неизвестная команда, попробуйте снова")
		}
	}
}

func ChooseClass() string {
	var chosenClass string

	for {
		fmt.Println("Выберите класс персонажа: knight, archer, sorcerer")
		_, err := fmt.Scanln(&chosenClass)
		if err != nil {
			fmt.Println(err)
		}

		if chosenClass == "knight" {
			fmt.Println("Вы выбрали рыцаря! У него высокая защита и крепкое здоровье.")
			return chosenClass
		} else if chosenClass == "archer" {
			fmt.Println("Вы выбрали лучника! Он наносит мощные атаки на расстоянии.")
			return chosenClass
		} else if chosenClass == "sorcerer" {
			fmt.Println("Вы выбрали мага! Магические атаки наносят огромный урон.")
			return chosenClass
		} else {
			fmt.Println("Неизвестный класс, попробуйте снова.")
		}
	}
}

func main() {
	fmt.Println("Добро пожаловать в мир приключений!")
	fmt.Print("Введите имя своего персонажа: ")
	var player string
	_, err := fmt.Scanln(&player)
	if err != nil {
		fmt.Println(err)
	}

	class := ChooseClass()
	hero := GeneratePlayer(player, class)
	Training(hero)
}
