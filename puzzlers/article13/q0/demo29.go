package main

import "fmt"

type AnimalCategory struct {
	kingdom string
	phylum  string
	class   string
	order   string
	family  string
	genus   string
	species string
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class,
		ac.order, ac.family, ac.genus, ac.species)
}

type Animal struct {
	scientificName string
	AnimalCategory
}

func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}

type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func main() {
	category := AnimalCategory{species: "cat"} // cat
	// The animal category: cat
	fmt.Printf("The animal category: %s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	// the animal: American Shorthair (category: cat)
	fmt.Printf("the animal: %s\n", animal)

	cat := Cat{
		name:   "little pig",
		Animal: animal,
	}
	// The cat: American Shorthair: (categary: cat, name: little pig)
	fmt.Printf("The cat: %s\n", cat)
}
