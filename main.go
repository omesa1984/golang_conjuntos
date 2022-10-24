package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

type humanos interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, p.sobrenome, "diz oi!")
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{"Juan", "Andarin"}

	dizerAlgumaCoisa(&p1)
}

/*
- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/
