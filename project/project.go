package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"ProductName"`
	CategoryId  int     `json:"Category"`
	UnitPrice   float64 `json:"Unitprice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"CategoryName"`
}

// ----------Tüm ürünleri döndüren fonksiyon-------------
func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodybytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodybytes, &products)

	return products, err
}

//--------kategorileridöndüren fonksiyon------

func GetAllCategories() []Category {
	response, err := http.Get("http://localhost:3000/categories")
	if err != nil {
		fmt.Println("HATA: ", err)
	}

	defer response.Body.Close()
	bodybytes, _ := ioutil.ReadAll(response.Body)

	var categories []Category
	json.Unmarshal(bodybytes, &categories)

	return categories
}

func Kategorilerigoster() {
	fmt.Printf("%v adet kategori bulunmakta.\n", len(GetAllCategories()))
	kategoriler := GetAllCategories()
	for i := 0; i < len(GetAllCategories()); i++ {
		fmt.Printf("%v : %v \n ", kategoriler[i].CategoryName, kategoriler[i].Id)
	}
}

// belirtilen katagorideki ürünlerin adını döndürür
func Kategoriyegoregetir() {
	Kategorilerigoster()

	var kategori int
	fmt.Print("kategori id si seçin: ")
	fmt.Scan(&kategori)

	urunler, _ := GetAllProducts()
	fmt.Println("[+] Ürünlerimiz:")
	for i := 0; i < len(urunler); i++ {
		if urunler[i].CategoryId == kategori {
			fmt.Println(urunler[i].ProductName, "Price: ", urunler[i].UnitPrice)
		}
	}

}

//----------------------ürün ekleme işlemi--------------------

func AddProduct(urunadi string, kategorino int, fiyat float64) {
	urun := Product{
		ProductName: urunadi,
		CategoryId:  kategorino,
		UnitPrice:   fiyat}
	//id tanımlamak şart değil(tanımlanmazsa otomatik olarak ayarlanır)

	jsonurun, err := json.Marshal(urun)
	if err != nil {
		fmt.Println("HATA: ", err)
	}

	response, err := http.Post(
		"http://localhost:3000/products",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonurun))
	if err != nil {
		fmt.Println("HATA: ", err)
	}

	defer response.Body.Close()
	bodybytes, _ := ioutil.ReadAll(response.Body)

	var postresponse Product
	json.Unmarshal(bodybytes, &postresponse)

	fmt.Println("Kayıt işlemi yapıldı. Durum: ", postresponse) //eğer ekleme işlemi başarılıysa eklenen değerler yazar
	//eğer ekleme işlemi başarısızsa(mesela var olan id tekrar kullanılmak istenrse başarısız olur) 0 değerleri yazar
}
