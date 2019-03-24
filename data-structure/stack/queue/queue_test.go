package queue

import "testing"

func TestCollection(t *testing.T) {
	collection := new(Collection)

	t.Run("Test Add", func(t *testing.T){
		collection.Add("a")
		if collection.Size() != 1 {
			t.Errorf("first added value must return a size of 1 but got %d\n", collection.Size())
		}
	})

	t.Run("Test Pop", func(t *testing.T){
		v := collection.Pop()
		if v != "a" {
			t.Errorf("Expecting value 'a' but got %s\n", v)
		}
		if collection.Size() != 0 {
			t.Errorf("after popping out the item that as been queued the size must be 0 but got %d\n", collection.Size())
		}
	})

	t.Run("Test Is Empty", func(t *testing.T){
		if !collection.IsEmpty() {
			t.Errorf("expecting to be empty")
		}
	})
}
