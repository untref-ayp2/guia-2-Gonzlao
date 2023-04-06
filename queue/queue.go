package queue

import (
	"errors"
	"guia2/stack"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	cola []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	(*q).cola = append((*q).cola, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	head := (*q).cola[0]
	(*q).cola = (*q).cola[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return (*q).cola[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len((*q).cola) == 0
}

type QueueS struct {
	pila1 stack.Stack
	pila2 stack.Stack
}

func (q *QueueS) Enqueue(v any) { // O(n)
	if q.IsEmpty() {
		q.pila1.Push(v)
	} else {
		for !q.pila2.IsEmpty() {
			aux, _ := q.pila2.Pop()
			q.pila1.Push(aux)
		}
		q.pila1.Push(v)
	}
}

func (q *QueueS) Dequeue() (any, error) { // O(n)
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	for !q.pila1.IsEmpty() {
		aux, _ := q.pila1.Pop()
		q.pila2.Push(aux)
	}
	head, _ := q.pila2.Pop()

	return head, nil
}

func (q *QueueS) IsEmpty() bool { // O(1)
	return q.pila1.IsEmpty() && q.pila2.IsEmpty()
}

func (q *QueueS) Front() (any, error) { // O(n)
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	for !q.pila1.IsEmpty() {
		aux, _ := q.pila1.Pop()
		q.pila2.Push(aux)
	}
	head, _ := q.pila2.Top()

	return head, nil
}
