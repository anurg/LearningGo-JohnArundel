package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        string
}

type Catalog map[int]Book

func Buy(b Book) Book {
	b.Copies--
	return b
}

func GetAllBooks(catalog []Book) []Book {
	// This is wrong thinking
	// result := []Book{}
	// for _, b := range catalog {
	// 	result = append(result, b)
	// }
	// return result
	return catalog
}
func GetBook(Catalog []Book, ID int) Book {
	for _, b := range Catalog {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}

func GetBookMap(Catalog map[int]Book, ID int) (Book, error) {
	b, ok := Catalog[ID]
	if ok {
		return b, nil
	}
	return Book{}, errors.New("book id doesnot exist") //nil-- for testing error condition
}

func GetAllBooksMap(Catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range Catalog {
		result = append(result, b)
	}
	return result
}

func (book Book) GetNetPriceCents() int {
	discount := (book.PriceCents * book.DiscountPercent) / 100
	return book.PriceCents - discount
}

func (c Catalog) GetAllBooksObject() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBookObject(ID int) (Book, error) {
	b, ok := c[ID]
	if ok {
		return b, nil
	}
	return Book{}, errors.New("book id doesnot exist") //nil-- for testing error condition
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("price cannot be set to %d", price)
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("category not allowed, %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
