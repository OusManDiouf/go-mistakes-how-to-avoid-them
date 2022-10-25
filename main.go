package main

import "fmt"

type ElementsSlice []string

func (e *ElementsSlice) add(str string) {
	*e = append(*e, str)
}

type ElementsMap map[int]string

func (m ElementsMap) add(k int, v string) {
	m[k] = v
}

func main() {
	elementsSlice := ElementsSlice{"Vanuatu", "Montenegro"}
	elementsSlice.add("Ethiopia")
	//fmt.Println(elements)

	elementsMap := ElementsMap{}
	elementsMap.add(1, "Germany")
	fmt.Println(elementsMap)
}
