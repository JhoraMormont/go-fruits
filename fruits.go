package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)
	// - Type of data: 'fruits' extends 'slice' (array with not fixed length);
type fruits [] string

/* - By convention go used 'first letter' of the type of data that this functions is going to receive.
	 - Receiver function, the () in the beginning is telling the function that any var of type 'fruits' have access to this function and
			uses that data. 
	 - Declaration and assignment (:=) of variable. go detects the type of content in this case.
*/
func (f fruits) print() {
	for i, fruit := range f {
		fmt.Println(i, fruit)
	}
}

/*

*/
func newFruits() fruits {
	
	var newFruitSlice fruits

	fruitsValues := []string { "pera", "manzana", "naranja", "ciruela"}
	fruitColors := []string {"roja", "verde", "amarilla"}

		// _ underscore is used so golang compiler dont throw and error because that var is not being used
		for _ , values := range fruitsValues {
			for _ , color := range fruitColors {
				newFruitSlice = append(newFruitSlice, values+" "+color)
			}
		}

	return newFruitSlice
}

/*
	- Double return from function, this is defined in the second ()
	- Slice support the return of elements based on the index passed to the [], discription: [fromIndex, toIndex]
*/
func deal(f fruits, quantity int) (fruits, fruits){
	return f[:quantity], f[quantity:]
}

/*
	- Double return from function, this is defined in the second ()
	- Slice support the return of elements based on the index passed to the [], discription: [fromIndex, toIndex]
*/
func (f fruits) toString() string {
	return strings.Join([]string(f),",")
}

/*
	- []byte(N) notation is a casting type declaration, in this case turns the 'string slice' into a 'byte slice'
*/
func (f fruits) writeFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(f.toString()), 0666)
}

/*
	- Last line is casting file(byte slice) to (string) and then using that in the input to split function and the result from that is 
		being cast to 'fruits' type of data
*/
func createFruitsFromFile(fileName string) fruits {
	file, err := ioutil.ReadFile(fileName)

	if err != nil{
		fmt.Println("Error:", err)
	}	

	fmt.Println("Succesfull reading, returning fruits")
	
	return fruits(strings.Split(string(file), ","))
}
