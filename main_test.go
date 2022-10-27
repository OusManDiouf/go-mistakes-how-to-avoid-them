package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSome(t *testing.T) {

	assert.True(t, true)
}

func TestMergeSort(t *testing.T) {

	theArray := []int{14, 25, 35, 47, 11, 25, 82, 5, 1, 3}
	result := []int{1, 3, 5, 11, 14, 25, 25, 35, 47, 82}

	mergeSort(theArray)

	assert.Equal(t, result, theArray)

}

func TestMergeSortParallel(t *testing.T) {

	theArray := []int{14, 25, 35, 47, 11, 25, 82, 5, 1, 3}
	result := []int{1, 3, 5, 11, 14, 25, 25, 35, 47, 82}

	mergeSortUsingParallel(theArray)

	assert.Equal(t, result, theArray)

}
