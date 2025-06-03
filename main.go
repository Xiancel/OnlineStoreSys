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
	Price       float64
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

// Управління товарами:
// Додавання нових товарів до каталогу
func addProducts(name, desc, category string, price float64, stock int) bool {
	if !productsExists(name, desc) {
		product := Product{
			ID:          len(products) + 1,
			Name:        name,
			Description: desc,
			Price:       price,
			Category:    category,
			Stock:       stock,
			IsActive:    true,
		}

		products = append(products, product)
		return true
	}
	return false
}

// Перевірка чи існує товар
func productsExists(name, desc string) bool {
	for _, p := range products {
		if p.Name == name && p.Description == desc {
			return true
		}
	}
	return false
}

// Видалення товарів з каталогу
func deleteProducts(id int) bool {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}

// Пошук товарів за назвою
func searchProductByName(name string) {
	found := false
	for _, p := range products {
		if p.Name == name {
			if !found {
				fmt.Printf("Товар %s Успішно Знайдений.\nВсі товари з токою ж назвою:\n", name)
				found = true
			}
			fmt.Printf("ID: %d | %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Description, p.Price, p.Stock)
		}
	}
	if !found {
		fmt.Printf("Товар %s не знайдено!\n", name)
	}
}

// Пошук товарів за назвою
func searchProductById(id int) {
	for _, p := range products {
		if p.ID == id {
			fmt.Println("Товар Успішно Знайдений.")
			fmt.Printf("ID: %d | %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Description, p.Price, p.Stock)
			return
		}
	}
	fmt.Println("Товар не знайдено!")
}

// Відображення всіх товарів
func displayAllProducts() {
	for _, p := range products {
		fmt.Printf("ID: %d | %s | Категорія: %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Category, p.Description, p.Price, p.Stock)
	}
}

// Відображення всіх товарів з наявністю
func displayAllProductsStock() {
	found := false
	for _, p := range products {
		if p.Stock > 0 {
			if !found {
				found = true
			}
			fmt.Printf("ID: %d | %s | Категорія: %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Category, p.Description, p.Price, p.Stock)
		}
	}
	if !found {
		fmt.Println("Немає товарів в наявності")
	}
}

// Оновлення ціни та кількості товару
func UpdatePriceStock(id int, newPrice float64, newStock int) {
	for i, p := range products {
		if p.ID == id {
			products[i].Price = newPrice
			products[i].Stock = newStock
			return
		}
	}

}

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
		fmt.Printf("Неможливо відобразити інформацію про клієнта: %s .Такого Кліента не існує ❌\n", name)
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
	} else {
		fmt.Printf("Неможливо оновити данні о клієнті %s.Такого Кліента не існує ❌\n", name)
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
	//перевірка
	addProducts("RTX 4060TI", "GYGABYTE GeForce RTX 4060TI", "Пк Комплектуючі", 20000, 5)
	addProducts("RTX 5060TI", "GYGABYTE GeForce RTX 5060TI", "Пк Комплектуючі", 30000, 5)
	addProducts("RTX 3060TI", "GYGABYTE GeForce RTX 3060TI", "Пк Комплектуючі", 12000, 0)
	fmt.Println("display all products")
	displayAllProducts()

	deleteProducts(2)

	fmt.Println("\nsearch products dy id")
	searchProductById(1)

	fmt.Println("\nsearch products by name")
	searchProductByName("RTX 5060TI")

	fmt.Println("\ndisplay all products")
	displayAllProducts()

	fmt.Println("\ndisplay all products stocks")
	displayAllProductsStock()

	fmt.Println("\nUpdate and Display")
	UpdatePriceStock(1, 19500, 3)
	displayAllProducts()
}
