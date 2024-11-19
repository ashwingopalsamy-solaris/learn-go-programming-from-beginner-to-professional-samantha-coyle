package main

import "fmt"

var GlobalLimit = 100

var MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, value string) {
	if len(cache) >= MaxCacheSize {
		return
	}
	cache[key] = value
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func SetBook(isbn, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func SetCD(sku, name string) {
	cacheSet(CacheKeyCD+sku, name)
}

func main() {
	cache = make(map[string]string)
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")
	fmt.Println(GetBook("1234-5678"))
	fmt.Println(GetCD("1234-5678"))
}
