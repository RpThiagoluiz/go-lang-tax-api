package main

import (
	"database/sql"
	"fmt"

	"github.com/RpThiagoluiz/go-lang-tax-api/internal/infra/database"
	"github.com/RpThiagoluiz/go-lang-tax-api/internal/usecase"
)

type Car struct {
	Model string
	Color string
}

func main() {
	/**
	*@atribuição e declaração
	* var car string
	* car = "Ferrari"
	* voce pode declarar direto
	* := declara é atribui a varável ao msm tempo;
	car:="Ferrari"
	car := Car{
		Model: "Ferrari",
		Color: "Red",
	}

	println(car.Color)
	*/

	db, err := sql.Open("sqlite3", "db.sqlite3")


	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)
	
	input := usecase.OrderInput{
		ID: "1234",
		Price: 10.0,
		Tax: 1.0,
	}

	output,err := uc.Execute(input)

	if err!=nil{
		panic(err)
	}

	fmt.Println(output)
}