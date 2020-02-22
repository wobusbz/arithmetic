package test12

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) ([]interface{}, error)
	Set(index int, newVal interface{}) error
	Insert(index int, newVal interface{}) error
	Append(newVal interface{}) error
	Clear()
	Delete(index int) error
	String() string
}

type ArrayList struct {
	DataSoure []interface{}
	IndexSize int
}

func NewArrayList() *ArrayList {
	obj := new(ArrayList)
	obj.DataSoure = make([]interface{}, 0, 10)
	obj.IndexSize = 0
	return obj
}
func (list *ArrayList) Size() int {
	return list.IndexSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 {
		return nil, errors.New("数组下标")
	}
	return list.DataSoure[list.IndexSize], nil
}

func (list *ArrayList) checkIsFull() {
	if list.IndexSize == cap(list.DataSoure) {
		newDataSoure := make([]interface{}, 2*list.IndexSize, 2*list.IndexSize)
		copy(newDataSoure, list.DataSoure)
		list.DataSoure = newDataSoure
	}
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index == list.IndexSize {
		return errors.New("数组下标越界")
	}
	list.DataSoure[index] = newVal
	return nil
}

func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index == list.IndexSize {
		return errors.New("数组下标越界")
	}
	list.checkIsFull()
	list.DataSoure = list.DataSoure[:list.IndexSize+1]
	for i := list.IndexSize; i > index; i-- {
		list.DataSoure[i] = list.DataSoure[i-1]
	}
	list.DataSoure[index] = newVal
	list.IndexSize++
	return nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.DataSoure = append(list.DataSoure, newVal)
	list.IndexSize++
}

func (list *ArrayList) Clear() {
	list.DataSoure = make([]interface{}, 0, 10)
	list.IndexSize = 0
}

func (list *ArrayList) Delete(index int) error {
	if index < 0 || index == list.IndexSize {
		return errors.New("数组下标越界")
	}
	list.DataSoure = append(list.DataSoure[:index], list.DataSoure[index+1:])
	list.IndexSize--
	return nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.DataSoure)
}
