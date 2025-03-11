package main

import (
	"SimplyConList/simlist"
	"fmt"
)

func main() {
	fmt.Println("Реализация односвязного списка.")
	var l simlist.SimList
	l.AppendToHead("Коля")
	l.AppendToHead("Федя")
	l.AppendToHead("Семен")
	l.AppendToHead("Александр")
	l.Print()
	fmt.Println()

	fmt.Printf("List as slice: %v\n", l.GetAsSlice())
	fmt.Println()

	fmt.Printf("поиск элемента <Федя>: \n")
	if v, idx, ok := l.FindByValue("Федя"); !ok {
		fmt.Println("Не удалось найти элемент!")
	} else {
		fmt.Printf("Найден элемент <%s> с индексом %d\n", v.Value, idx)
	}
	fmt.Println()

	fmt.Println("Вставка <Дормидонт>:")
	if ok := l.Insert(1, "Дормидонт"); !ok {
		fmt.Println("Не удалось вставить элемент")
	} else {
		l.Print()
	}
	fmt.Println()

	fmt.Println("Удаление элемента <Федя>:")
	l.DeleteByValue("Федя")
	l.Print()
}
