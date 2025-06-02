package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Stock       int
	IsActive    bool
}

type Customer struct {
	Name     string
	Surname  string
	Password string
}

var products = make([]Product, 0)
var customers = make([]Customer, 0)

// var orders = make([]Order, 0)
// var carts = make(map[int]Cart)

// читачь строки
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// Управління клієнтами

// Реєстрація нових клієнтів
func registerClient(name, surname, password string) bool {
	if !clientExists(name) {

		client := Customer{
			Name:     name,
			Surname:  surname,
			Password: password,
		}

		customers = append(customers, client)
		return true
	}
	return false
}

// Перевірка чи існує кліент
func clientExists(name string) bool {
	for _, c := range customers {
		if c.Name == name {
			return true
		}
	}
	return false
}

// Перегляд інформації про клієнта
func checkClientInfo(name string) {
	if clientExists(name) {
		for _, c := range customers {
			if c.Name == name {
				fmt.Printf("Прізвище: %s\nІм'я: %s\nПароль: %s\n", c.Surname, c.Name, c.Password)
			}
		}
	} else {
		fmt.Printf("Неможливо відобразити інформацію про клієнта: %s .Такого Кліента не існує\n", name)
	}
}

// Оновлення контактних даних клієнта
func updateClient(name string, change int, newName string) {
	if clientExists(name) {
		if change == 1 {
			for i, c := range customers {
				if c.Name == name {
					customers[i].Name = newName
					break
				}
			}
			fmt.Println("Ім'я Успішно Оновлене")
		} else if change == 2 {
			for i, c := range customers {
				if c.Name == name {
					customers[i].Surname = newName
					break
				}
			}
			fmt.Println("Прізвище Успішно Оновлене")
		}

	}
}

// Отримує текстове введення
func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	input, _ = reader.ReadString('\n')
	return input
}

// Отримує числове введення
func getIntInput(prompt string) int {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	//перевірка на пробіли
	if strings.ContainsAny(input, " \t") {
		fmt.Println("Некоректне введення. Ведіть ціле число без пробілів. ❌")
		return -1
	}

	//перевірка на введеня числа
	var value int
	_, err := fmt.Sscanf(input, "%d", &value)
	if err != nil {
		fmt.Println("Некоректне введення. Ведіть ціле число. ❌")
		return -1
	}
	return value
}

func main() {

}
