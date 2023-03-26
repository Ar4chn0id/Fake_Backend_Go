package main

// typicode/json-server
// json server başlatma: json-server --watch db.json

import (
	"Proje/project"
)

func main() {

	// fmt.Println("Here we go")
	// urunler := project.GetAllProducts()
	// fmt.Println("tüm ürünler: \n", urunler)
	// project.Kategoriyegoregetir()
	project.AddProduct("Fruitjuice", 2, 10.99)

}

//ctrl k + ctrl c -> comment
//ctrl k + ctrl u -> uncomment
