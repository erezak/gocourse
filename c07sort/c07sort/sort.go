/*
	The sort package provides interfaces and algorithms that allow sorting
*/
package c07sort

/*
	A list of functions that the collection must implement to be sortable
*/
type Sortable interface {
	// Implement this method to return whether the item in the first index is 
	// smaller than the item in the second index, to allow sorting
	IsFirstLessThanSecondByIndex(firstIndex int, secondIndex int) bool
	// Implement this method to return the length of the collection
	// (i.e. number of items)
	Length() int
	// Implement this method to swap the positions of the
	// first and second items accroding to their indexes
	Swap(firstIndex int, secondIndex int)
}

/*
	A naive implementation of bubble sort
*/
func BubbleSort(s Sortable) (rounds int) {
	length := s.Length()
	for {
		wasSwapped := false
		rounds++
		for i:=1; i < length ; i++ {
			if s.IsFirstLessThanSecondByIndex(i, i-1) {
				s.Swap(i, i-1)
				wasSwapped = true
			}
		}
		if !wasSwapped {
			break
		}
	}
	return rounds
}
