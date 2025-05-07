package main

import (
	"fmt"
	"sync"
)

func insertionSort(tab []float64) {
	for i := 1; i < len(tab); i++ {
		j := i - 1
		var value float64 = tab[i]
		for ; j >= 0; j-- {
			if tab[j] > value {
				tab[j+1] = tab[j]
			} else {
				break
			}

		}
		tab[j+1] = value

	}
	wg.Done()
}
func fct(line []float64) {
	for _, v := range line {
		fmt.Printf("%f, ", v)
	}
}

func transpose(tab *[][]float64) {
	tmp := *tab
	*tab = [][]float64{}
	for i := 0; i < len(tmp[0]); i++ {
		var line []float64
		for j := 0; j < len(tmp); j++ {
			line = append(line, tmp[i][j])
		}
		*tab = append(*tab, line)
	}
}

func printA(A [][]float64) {
	for _, line := range A {
		fct(line)
		fmt.Println()
	}
}

var wg sync.WaitGroup

func sortRows(tab [][]float64) {
	wg.Add(len(tab))
	for _, line := range tab {
		go insertionSort(line)

	}
	wg.Wait()
}

func main() {
	//t := []float64{5, 3, 4, 1, 99}
	//insertionSort(t)
	//fmt.Print(t)

	//array := [][]float64{{7.1, 2.3, 1.1},
	//	{4.3, 5.6, 6.8},
	//	{2.3, 2.7, 3.5},
	//	{4.5, 8.1, 6.6}}
	//
	//println("Print input ")
	//printA(array)
	//
	////println()
	////
	////println()
	//println("Print sortrows ")
	//sortRows(array)
	//printA(array)
	//
	//println("Print transpose")
	//tanspose(&array)
	//printA(array)
	//
	//println("sort again")
	//sortRows(array)
	//printA(array)
	//
	//println("after all")
	//printA(array)
	//
	//////////////////////////
	//array2 := [][]float64{{1.1, 7.3, 3.2, 0.3, 3.1},
	//	{4.3, 5.6, 1.8, 5.3, 3.1},
	//	{1.3, 2.7, 3.5, 9.3, 1.1},
	//	{7.5, 5.1, 0.6, 2.3, 3.9}}
	//
	////
	//
	//println("Print input ")
	//printA(array2)
	//
	////println()
	////
	////println()
	//println("Print sortrows ")
	//sortRows(array2)
	//printA(array2)
	//
	//println("Print transpose")
	//tanspose(&array2)
	//printA(array2)
	//
	//println("sort again")
	//sortRows(array2)
	//printA(array2)
	//
	//println("after all")
	//printA(array2)

	array := [][]float64{{3, 2, 1},
		{6, 5, 4}}

	println("Print input ")
	printA(array)
	
	println("Print sortrows ")
	sortRows(array)
	printA(array)

	println("Print transpose")
	transpose(&array)
	printA(array)

	println("sort again")
	sortRows(array)
	printA(array)

	println("after all")
	printA(array)
}
