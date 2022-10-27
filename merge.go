package main

import "sync"

func mergeSort(theArray []int) {

	nElems := len(theArray)
	workspace := make([]int, nElems)
	recMergeSort(theArray, workspace, 0, nElems-1)
}
func recMergeSort(theArray []int, workspace []int, lowerBound int, upperBound int) {
	if lowerBound == upperBound {
		return
	} else {
		mid := (lowerBound + upperBound) / 2

		recMergeSort(theArray, workspace, lowerBound, mid)
		recMergeSort(theArray, workspace, mid+1, upperBound)
		merge(theArray, workspace, lowerBound, mid+1, upperBound)
	}
}

func mergeSortUsingParallel(theArray []int) {

	nElems := len(theArray)
	workspace := make([]int, nElems)
	recMergeSortUsingParallel(theArray, workspace, 0, nElems-1)
}

func recMergeSortUsingParallel(theArray []int, workspace []int, lowerBound int, upperBound int) {
	if lowerBound == upperBound {
		return
	}
	mid := (lowerBound + upperBound) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		recMergeSortUsingParallel(theArray, workspace, lowerBound, mid)
	}()

	go func() {
		defer wg.Done()
		recMergeSortUsingParallel(theArray, workspace, mid+1, upperBound)
	}()
	wg.Wait()
	merge(theArray, workspace, lowerBound, mid+1, upperBound)
}

func merge(theArray []int, workspace []int, lowerPtr int, highPtr int, upperBound int) {
	j := 0
	lowerBound := lowerPtr
	mid := highPtr - 1
	n := upperBound - lowerBound + 1

	for lowerPtr <= mid && highPtr <= upperBound {
		if theArray[lowerPtr] < theArray[highPtr] {
			workspace[j] = theArray[lowerPtr]
			j++
			lowerPtr++
		} else {
			workspace[j] = theArray[highPtr]
			j++
			highPtr++
		}
	}

	for lowerPtr <= mid {
		workspace[j] = theArray[lowerPtr]
		j++
		lowerPtr++
	}
	for highPtr <= upperBound {
		workspace[j] = theArray[highPtr]
		j++
		highPtr++
	}

	for i := 0; i < n; i++ {
		theArray[lowerBound+i] = workspace[i]
	}
}
