package main

import "fmt"

type boardColumn struct {
	previous *boardColumn
	next     *boardColumn
	column   []int
}

func (bc *boardColumn) print() error {
	fmt.Printf("%+v\n", bc.column)
	return nil
}

func fillCollumn(column []int, val, rangeFrom, rangeTo int) {
	for index := rangeFrom; index < rangeTo; index++ {
		column[index] = val
	}
}
