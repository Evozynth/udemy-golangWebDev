package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type menuItem struct {
	Name string
	Price string
	Description string
}

type menuSection struct {
	Name string
	Items []menuItem
}

type menu []menuSection

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

func main() {

	m := menu{
		menuSection{
			Name: "Breakfast",
			Items: []menuItem{
				{
					Name: "Hungry Dan",
					Price: "9.49",
					Description: "2 Fresh eggs with a generous portion of crispy fried red potatoes or hash browns and your choice of breakfast meat, served with 2 slices of toast or 2 pancakes",
				},
				{
					Name: "Lumberyard Sandwich",
					Price: "7.99",
					Description: "2 Fresh eggs scrambled and stacked with your choice of breakfast meat, Wisconsin cheddar and American cheese on grilled sourdough, served with choice of breakfast potato",
				},
			},
		},
		menuSection{
			Name: "Lunch",
			Items: []menuItem{
				{
					Name: "Nachos",
					Price: "9.99",
					Description: "Your choice of steak or chicken over refried beans, fresh pico de gallo, crisp tortilla chips, topped with Wisconsin cheddar cheese. Served with sour cream and jalapenos",
				},
				{
					Name: "Beef Tender Bites",
					Price: "7.99",
					Description: "A lumberyard exclusive. Hand cut and battered beef tender tips. Deep fried until perfectly done. Served with tiger sauce",
				},
			},
		},
		menuSection{
			Name: "Dinner",
			Items: []menuItem{
				{
					Name: "Quesadilla",
					Price: "8.49",
					Description: "Large tortilla filled with pico de gallo and melted cheddar and mozzarella cheese. Choice of steak or chicken",
				},
				{
					Name: "Chips and Queso",
					Price: "6.99",
					Description: "Crisp tortilla chips served with our jalapeno cheese sauce topped with pico de gallo",
				},
			},
		},
	}

	r := restaurants{
		restaurant{
			Name: "Lumberyard Bar & Grill",
			Menu: m,
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
