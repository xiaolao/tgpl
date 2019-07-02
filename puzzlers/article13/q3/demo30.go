package main

import "fmt"

type Cat struct {
	name           string
	scientificName string
	category       string
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func main() {
	cat := New("little pig", "American shorthair", "cat")
	cat.SetName("monster")           // (&cat).SetName("monster")
	fmt.Printf("The cat: %s\n", cat) // monster

	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat) // The cat monster

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}
