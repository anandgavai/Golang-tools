package main

import (
	"fmt"
)

type Animal interface {
Speaks() string
}

type Dog struct {
	species string
	name string
	legs int
}

type Fish struct {
	species string
	name string
}

func (a Dog) Speaks() string{
		s:="Woof"
		return s
}


func(f Fish) Speaks() string{
	m:="Blob Blob...."
	return m
}


func behaviour(a Animal){
	fmt.Println(a.Speaks())
}

func main() {
	a:= Dog{"Mammal","Dog",4}
	b:= Fish{"Pices","Dorada"}
  behaviour(a)
	behaviour(b)

}
