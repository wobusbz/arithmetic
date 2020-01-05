package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CalcRunTime(t string, f func(intArr []int, s1, s2 int), intArr []int, s1, s2 int) {
	s := time.Now()
	f(intArr, s1, s2)
	e := time.Since(s)
	fmt.Printf("%s 一共消耗时间: %v\n", t, e)
}

func CalcRunTime1(t string, f func(intArr []int), intArr []int) {
	s := time.Now()
	f(intArr)
	e := time.Since(s)
	fmt.Printf("%s 一共消耗时间: %v\n", t, e)
}

func CalcRunTime2(t string, f func(n int) int, n int) {
	s := time.Now()
	f(n)
	e := time.Since(s)
	fmt.Printf("%s 一共消耗时间: %v\n", t, e)
}

func RandList(n int) (list []int) {
	for i := 0; i < n; i++ {
		val := rand.Intn(1000000000)
		list = append(list, val)
	}
	return
}

var (
	IntArr = []int{10000, 199, 12, 5, 7, 22, 6, 100, 99, 33}
)
