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
	ClientID int
	Name     string
	Surname  string
}

type Cart struct {
	ClientID  int
	ProductID int
	Quantity  int
}

var products = make([]Product, 0)
var customers = make([]Customer, 0)

// var orders = make([]Order, 0)
var carts = make(map[int]Cart)

// читачь строки
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// Управління товарами:

// Перевірка чи існує товар
func productsExists(name, desc string) bool {
	for _, p := range products {
		if p.Name == name && p.Description == desc {
			return true
		}
	}
	return false
}

// Додавання нових товарів до каталогу
func (p *Product) addProducts() bool {
	if !productsExists(p.Name, p.Description) {
		p.ID = len(products) + 1
		p.IsActive = true
		products = append(products, *p)
		return true
	}
	return false
}

// Видалення товарів з каталогу
func (p Product) deleteProducts() bool {
	for i, prod := range products {
		if prod.ID == p.ID {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}

// Оновлення ціни та кількості товару
func (p *Product) UpdatePriceStock(newPrice float64, newStock int) {
	p.Price = newPrice
	p.Stock = newStock
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

// Пошук товарів за ID
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

// Управління клієнтами

// Перевірка чи існує кліент
func clientExists(name string) bool {
	for _, c := range customers {
		if c.Name == name {
			return true
		}
	}
	return false
}

// Реєстрація нових клієнтів
func (c *Customer) registerClient() bool {
	if !clientExists(c.Name) {
		c.ClientID = len(customers) + 1
		customers = append(customers, *c)
		return true
	}
	return false
}

// Перегляд інформації про клієнта
func checkClientInfo(name string) {
	if clientExists(name) {
		for _, c := range customers {
			if c.Name == name {
				fmt.Printf("ID: %d\nПрізвище: %s\nІм'я: %s\n", c.ClientID, c.Surname, c.Name)
			}
		}
	} else {
		fmt.Printf("Такого Кліента не існує ❌\n", name)
	}
}

// Пошук индексу клієнта по имені
func findCustomerIndex(name string) int {
	for i, c := range customers {
		if c.Name == name {
			return i
		}
	}
	return -1
}

// Оновлення контактних даних клієнта
func (c *Customer) updateClient(change int, newValue string) {
	index := findCustomerIndex(c.Name)
	if index == -1 {
		fmt.Printf("Такого Кліента не існує ❌\n", c.Name)
		return
	}

	switch change {
	case 1:
		customers[index].Name = newValue
		fmt.Println("Ім'я Успішно Оновлене")
	case 2:
		customers[index].Surname = newValue
		fmt.Println("Прізвище Успішно Оновлене")
	default:
		fmt.Println("Невірний параметр зміни ❌")
	}
}

// Система кошика:

// Додавання товарів до кошика
func (c *Cart) addCarts() bool {
	found := false
	for _, p := range products {
		if p.ID == c.ProductID {
			found = true
			break
		}
	}
	if !found {
		return false
	}

	carts[c.ProductID] = *c
	return true
}

// Видалення товарів з кошика
func (c Cart) deleteProductFromCart() bool {
	for i, item := range carts {
		if item.ProductID == c.ProductID {
			delete(carts, i)
			return true
		}
	}
	return false
}

// Перегляд вмісту кошика
// переделать функцию
func CheckCartItem(name string) {
	var client *Customer
	for _, c := range customers {
		if c.Name == name {
			client = &c
			break
		}
	}

	if client == nil {
		fmt.Println("Клієнта не знайдено.")
		return
	}

	fmt.Printf("Клієнт: %s %s\n", client.Name, client.Surname)

	found := false

	for _, cart := range carts {
		if cart.ClientID == client.ClientID {
			for i, prod := range products {
				if prod.ID == cart.ProductID {
					fmt.Printf("%d. %s x%d - %.2f грн\n", i+1, prod.Name, cart.Quantity, prod.Price*float64(cart.Quantity))
					found = true
					break
				}
			}
		}
	}
	if !found {
		fmt.Println("Кошик порожній.")
		return
	}

	totalsum := calculateCartTotal(client.ClientID)
	fmt.Println("Знижка: 0%") //добавить потом знижку глобальную переменую и функцию для расчета
	fmt.Printf("Загальна сума: %.2f грн\n", totalsum)
}

// Підрахунок загальної суми в кошику
func calculateCartTotal(ClientID int) float64 {
	totalsum := 0.0
	for _, cart := range carts {
		if cart.ClientID == ClientID {
			for _, prod := range products {
				if prod.ID == cart.ProductID {
					totalsum += prod.Price * float64(cart.Quantity)
					break
				}
			}
		}
	}
	return totalsum
}

// Розрахунок загальної суми з урахуванням знижок
func calculateTotalSum(cli int, confirm bool) {
	del := 100.0
	totalSum := calculateCartTotal(cli)

	fmt.Printf("Вартість доставки: %.2f грн\n", del)
	fmt.Printf("Загальна сума до сплати: %.2f грн\n", totalSum+del)

	// сделать створення замовлення для кошика в Система замовлень і добавити сюди
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
	p := Product{
		Name:        "RTX 3060",
		Description: "GYGABYTE RTX 3060",
		Price:       8000,
		Category:    "Пк Комплектуючі",
		Stock:       4,
	}
	p2 := Product{
		Name:        "RTX 4060",
		Description: "GYGABYTE RTX 4060",
		Price:       12000,
		Category:    "Пк Комплектуючі",
		Stock:       2,
	}
	c := Customer{
		Name:    "Jotaro",
		Surname: "Kujo",
	}
	k := Cart{
		ClientID:  c.ClientID + 1,
		ProductID: 1,
		Quantity:  2,
	}
	k2 := Cart{
		ClientID:  c.ClientID + 1,
		ProductID: 2,
		Quantity:  1,
	}
	p.addProducts()
	p2.addProducts()
	c.registerClient()
	k.addCarts()
	k2.addCarts()

	fmt.Println(products)
	fmt.Println(customers)
	fmt.Println(carts)

	//k2.deleteProductFromCart()
	fmt.Println(carts)
	CheckCartItem("Jotaro")
	CheckCartItem("Зщзф")

	calculateTotalSum(1, true)
}
