package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

/*
*Ponteiro vc usa para referenciar uma variável que poderá sofrer alteração.
* A msm após sofrer alteração será referenciada, com a alteração nos demais lugares que ela é utilizada
* Quando a função é executada ela é limpada da memoria.
* Logo a alteração não ficará salva para as demais ações.
* e o valor vai ser pego do objeto que foi inserido
* Obs.: em alguns casos o valor não deverá ser alterados, e em outro precisa ser.
Resumo o ponteiro irá alterar o valor original
ex:
func (o Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
	println(o.FinalPrice)
}

func (o Order) CalculateFinalPrice2(price) {
	o.price = price -> esta sendo duplicado na memoria
	o.FinalPrice = o.Price + o.Tax
	println(o.FinalPrice)
	a:= o.price
b:= &a -> endereço na memoria aonde 'a' é salvo
Como ambos apontam para o msm endereço se o 'b' for alterado 'a' tbm sera
}
*/

func NewOrder(id string, price float64, tax float64) (*Order,error) {
	// Return uma nova struct ou error
	// sempre retornara um dos dois

	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()
	if err != nil {
		return nil,err
	}
	return order,nil
}

func (o *Order) Validate() error {
	// sem o ponteiro o final price só vai ter o valor alterado dentro da function
	if o.ID == "" {
		return errors.New("id is required")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.Tax < 0 {
		return errors.New("tax must be than or equal than zero")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	// sem o ponteiro o final price só vai ter o valor alterado dentro da function
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()
	if err != nil{
		return err
	}
	return nil
}

// Pode executar ele assim
// order,err := entity.NewOrder("1",10,0)
// if err != nill { println(err.Error())} else {restante da logica}