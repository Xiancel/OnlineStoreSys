package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// структура Товарів
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    string
	Stock       int
	IsActive    bool
}

// структура клієнтів
type Customer struct {
	ClientID int
	Name     string
	Surname  string
}

// структіра кошику
type Cart struct {
	ClientID  int
	ProductID int
	Quantity  int
}

// структура замовлень
type Order struct {
	OrdersID   int
	ClientID   int
	Sum        float64
	Items      []Cart
	Status     string
	CreateData time.Time
}

// глобальні змінні які хранять в собі данні а саме:
var products = make([]Product, 0)   // товари
var customers = make([]Customer, 0) // клієнтів
var orders = make([]Order, 0)       // замовлення
var carts = make(map[int]Cart)      // кошик

// глобальні змінні
var del float64 = 100.0                // доставка
var globalDisc float64 = 0.0           // процент знижки
var StoreName string = "SuuupeerStore" // назва магазину

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
	// перевірка на наявність товара
	if !productsExists(p.Name, p.Description) {
		p.ID = len(products) + 1        // присвоення ID товару
		p.IsActive = true               // присвоення активності товару
		products = append(products, *p) // додавання до слайсла
		return true
	}
	return false
}

// Видалення товарів з каталогу
func (p Product) deleteProducts() bool {
	for i, prod := range products {
		if prod.ID == p.ID {
			products = append(products[:i], products[i+1:]...) // видалення товару зі слайсу
			return true
		}
	}
	return false
}

// Оновлення ціни та кількості товару
func (p *Product) UpdatePriceStock(newPrice float64, newStock int) {
	p.Price = newPrice // міняє ціну
	p.Stock = newStock // мінає колічество
}

