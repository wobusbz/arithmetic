package tailStrings

import "testing"

func TestTailStrings(t *testing.T) {
	str := "hello  wuhaurou "

	t.Log(TailStrings(str))
}
