package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

// this method should be declared within same package where you defined the struct
func (v Vertex) abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Pointer receiver. Only thing needs to be mentioned is that the T of *T cannot be a pointer
// I like this
func (v *Vertex) scale(f float64){
	v.X *= f
	v.Y *= f
}

func scaleFunc(v *Vertex, f float64){
	v.X *= f
	v.Y *= f
}

func add(x,y int) int{
	return x + y
}

func main(){
	v := Vertex{3.0, 4.0}
	// Feel weird here becasue v is Vertex (not a pointer) but scale function expects a *Vertex to call it
	// But kinda understand why they do this, go already interprete T or *T for us is using a pointer receiver
	v.scale(10)
	scaleFunc(&v , 10)
	fmt.Println(v.abs())

	p := &Vertex{3.0, 4.0} // [TODO] : this like is super suspicious
	p.scale(5)
	scaleFunc(p, 10)
	// I found interesting part here is `func (v Vertex)` and `func (v *Vertex)` are interchangeable
	// value receiver and pointer receiver are interchangeable
	fmt.Println(p.abs())

}