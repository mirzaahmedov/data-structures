package list

import "testing"

func TestLinkedListAdd(t *testing.T){
	linkedList := NewLinkedList()

	linkedList.Add(10)
	linkedList.Add(20)

	if linkedList.Head().Val != 10 || linkedList.Head().Next.Val != 20 {
		t.Fatal("Should add new element to the list")
	}
}

func TestLinkedListSize(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Add(20)

	if linkedList.Size() != 1 {
		t.Fatal("Should update list size when new nodes added")
	}
}

func TestLinkedListForEach(t *testing.T) {
	items := [3]int{ 2, 5, 10 }

	linkedList := NewLinkedList()

	for i := range items {
		linkedList.Add(items[i])
	}

	count := 0
	linkedList.ForEach(func(i int) {
		if items[count] != i {
			t.Fatal("Should iterate list")
		}
		count++
	})

	if count != len(items) {
		t.Fatal("Should iterate list")
	}
}