package main

import "testing"

//Simple test by calling Add function and expecting a result
func TestAdd(t *testing.T) {

	result := Add(1, 2)
	if result != 3 {
		t.Error("Existing 1 + 2 to equal 3")
	}
}

//struct that defines what values we pass
//as params and what we expect back
type testTableAdd struct {
	x        int
	y        int
	expected int
}

//list of inputs and results. this tests several variations
//of inputs
var testingArray = []testTableAdd{
	{1, 1, 2},
	{2, 2, 4},
	{4, 4, 8},
}

//executes the table / struct driven test
func TestTableAdd(t *testing.T) {

	//iterate over test array
	for _, test := range testingArray {

		//call Add and get the result
		result := Add(test.x, test.y)

		//compare the result to expected. return error if failed
		if result != test.expected {
			t.Error("Testing failed")
		}
	}

}
