package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	name    string
	class   string
	health  int
	attack  int
	defence int
	special int
}

func generate_player(name string, class string) player {
	var hp, atk, def, sp int

	if class == "knight" {
		hp = 100
		atk = 15
		def = 20
		sp = 30
	} else if class == "archer" {
		hp = 75
		atk = 25
		def = 10
		sp = 50
	} else if class == "sorcerer" {
		hp = 60
		atk = 35
		def = 5
		sp = 70
	} else {
		fmt.Println("Unknown class, setting default values")
		hp = 80
		atk = 10
		def = 10
		sp = 20
	}
	return player{name, class, hp, atk, def, sp}
}

func attack_enemy(p player) string {
	rand.Seed(time.Now().UnixNano())
	damage := p.attack + rand.Intn(10)
	return fmt.Sprintf("%s атакует врага и наносит %d урона!", p.name, damage)
}

func defend(p player) string {
	blocked := p.defence / 2
	return fmt.Sprintf("%s защищается, блокируя %d урона.", p.name, blocked)
}

func special_ability(p player) string {
	if p.class == "knight" {
		return fmt.Sprintf("%s использует 'Рыцарский рывок'! Его защита увеличена на %d.", p.name, p.special)
	} else {
		if p.class == "archer" {
			return fmt.Sprintf("%s использует 'Меткий выстрел'! Атака увеличена на %d.", p.name, p.special)
		} else {
			if p.class == "sorcerer" {
				return fmt.Sprintf("%s использует 'Огненный шар'! Магическая атака на %d.", p.name, p.special)
			} else {
				return "Способность неизвестна"
			}
		}
	}
}

func train_player(p player) {
	fmt.Println("Добро пожаловать на тренировку,", p.name)
	fmt.Println("Ты выбрал класс:", p.class)
	fmt.Println("Твои характеристики: Здоровье:", p.health, "Атака:", p.attack, "Защита:", p.defence)

	fmt.Println("Доступные команды: attack, defend, special, exit")

	var command string
	for command != "exit" {
		fmt.Print("Введите команду: ")
		fmt.Scanln(&command)

		if command == "attack" {
			fmt.Println(attack_enemy(p))
		} else if command == "defend" {
			fmt.Println(defend(p))
		} else if command == "special" {
			fmt.Println(special_ability(p))
		} else if command == "exit" {
			break
		} else {
			fmt.Println("Неизвестная команда, попробуйте снова")
		}
	}
}

func choose_class() string {
	var chosen_class string
	fmt.Println("Выберите класс персонажа: knight, archer, sorcerer")
	fmt.Scanln(&chosen_class)

	if chosen_class == "knight" {
		fmt.Println("Вы выбрали рыцаря! У него высокая защита и крепкое здоровье.")
	} else if chosen_class == "archer" {
		fmt.Println("Вы выбрали лучника! Он наносит мощные атаки на расстоянии.")
	} else if chosen_class == "sorcerer" {
		fmt.Println("Вы выбрали мага! Магические атаки наносят огромный урон.")
	} else {
		fmt.Println("Неизвестный класс, попробуйте снова.")
		return choose_class()
	}
	return chosen_class
}

func main() {
	fmt.Println("Добро пожаловать в мир приключений!")
	fmt.Print("Введите имя своего персонажа: ")
	var player_name string
	fmt.Scanln(&player_name)

	class := choose_class()
	hero := generate_player(player_name, class)
	train_player(hero)
}
