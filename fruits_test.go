package main

import (
	"os"
	"testing"
)

/*
	- For testing the name of the file has to be 'NameOfTestingFile_test.go, so go test command can read it.
*/

func TestNewFruits(t *testing.T){
	fruits := newFruits()

	if(len(fruits) != 12){
		t.Errorf("Expected a number of 12 fruits, but got %v",len(fruits))
	}
}

func TestWriteFileAndCreateFruitsFromFile(t *testing.T){
	os.Remove("_test_fruits.txt")

	fruits := newFruits()
	err := fruits.writeFile("_test_fruits.txt")
		
		if err != nil {
			t.Errorf("Error writing fruits data in file: %v", err)
		}

	fruitsFromFile := createFruitsFromFile("_test_fruits.txt")

		if len(fruitsFromFile) != 12 {
			t.Errorf("Expected a number of 12 fruits from the readed file, but got %v",len(fruits))
		}
}