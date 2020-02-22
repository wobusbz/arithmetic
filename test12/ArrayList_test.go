package test12

import "testing"

func TestNewArrayList(t *testing.T) {
	var list List = NewArrayList()

	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	t.Log(list.String())
}