// Пошук товарів за назвою
func searchProductByName(name string) {
	// змінна для зберегання найдених товарів
	found := false
	for _, p := range products {
		if p.Name == name {
			if !found {
				// виводить ім'я знайденого товару
				fmt.Printf("\nТовар %s Успішно Знайдений.✅\nВсі товари з токою ж назвою:\n", name)
				found = true
			}
			// виводить всі товари з таким же ім'ям
			fmt.Printf("ID: %d | %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Description, p.Price, p.Stock)
		}
	}
	// якщо товар не був знайдений видає повідомлення
	if !found {
		fmt.Printf("Товар %s не знайдено!❌\n", name)
	}
}

// Пошук товарів за ID
func searchProductById(id int) {
	for _, p := range products {
		if p.ID == id {
			// виводить повідомлення про знайдений товар
			fmt.Println("\nТовар Успішно Знайдений.✅")
			// виводить товар і його опис(характеристики)
			fmt.Printf("ID: %d | %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Description, p.Price, p.Stock)
			return
		}
	}
	fmt.Println("Товар не знайдено!❌")
}

// Відображення всіх товарів
func displayAllProducts() {
	for _, p := range products {
		// виводить всі товари в каталозі
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
			// виводить всі товари в каталозі з наявністю
			fmt.Printf("ID: %d | %s | Категорія: %s | Опис: %s | Ціна: %.2f грн | Наявність: %d шт.\n", p.ID, p.Name, p.Category, p.Description, p.Price, p.Stock)
		}
	}
	if !found {
		fmt.Println("Немає товарів в наявності!❌")
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
	// перевірка на інування клієнта
	if !clientExists(c.Name) {
		c.ClientID = len(customers) + 1   // присвоення ID клієнту
		customers = append(customers, *c) // додавання до слайсу
		return true
	}
	return false
}

// Перегляд інформації про клієнта
func checkClientInfo(name string) {
	// перевірка на існування клієнта
	if clientExists(name) {
		for _, c := range customers {
			if c.Name == name {
				// виводить інформацію про клієнта
				fmt.Printf("ID: %d\nІм'я: %s\nПрізвище: %s\n", c.ClientID, c.Name, c.Surname)
			}
		}
	} else {
		fmt.Println("Такого Кліента не існує!❌")
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
	// получення індексу клієнта
	index := findCustomerIndex(c.Name)
	// перевірка на існування клієнта
	if index == -1 {
		fmt.Printf("Такого Кліента не існує ❌\n", c.Name)
		return
	}

	// змінна данних кліента
	switch change {
	case 1:
		customers[index].Name = newValue
		fmt.Println("Ім'я Успішно Оновлене ✅")
	case 2:
		customers[index].Surname = newValue
		fmt.Println("Прізвище Успішно Оновлене ✅")
	default:
		fmt.Println("Невірний параметр зміни ❌")
	}
}

// Система кошика:

// Додавання товарів до кошика
func (c *Cart) addCarts() bool {
	for i, _ := range products {
		// перевірка ID товару
		if products[i].ID == c.ProductID {
			// перевірка на достатность товару
			if c.Quantity > products[i].Stock {
				fmt.Printf("Помилка.Недостатньо товару на складі.Доступно %d шт.❌\n", products[i].Stock)
				return false
			}

			// перевірка на інування товара у кошику
			existing, exists := carts[c.ProductID]
			if exists {
				// обчислює сумарну кількість товару
				totalQty := existing.Quantity + c.Quantity

				// перевірка на перевищення сумарної кількості товару
				if totalQty > products[i].Stock {
					fmt.Printf("Сумарна кількість у кошику превищує доступну.Доступна %d шт.❌\n", products[i].Stock)
					return false
				}

				// оновлення кількості
				c.Quantity = totalQty
			}

			// зменшення кількості товару на складі
			products[i].Stock -= c.Quantity

			// додавання або оновлення товара у кошику
			carts[c.ProductID] = *c
			return true
		}
	}
	return false
}

// Видалення товарів з кошика
func (c Cart) deleteProductFromCart() bool {
	for i, item := range carts {
		// перевірка на інування такого айди товару і клієнта
		if item.ProductID == c.ProductID && item.ClientID == c.ClientID {

			for i, _ := range products {
				// якшо товар з таким айди є у кошику то кількість товару у кошику повертаєтсья на склад
				if products[i].ID == c.ProductID {
					products[i].Stock += item.Quantity
					break
				}
			}
			// видалення товару з кошика
			delete(carts, i)
			return true
		}
	}
	return false
}

// обчислює знижку
func calculateDiscount(sum float64) float64 {
	return sum * globalDisc / 100
}

// Перегляд вмісту кошика
func CheckCartItem(name string) {
	// змінна для храніня кліенту
	var client *Customer

	// пошук клієнта за ім'ям
	for _, c := range customers {
		if c.Name == name {
			client = &c // зберегання вказівник на знайденого клієнта
			break
		}
	}

	// якщо клієнта не знайдено виводится повідомлення
	if client == nil {
		fmt.Println("Клієнта не знайдено.❌")
		return
	}

	// виводить ім'я та прізвище знайденого клієнта
	fmt.Printf("Клієнт: %s %s\n", client.Name, client.Surname)

	// змінна для зберегання знайденого кліента
	found := false

	for _, cart := range carts {
		// перевірка чи в корзині клієнта є такий товар
		if cart.ClientID == client.ClientID {
			// шукає товар у каталозі
			for i, prod := range products {
				if prod.ID == cart.ProductID {
					// виводимо інформацію про товар у кошику
					fmt.Printf("%d. %s x%d - %.2f грн\n", i+1, prod.Name, cart.Quantity, prod.Price*float64(cart.Quantity))
					found = true
					break
				}
			}
		}
	}
	if !found {
		fmt.Println("Кошик порожній.❌")
		return
	}

	// обчислення загальної суми товарів у кошику
	totalsum := calculateCartTotal(client.ClientID)
	// обчислення знижки
	discount := calculateDiscount(totalsum)
	// обчислення фінальної суми
	finalsum := totalsum - discount

	// вивід повідомлень
	fmt.Println("Знижка: %.0f%%", globalDisc)
	fmt.Printf("Загальна сума: %.2f грн\n", finalsum)
}

// Підрахунок загальної суми в кошику
func calculateCartTotal(ClientID int) float64 {
	// змінна для зберегання загальної суми в кошику
	totalsum := 0.0
	for _, cart := range carts {
		// перевірка айди клієнта на існування
		if cart.ClientID == ClientID {
			for _, prod := range products {
				// перевірка на наявність товару в кошику
				if prod.ID == cart.ProductID {
					// розрахунок загальної суми
					totalsum += prod.Price * float64(cart.Quantity)
					break
				}
			}
		}
	}
	return totalsum
}

// Розрахунок загальної суми з урахуванням знижок
func calculateTotalSum(cli int) {
	// змінна для зберегання загальної сумми
	totalSum := calculateCartTotal(cli)

	// вивід повідомлень
	fmt.Printf("Вартість доставки: %.2f грн\n", del)
	fmt.Printf("Загальна сума до сплати: %.2f грн\n", totalSum+del)

	// получення відповіді від користувача
	answer := getStringInput("Підтвердити замовлення? (y/n): ")

	// згідно відповіді формується заказ чи ні
	if answer == "y" || answer == "Y" {
		var o Order
		o.ClientID = cli
		o.Sum = totalSum + del
		o.createOrders()
	} else {
		fmt.Println("Замовленя скасовано.❌")
	}
}

// Застосування знижки
func setDiscount() {
	newDisc := getNumInput("Введіть новий відсток знижки: ")

	// перевірка на коректність вводу
	if newDisc < 0 || newDisc > 100 {
		fmt.Println("Некоректне введене значення.Введіть від 0 до 100.❌")
		return
	}

	// змінює знижку
	globalDisc = newDisc
	// вивід повідомлення
	fmt.Printf("Знижку встановлено: %.0f%% ✅\n", globalDisc)
}

// Система замовлень
// Створення замовлення з кошика
func (o *Order) createOrders() bool {
	// змінна для зберегання всіх товарів клієнта
	var clientCartItems []Cart
	for _, cart := range carts {
		// перевірка айди клієнта
		if cart.ClientID == o.ClientID {
			// додавання товарів
			clientCartItems = append(clientCartItems, cart)
		}
	}

	// перевірка на порожність кошика
	if len(clientCartItems) == 0 {
		fmt.Println("Кошик порожній.Неможливо створити замовлення.❌")
		return false
	}

	// розрахуннок загальної суми
	totalSum := calculateCartTotal(o.ClientID)

	// заповненя замовлення
	o.OrdersID = len(orders) + 1 // ID замовлення
	o.Items = clientCartItems    // Список Товарів
	o.Sum = totalSum             // Підумкова сума
	o.Status = "pending"         // Статус
	o.CreateData = time.Now()    // поточна дата і час

	// додавання замовлення в слайс
	orders = append(orders, *o)

	// очищення кошику клієнта після оформлення заказу
	for id, cart := range carts {
		if cart.ClientID == o.ClientID {
			delete(carts, id)
		}
	}

	// вивід повідомлення
	fmt.Printf("Замовлення #%d успішно створено! ✅\n", o.OrdersID)
	fmt.Printf("Статус: %s\n", o.Status)
	return true
}

// Перегляд історії замовлень клієнта
func displayHistoryOrders(id int) {
	found := false
	for _, order := range orders {
		// перевірка чи належить замвлення данному клієнту
		if order.ClientID == id {
			found = true
			// виводимо інформацію про замовлення
			fmt.Printf("Замовленя #%d | Дата: %s | Статус: %s\n", order.OrdersID, order.CreateData.Format("02.01.2006 15:04"), order.Status)
			// виводимо вміст кошика
			fmt.Println("Вміст Кошика.")
			for i, item := range order.Items {
				// шукаємо товар за ID
				for _, prod := range products {
					if prod.ID == item.ProductID {
						// виводимо інформації про товар у замовленні
						fmt.Printf("%d. %s x%d - %.2f грн\n", i+1, prod.Name, item.Quantity, prod.Price*float64(item.Quantity))
						break
					}
				}
			}
			// вивід загальної суми
			fmt.Printf("Загальна сума: %.2f\n", order.Sum)
		}
	}

	if !found {
		fmt.Println("Історія замовлень клієнта порожня.❌")
	}

}

// Зміна статусу замовлення
func (o *Order) changeStatusOrder(newStatus string) bool {
	for i, order := range orders {
		// перевірка чи знайдено потрібне замовлення
		if order.OrdersID == o.OrdersID {
			// оновлення статусу замовлення
			orders[i].Status = newStatus
			o.Status = newStatus
			return true
		}
	}
	return false
}

// Розрахунок загальної вартості з доставкою
func calculeteOrdersSum(id int) {
	for _, order := range orders {
		//  знаходження замовлденя по айди
		if order.OrdersID == id {
			// підрахунок загальної суми
			total := order.Sum + del
			// вивід інформації
			fmt.Printf("Замовленя #%d\n", order.OrdersID)
			fmt.Printf("Сума товарів: %.2f грн\n", order.Sum)
			fmt.Printf("Сума доставки: %.2f грн\n", del)
			fmt.Printf("Загальна сума: %.2f грн\n", total)
			return
		}
	}
	// вивід повідомлення якщо замовлення не знайдено
	fmt.Println("Замовленя не знайдено.❌")
}

// Отримує текстове введення
func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	input, _ = reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Отримує ціло числене введення
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

// Отримує числове введення
func getNumInput(prompt string) float64 {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	//перевірка на пробіли
	if strings.ContainsAny(input, " \t") {
		fmt.Println("Некоректне введення. Ведіть ціле число без пробілів. ❌")
		return -1
	}

	//перевірка на введеня числа
	var value float64
	_, err := fmt.Sscanf(input, "%f", &value)
	if err != nil {
		fmt.Println("Некоректне введення. Ведіть число. ❌")
		return -1
	}
	return value
}

// Функції меню

// головне меню
func mainMenu() {
	fmt.Println("\n📋 Головне меню:")
	fmt.Println("1. 📦 Управління товарами")
	fmt.Println("2. 👤 Управління клієнтами")
	fmt.Println("3. 🛒 Кошик покупок")
	fmt.Println("4. 📦 Замовлення")
	fmt.Println("5. 📊 Статистика магазину")
	fmt.Println("6. 🚪 Вихід")
}

// меню управлінняя товарами
func productMenu() {
	fmt.Println("1. 📦 Додати товар")
	fmt.Println("2. 📋 Переглянути всі товари")
	fmt.Println("3. ✅ Переглянути всі товари в наявності")
	fmt.Println("4. 🆔 Знайти товар за ID")
	fmt.Println("5. 🔎 Пошук за назвою")
	fmt.Println("6. ✏️ Оновити товар")
	fmt.Println("7. 🗑️ Видалити товар з каталогу")
	fmt.Println("8. 🔙 Повернутися до головного меню")
}

// меню управління клієнтами
func clientMenu() {
	fmt.Println("1. 📝 Реєстрація клієнта")
	fmt.Println("2. 👁️ Переглянути інформацію про клієнта")
	fmt.Println("3. 🔄 Оновлення контактних даних клієнта")
	fmt.Println("4. 🔙 Повернутися до головного меню")
}

// менб управління кошиком
func cartMenu() {
	fmt.Println("1. ➕ Додати товар до кошика")
	fmt.Println("2. ➖ Видалити товар з кошика")
	fmt.Println("3. 👀 Переглянути кошик")
	fmt.Println("4. 💸 Застосувати знижку")
	fmt.Println("5. ✅ Оформити замовлення")
	fmt.Println("6. 🔙 Повернутися до головного меню")
}

// меню управління замовленями
func orderMenu() {
	fmt.Println("1. 📜 Перегляд історії замовлень")
	fmt.Println("2. 🔁 Зніма статусу замовлення")
	fmt.Println("3. 💰 Розрахунок загальної вартосі")
	fmt.Println("4. 🔙 Повернутися до головного меню")
}

// вібір із меню товарів
func productChoise() {
	fmt.Println("\n--- Меню товарів ---")
	productMenu()
	for {
		choise := getIntInput("> ")
		switch choise {
		case 1:
			choiseAddProduct()
		case 2:
			fmt.Println("\n--- Всі товари ---")
			displayAllProducts()
		case 3:
			fmt.Println("\n--- Всі товари в наявності ---")
			displayAllProductsStock()
		case 4:
			choiseSearchProdId()
		case 5:
			choiseSearchProdName()
		case 6:
			choiseUpdateProduct()
		case 7:
			choiseDeleteProduc()
		case 8:
			return
		default:
			fmt.Println("\nТакого вибору немає!❌")
		}
	}
}

// додавання товару
func choiseAddProduct() {
	fmt.Println("\n--- Додавання товару ---")

	// запрос у пользователя, імені опису ціни категорії кількість товару
	name := getStringInput("Введіть назву товару: ")
	desc := getStringInput("Введіть опис: ")
	price := getNumInput("Введіть ціну: ")
	categ := getStringInput("Введіть категорію: ")
	stock := getIntInput("Введіть кількість на складі:")

	// создання товару
	prod := Product{
		Name:        name,
		Description: desc,
		Price:       price,
		Category:    categ,
		Stock:       stock,
	}

	// вивід інформації про создання товара
	if prod.addProducts() {
		fmt.Printf("\nТовар %v успішно додано до каталогу! ✅\n", name)
	} else {
		fmt.Println("Помилка в додаванні товара.Такий товар вже існує!❌\n")
	}
}

// пошук товара за його айди
func choiseSearchProdId() {
	fmt.Println("\n--- Пошук по ID ---")

	// запрос айди у пользователя
	id := getIntInput("Введіть ID товару: ")
	searchProductById(id)
}

// пошук товара за іменем
func choiseSearchProdName() {
	fmt.Println("\n--- Пошук по імені ---")

	// запрос назви товару у пользователя
	name := getStringInput("Введіть назву товару: ")
	searchProductByName(name)
}

// оновлення товару
func choiseUpdateProduct() {
	fmt.Println("\n--- Оновлення товару ---")

	// запрос у пользователя ID новой ціни і нову кількість
	id := getIntInput("Введіть ID товару: ")
	newPrice := getNumInput("Введіть нову ціну: ")
	newStock := getIntInput("Введіть нову кількість на складу: ")

	// пошук товара за айди и оновлення товару якщо такий товар є
	for i, _ := range products {
		if products[i].ID == id {
			products[i].UpdatePriceStock(newPrice, newStock)
			fmt.Println("Товар успішно оновлено. ✅\n")
			return
		}
	}
	fmt.Println("Товара з таким ID не знайдено!❌\n")
}

// видалення товару з каталогу
func choiseDeleteProduc() {
	fmt.Println("\n--- Видалення товара ---")

	// запрос айді у пользователя
	id := getIntInput("Введіть ID товару: ")

	// пошук товара за його айди щоб видалити
	for _, prod := range products {
		if prod.ID == id {
			if prod.deleteProducts() {
				fmt.Println("Товар успішно видалено з каталогу. ✅\n")
			} else {
				fmt.Println("Помилка при видаленні товара.❌\n")
			}
			return
		}
	}

	fmt.Println("Товара з таким ID не знайдено.❌\n")
}

// вібір з меню клієнтів
func clientChoise() {
	fmt.Println("\n--- Меню клієнтів ---")
	clientMenu()
	for {
		choise := getIntInput("> ")
		switch choise {
		case 1:
			choiseAddClient()
		case 2:
			choiseDisplayClientInfo()
		case 3:
			choiseUpdateClient()
		case 4:
			return
		default:
			fmt.Println("\nТакого вибору немає!❌")
		}
	}
}

// додавання клієнта
func choiseAddClient() {
	fmt.Println("\n--- Реєстрація клієнта ---")

	// запрос імені та прізвище
	name := getStringInput("Введіть ім'я: ")
	surname := getStringInput("Введіть прізвище: ")

	// создання клієнта
	client := Customer{
		Name:    name,
		Surname: surname,
	}

	// додавання клієнта та вівід повідомлення про додавання
	if client.registerClient() {
		fmt.Println("Клієнт успішно додано. ✅")
	} else {
		fmt.Println("Помилка в додаванні клаєнта.Такий клієнт вже існує!❌")
	}
}

// перегляд інформації про клієнта
func choiseDisplayClientInfo() {
	fmt.Println("\n--- Перегляд інформації ---")

	// запрос імені клієнта
	name := getStringInput("Введіть ім'я клієнта: ")

	// вивід інформації про клієнта
	checkClientInfo(name)
}

// оновлення контактних данних клієнта
func choiseUpdateClient() {
	fmt.Println("\n--- Оновлення данних ---")

	// запросс айди клієнта
	id := getIntInput("Введіть ID клієнта: ")

	//змінна для храніння клієнта
	var client *Customer

	// пошук клієнта за введенем айді
	for i, _ := range customers {
		if customers[i].ClientID == id {
			client = &customers[i] //передача ссилки на клієнта
			break
		}
	}

	// вивід повідомлення якщо клієнта не знайдено
	if client == nil {
		fmt.Println("Клієнта не знайдено.❌")
		return
	}

	// запрос у пользователя зміни щоб змінити ім'я чи прізвищя
	change := getIntInput("Оберіть що бажаєте змінити (1 - Ім'я, 2 - Прізвище): ")
	// вивід повідомленя якщо вибор був не правельний

	if change != 1 && change != 2 {
		fmt.Println("Неправильний вибір будь ласка, вибуріть 1 або 2.❌")
		return
	}

	// змінна ім'я чи прівзвищя за обраним зміной клієнта
	switch change {
	case 1:
		// запрос імені
		name := getStringInput("Введіть нове ім'я: ")
		client.updateClient(change, name)
	case 2:
		// запрос прізвищя
		surName := getStringInput("Введіть нове прізвище: ")
		client.updateClient(change, surName)
	}
}

// вибір з меню кошика
func cartChoise() {
	fmt.Println("\n--- Меню кошика ---")
	cartMenu()
	for {
		choise := getIntInput("> ")
		switch choise {
		case 1:
			choiseAddToCarts()
		case 2:
			choiseDeleteCartProd()
		case 3:
			choiseCheckCart()
		case 4:
			fmt.Println("\n--- Застосування знижки ---")
			setDiscount()
		case 5:
			choiseMakeOrder()
		case 6:
			return
		default:
			fmt.Println("\nТакого вибору немає!❌")
		}
	}
}

// додавання в кошик
func choiseAddToCarts() {
	fmt.Println("\n--- Додавання до кошика ---")

	// запрос ID клієнта ID товару, кількості товару
	cliID := getIntInput("Введіть ID клієнта: ")
	prodID := getIntInput("Введіть ID товару: ")
	qty := getIntInput("Введіть кількість товару: ")

	//  создання кошика
	cartItem := Cart{
		ClientID:  cliID,
		ProductID: prodID,
		Quantity:  qty,
	}

	//  додавання в кошик та вивід повідомеленя
	if cartItem.addCarts() {
		fmt.Println("Товар успішно додано до кошика. ✅")
	} else {
		fmt.Println("Товар з таким ID не знайдено!❌")
	}
}

// видалення товару з кошика
func choiseDeleteCartProd() {
	fmt.Println("\n--- Видалення з кошика ---")

	//  запрос айді клієнта та айді товару
	cliID := getIntInput("Введіть ID клієнта: ")
	prodID := getIntInput("Введіть ID товару: ")

	//  создання кошика
	cartDel := Cart{
		ClientID:  cliID,
		ProductID: prodID,
	}

	// видалення товару з кошика та вивід повідомлення
	if cartDel.deleteProductFromCart() {
		fmt.Println("Товар успішно видален. ✅")
	} else {
		fmt.Println("Товар не знайдено у кошику клієнта!❌")
	}
}

// перевірка кошика
func choiseCheckCart() {
	fmt.Println("\n--- Ваш кошик ---")

	// запрос імені у пользователя
	name := getStringInput("Введіть ім'я клієнта: ")

	// вивід інформації про коошик клієнта
	CheckCartItem(name)
}

// оформлення замовлення
func choiseMakeOrder() {
	fmt.Println("\n--- Оформлення замовлення ---")

	// запрос айді у пользователя
	id := getIntInput("Введіть ID клієнта: ")
	// оформлення замовлення
	calculateTotalSum(id)
}

// вибір з меню замовлень
func orderChoise() {
	fmt.Println("\n--- Меню замовлень ---")
	orderMenu()
	for {
		choise := getIntInput("> ")
		switch choise {
		case 1:
			choiseCheckHistory()
		case 2:
			choiseChangeStatus()
		case 3:
			choiseCalculateOrderSum()
		case 4:
			return
		default:
			fmt.Println("\nТакого вибору немає!❌")
		}
	}
}

// просмотр історія замовлень клієнта
func choiseCheckHistory() {
	fmt.Println("\n--- Історія Замовлень ---")

	// запрос айді клієнта у пользователя
	id := getIntInput("Введіть ID клієнта: ")

	// вівід історії замовлень пользователя
	displayHistoryOrders(id)
}

// зміна статусу замовлень
func choiseChangeStatus() {
	fmt.Println("\n--- Зміна статусу ---")

	// запрос номер замовлення
	id := getIntInput("Введіть номер замовлення: ")

	// змінна для зберегання кошику клієнта
	var order *Order
	// пошук номера замовлення
	for i, _ := range orders {
		if orders[i].OrdersID == id {
			order = &orders[i] // передача ссілки на замовлення
			break
		}
	}

	// вивід повідомлення якщо замовлення не знайдено
	if order == nil {
		fmt.Println("Замовлення з таким ID не знайдено.❌")
		return
	}

	// запрос нового статусу замовлення
	newStatus := getStringInput("Введіть новий статус: ")

	// зміна статусу замовлення та вівід повідомлення
	if order.changeStatusOrder(newStatus) {
		fmt.Println("Статус замовлення успішно оновлено! ✅")
	} else {
		fmt.Println("Не вдалося змінити статус замовлення.❌")
	}
}

// розрахунок вартості
func choiseCalculateOrderSum() {
	fmt.Println("\n--- Розрахунок вартості ---")

	// запрос айді клієнта
	id := getIntInput("Введіть ID клієнта: ")

	// вивід розрахунку вартосі
	calculeteOrdersSum(id)
}

// вивід статистики магазину
func showShopStats() {
	fmt.Print("\n--- Статистика магазину ---\n")

	// підрахунок активних товарів в кошику
	totalCartItem := 0
	for _, cart := range carts {
		totalCartItem += cart.Quantity
	}

	// підрахунок загального прибутку магазину
	var totalProf float64
	for _, order := range orders {
		totalProf += order.Sum
	}

	// пошук популярного товару
	popularItem := make(map[int]int)
	for _, order := range orders {
		for _, item := range order.Items {
			popularItem[item.ProductID] += item.Quantity
		}
	}
	// змінни для храніння:
	maxSold := 0              // продаж популярного товару
	var topProductName string // інем популярного товару

	// пошук самого популярного товара та його колічество продажів
	for _, prod := range products {
		if popularItem[prod.ID] > maxSold {
			maxSold = popularItem[prod.ID]
			topProductName = prod.Name
		}
	}
	// вивід повідомлень
	fmt.Printf("📦 Кількість товарів у каталозі: %d \n", len(products))
	fmt.Printf("👥 Зареєстровано клієнтів: %d \n", len(customers))
	fmt.Printf("🛒 Активні товари у кошиках: %d \n", totalCartItem)
	fmt.Printf("📃 Усього Замовлень: %d \n", len(orders))
	fmt.Printf("💰 Загальний Прибуток: %.2f грн \n", totalProf)
	if maxSold > 0 {
		fmt.Printf("📦 Найпопулярніший товар: %s (%d продажів) \n", topProductName, maxSold)
	}
}

// основна функція
func main() {
	fmt.Printf("=== Онлайн-магазин \"%s\" ===\n", StoreName)
	// вібір з головного меню
	for {
		mainMenu()
		choise := getIntInput("> ")
		switch choise {
		case 1:
			productChoise()
		case 2:
			clientChoise()
		case 3:
			cartChoise()
		case 4:
			orderChoise()
		case 5:
			showShopStats()
		case 6:
			fmt.Printf("👋 Допобачення!\nНадіюсь вам понравилось в нашому %s!\n", StoreName)
			return
		}
	}
}
