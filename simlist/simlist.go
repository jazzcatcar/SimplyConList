package simlist

import "fmt"

type SimNode struct {
	Value string
	prev  *SimNode
}

type SimList struct {
	SimNode
	head *SimNode
}

// Len возвращает длину списка
func (sl *SimList) Len() int {
	var cnt int
	for curr := sl.head; curr != nil; curr = curr.prev {
		cnt += 1
	}
	return cnt
}

// AppendToHead добавление элемента в список
func (sl *SimList) AppendToHead(pValue string) {
	v := new(SimNode)
	v.Value = pValue
	v.prev = sl.head
	sl.head = v
}

// AppendToEnd добавляет элемент в конец списка
func (sl *SimList) AppendToEnd(pValue string) {
	curr := sl.head
	for ; curr != nil; curr = curr.prev {
		if curr.prev == nil {
			tmp := new(SimNode)
			tmp.Value = pValue
			curr.prev = tmp
		}
	}
}

// Insert добавление элемента в список после элемента с индексом pIdx
func (sl *SimList) Insert(pIdx int, pValue string) bool {
	curr := sl.head
	for idx := 1; curr != nil; idx++ {
		if idx == pIdx {
			tmp := new(SimNode)
			tmp.Value = pValue
			tmp.prev = curr.prev
			curr.prev = tmp
			return true
		}
		curr = curr.prev
	}
	return false
}

// GetAsSlice возвращает элементиы списка в виде среза
func (sl *SimList) GetAsSlice() []string {
	var res []string
	for v := sl.head; v != nil; v = v.prev {
		res = append(res, v.Value)
	}
	return res
}

// Clear удаляет всех элементов списка
func (sl *SimList) Clear() {
	for {
		if sl.head == nil {
			return
		}
		vTemp := sl.head.prev
		sl.head.prev = nil
		sl.head = vTemp
	}
}

// FindByValue возвращает узел с найденным значением Value
func (sl *SimList) FindByValue(pValue string) (*SimNode, int, bool) {
	//поиск элемента в списке
	curr := sl.head
	for idx := 1; curr != nil; idx++ {
		if curr.Value == pValue {
			return curr, idx, true
		}
		curr = curr.prev
	}
	return nil, -1, false
}

// DeleteByValue удаление элемента со значением Value
func (sl *SimList) DeleteByValue(pValue string) bool {
	var i int
	var previous *SimNode
	for curr := sl.head; curr != nil; {
		i++
		if curr.Value == pValue {
			if curr == sl.head {
				sl.head = sl.head.prev
				curr.prev = nil
			} else {
				previous.prev = curr.prev
				curr.prev = nil
			}
			return true
		}
		if i != 0 {
			previous = curr
		}
		curr = curr.prev
	}
	return false
}

// Print печать списка
func (sl *SimList) Print() {
	if sl.head == nil {
		fmt.Println("Simple list is empty")
	} else {
		for v := sl.head; v != nil; v = v.prev {
			fmt.Printf("%s\n", v.Value)
		}
	}
}
