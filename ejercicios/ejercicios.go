package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
	"strings"
)

func InvertirCadena(cadena string) string { // orden O(2n)
	var pila stack.Stack
	var nueva_cadena string
	for _, value := range cadena { // O(n)
		pila.Push(string(value))
	}

	for i := 0; i < len(cadena); i++ { // O(n)
		caracter, _ := pila.Pop()
		nueva_cadena += caracter.(string)
	}
	return nueva_cadena
}

func Palindromo(cadena string) bool { // O(2n)
	polindromo := InvertirCadena(cadena) // O(2n)
	return polindromo == cadena
}

func Balanceada(simbolos string) bool { // O(n)
	var pila stack.Stack
	balanceada := true
	i := 0
	for i < len(simbolos) && balanceada { // O(n)
		simbolo := string(simbolos[i])
		if strings.Contains("([{", simbolo) {
			pila.Push(simbolo)
		} else {
			if pila.IsEmpty() {
				balanceada = false
			} else {
				simbolo_abierto, _ := pila.Pop()
				if !(CompararOpuesto(simbolo_abierto.(string), simbolo)) {
					balanceada = false
				}
			}
		}
		i++
	}
	return balanceada
}

func CompararOpuesto(abierto, cerrado string) bool {
	abiertos := "([{"
	cerrados := ")]}"
	return strings.Index(abiertos, abierto) == strings.Index(cerrados, cerrado)

}

func UnirColas(q1, q2 queue.Queue) queue.Queue { // O(n)
	for !q2.IsEmpty() { // O(n)
		value, _ := q2.Dequeue()
		q1.Enqueue(value)
	}
	return q1
}
