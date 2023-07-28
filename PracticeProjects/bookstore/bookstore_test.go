package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Arundale",
		Copies: 10,
	}
	want := 9
	got := bookstore.Buy(b)
	if want != got.Copies {
		t.Errorf("want:%d, got:%d", want, got.Copies)
	}

}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go-Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go-Tools"},
	}
	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go-Tools"},
	}
	want := bookstore.Book{
		ID: 2, Title: "The Power of Go-Tools",
	}
	got := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBookMap(t *testing.T) {
	t.Parallel()
	Catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go-Tools"},
	}
	Catalog[3] = bookstore.Book{ID: 3, Title: "Go Joy"}
	b := Catalog[3]
	b.Title = "Go Joyyy"
	Catalog[3] = b
	want := bookstore.Book{
		ID: 2, Title: "The Power of Go-Tools",
	}
	got, err := bookstore.GetBookMap(Catalog, 2)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(Catalog[1].Title)
	// fmt.Println(Catalog[2].Title)
	// fmt.Println(Catalog[3].Title)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestBookGetBadIDReturns(t *testing.T) {
	t.Parallel()
	Catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go-Tools"},
	}
	_, err := bookstore.GetBookMap(Catalog, 3)
	if err == nil {
		t.Error("Should have got errot but got nil")
	}
}

func TestGetAllBooksMap(t *testing.T) {
	t.Parallel()
	Catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go-Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go-Tools"},
	}
	got := bookstore.GetAllBooksMap(Catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		ID: 1, Title: "For the Love of Go", Author: "John Arundel", PriceCents: 400, DiscountPercent: 50,
	}
	want := 200
	// got := bookstore.GetNetPriceCents(book)
	got := book.GetNetPriceCents()
	if want != got {
		t.Errorf("want:%d, got:%d", want, got)
	}
}

func TestGetAllBooksObject(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go-Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go-Tools"},
	}
	got := catalog.GetAllBooksObject()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBookObject(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go-Tools"},
	}
	want := bookstore.Book{
		ID: 2, Title: "The Power of Go-Tools",
	}
	got, err := catalog.GetBookObject(2)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	want := 3000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want:%d, got:%d", want, got)
	}

}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory("bogus")
	if err == nil {
		t.Fatal("should have got error for setting category, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the love of Go",
	}
	want := "Autobiography"
	err := b.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.Category()
	if want != got {
		t.Errorf("want:%q,got:%q", want, got)
	}
}
