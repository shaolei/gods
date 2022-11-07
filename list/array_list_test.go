package list

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	arrayList := NewArrayList()
	for i := 0; i < 5; i++ {
		arrayList.AddLast(i)
	}
	arrayList.Remove(3)
	arrayList.Add(1, 9)
	arrayList.AddFirst(100)
	arrayList.RemoveLast()
	arrayList.Set(1, 10)
	for i := 0; i < arrayList.Size(); i++ {
		v, _ := arrayList.Get(i)
		t.Logf("元素%d: %d", i, v)
	}
}
