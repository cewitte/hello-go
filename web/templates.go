package main

import (
	"log"
	"os"
	"text/template"
)

// Restaurant contains the restaurant name and menu.
type Restaurant struct {
	Name string
	Menu []Meal
}

// Meal is for the meal name and its items.
type Meal struct {
	Name  string
	Items []string
}

// Restaurants adds support to many restaurants.
type Restaurants []Restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	rs := Restaurants{
		Restaurant{
			Name: "Jack in the Box",
			Menu: []Meal{
				Meal{
					Name: "Breakfast",
					Items: []string{
						"Coffee with Cream",
						"Egg Sandwich",
						"Grapefruit",
						"Orange Juice",
					},
				},
				Meal{
					Name: "Lunch",
					Items: []string{
						"California Jack",
						"French fries",
						"Coke Zero",
						"Apple Pie",
					},
				},
				Meal{
					Name: "Dinner",
					Items: []string{
						"Chicken Salad",
						"Apple Fruit",
						"Fruit Salad",
					},
				},
			},
		},
		Restaurant{
			Name: "McDonald's",
			Menu: []Meal{
				Meal{
					Name: "Breakfast",
					Items: []string{
						"Capuccino",
						"Ham & Cheese Croissant",
						"Brazilian Cheese Bread",
						"Apple Juice",
					},
				},
				Meal{
					Name: "Lunch",
					Items: []string{
						"Peach Juice",
						"Double Bacon Cheeseburger",
						"French Fries",
						"Banana Pie",
					},
				},
				Meal{
					Name: "Dinner",
					Items: []string{
						"Green Tea",
						"Tomatoe Soup",
						"Selected Grapes",
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, rs)

	if err != nil {
		log.Fatal(err)
	}
}
