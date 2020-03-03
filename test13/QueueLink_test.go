package test13

import (
	"io/ioutil"
	"path"
	"testing"
)

func testQueueStack(t *testing.T) {
	var node = &QueueStack{}
	for i := 0; i < 10; i++ {
		node.Push(i)
	}
	t.Log(node.Length())
	for data := node.Pop(); data != nil; data = node.Pop() {
		t.Log(data)
	}
}

func TestQueueStack_1(t *testing.T) {
	var node = &QueueStack{}
	node.Push("E:/")
	for {
		paths := node.Pop().(string)
		read, _ := ioutil.ReadDir(paths)
		t.Log(read)
		for _, finfo := range read {
			if finfo.IsDir() {
				node.Push(path.Join(paths + finfo.Name()))
			}
		}
		for data := node.Pop(); data != nil; data = node.Pop() {
			t.Log(data)
		}
		if node.Length() == 0 {
			break
		}
	}
}
