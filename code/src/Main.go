package main

import(
	"fmt"
	"test/src/go-packages/geometry"
	"test/src/go-packages/temp"
	"test/src/go-packages/temp2"
	"test/src/go-packages/temp3"
)



func main(){
	fmt.Println("heyy")
	c := pac1.Cube{Length: 3}
	geometry.CalculateVolume(c,"Cube")

	b:=pac2.Box{Length:2,Width:3,Height:4}
	geometry.CalculateVolume(b,"Box")

	s:=pac3.Sphere{Radius:4}
	geometry.CalculateVolume(s,"Sphere")
	

}