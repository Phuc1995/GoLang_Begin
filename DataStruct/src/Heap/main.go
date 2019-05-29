package main

// integerHeap a type
type IntegerHeap []int
// IntegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i,j int) bool {
	return iheap[i]<iheap[j]
}

// IntegerHeap method -swaps the element of i to j index

