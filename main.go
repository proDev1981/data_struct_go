package main

import . "object/core" // name proyect => object
import "fmt"

func main() {

	person := Tupla(
		Struct(
			Int("#age", 20),
			String("#name", "alberto"),
			Bool("#casado", false),
			Float("#altura", 1.81),
		),
		Struct(
			Int("#age", 21),
			String("#name", "paco"),
			Bool("#casado", true),
			Float("#altura", 1.76),
		),
	)

	fmt.Println(person.Struct("uno").Float("altura"))

	alberto := person.Filter(func(item Value) bool {
		return item.Int("age") >= 20
	})
	fmt.Println(alberto)

	person.ForEach(func(index int, item Value) {
		fmt.Println(item.String("name"))
		fmt.Println(item.Int("age"))
		fmt.Println(item.Bool("casado"))
		fmt.Println(item.Float("altura"))
	})

}
