package common

import ()

type Array struct {
	Array  [][]int32
	Width  int
	Height int
}

func NewArray(w, h int) *Array {
	a := new(Array)
	a.Width = w
	a.Height = h
	a.Array = make([][]int32, w)
	for i := 0; i < w; i++ {
		a.Array[i] = make([]int32, h)
	}
	return a
}

func CloneArray(a *Array) *Array {
	newArray := NewArray(a.Width, a.Height)

	//more stuff here

	return newArray
}
