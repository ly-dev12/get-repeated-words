package main

import (
	"fmt"
	"strings"
)

func main() {

	var texto string = "Hola mundo que tal hola aqui un repetetidor de palabras normal sin normalize, si repetimos palabras : tal hola aqui repetidor"
	var strs = make(map[string]int, 1)

	slicing := strings.Split(texto, " ")
	//slicing = append(slicing, "dfsfd") 11 11, se duplica la capacidad al ser un slice dinamico y no entrar mas data

	fmt.Println(slicing, len(slicing), cap(slicing))

	for _, trozoTexto := range slicing {
		strs[trozoTexto]++
	}

	// el 1 fmt.Println(strs, len(strs))

	for key, value := range strs {
		fmt.Printf("%s = %d\n", key, value)
	}

	fmt.Println(strs)

	change(strs) // LOS MAPAS YA SON REFERENCIADOS, NO HACE FALTA UNA MEMORIA DE PUNTERO

	fmt.Println(strs)

}

func change(mapa map[string]int) {
	fmt.Println(mapa)
	mapa["Hola"] = 2
	mapa["aqui"] = 45
}
