package golang-area

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) calculadesc() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.50,
	}
	fmt.Println(produto1, produto1.calculadesc())

	produto2 := produto{"caneta", 2.00, 0.10}
	fmt.Println(produto2, produto2.calculadesc())
}
