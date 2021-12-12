package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int64
	next  *Node
}

func (head *Node) Print() {
	for head != nil {
		fmt.Println((*head).value)

		head = head.next
	}
}

func (head *Node) Reverse() *Node {
	if head == nil {
		return head
	}

	var prevNode, nextNode *Node
	currentNode := head

	for currentNode.next != nil {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	currentNode.next = prevNode
	head = currentNode

	return head
}

func main() {
	var head, last *Node

	scanner := bufio.NewScanner(os.Stdin)

	println("Enter numbers: ")

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			panic(err)
		}

		// создадим временную переменную с указателем на ноду с введенным числом
		tmp := &Node{num, nil}

		// если указатель пустой, значит ввели первое число
		if head == nil {
			// запомнили первую ноду - начало связанного списка,
			// и дублируем в переменную для последней на текущий момент ноды
			// т.е. запишем в head и last один и тот же указатель на место/ноду с введенным первым числом
			// nil , nil
			head, last = tmp, tmp
			// &{num, nil} &{num, nil}
		} else {
			// &{num, nil} - сейчас last - это указатель на место, где лежит нода с предыдущей итерации, с предыдущим числом и без указателя
			// добавляем в эту ноду указатель на место с созданной в этой итерацации временной нодой
			last.next = tmp // &{num, pointer}
			// перезапишем last указателем на временную ноду созданую в текущей итерации
			last = tmp // &{newNum, nil}
		}
	}

	println("Reversed list: ")
	head.Reverse().Print()
}
