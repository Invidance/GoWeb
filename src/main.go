package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
	Price float64
	Name  string
	Count int
}

func initPages() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/products/", futuresPage)
	http.HandleFunc("/auth/", authPage)
}

func main() {
	initPages()
	fmt.Println("Pages inited")
	http.ListenAndServe("", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Open main page")
	page, err := template.ParseFiles("templates/header.html", "templates/index.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	page.ExecuteTemplate(w, "index", nil)
}

func futuresPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Open product page")

	page, err := template.ParseFiles("templates/header.html", "templates/futures.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	list := []Product{}
	list = append(list, Product{30.0, "Milk", 10})
	list = append(list, Product{50.0, "Tamato", 1000})
	list = append(list, Product{120.0, "Cucumber", 5000})
	list = append(list, Product{53.0, "Onion", 2000})
	list = append(list, Product{55.0, "Carrot", 600})
	list = append(list, Product{25.0, "Bread", 700})
	list = append(list, Product{150.0, "Butter", 300})
	list = append(list, Product{200.0, "Cheese", 250})
	list = append(list, Product{45.0, "Mushroom", 1200})
	list = append(list, Product{75.0, "Oil", 750})

	list = sort(list)

	page.ExecuteTemplate(w, "futures", list)
}

func authPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Open auth page")
	page, err := template.ParseFiles("templates/header.html", "templates/auth.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	page.ExecuteTemplate(w, "auth", nil)
}

func sort(arr []Product) []Product {
	var count int = 1
	for count > 0 {
		count = 0
		for i := 1; i < len(arr); i++ {
			arr[i-1], arr[i] = swap(arr[i-1], arr[i], &count)
		}
	}

	return arr
}

func swap(p1 Product, p2 Product, count *int) (Product, Product) {
	if p1.Price > p2.Price {
		(*count) += 1
		return p2, p1
	} else {
		return p1, p2
	}
}
