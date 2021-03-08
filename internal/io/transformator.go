package io

// Transformator is a function that receives a byte array and returns a transformed array
type Transformator func(input []byte) []byte

// NoOpTransformator returns the input
var NoOpTransformator = func(input []byte) []byte {
	return input
}
