package main

import "errors"

func removeElement(s []int, index int) ([]int, error) {

	newSlice := []int{}

	if index > (len(s)-1) || index < 0 {
		return s, errors.New("Заданного индекса нет в слайсе")
	}

	newSlice = append(s[:index], s[(index+1):]...)

	return newSlice, nil
}
