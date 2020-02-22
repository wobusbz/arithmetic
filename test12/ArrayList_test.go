package test12

import "testing"

func TestNewArrayList(t *testing.T) {
	var arrayList = NewArrayList()

	arrayList.Append(2)
	arrayList.Append(2)
	arrayList.Append(2)
	arrayList.Append(2)
	arrayList.Append(2)

	for i := 0; i < 20; i++ {
		arrayList.Insert(i, "cx")
	}

	t.Log(arrayList.String())
}
