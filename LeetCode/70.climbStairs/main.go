package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}


func climbStairs(n int) int {
    var dps = make([]int,n+1)
    dps[0] = 1
    dps[1] = 1
    for i := 2; i <= n; i ++ {
        dps[i] = dps[i-1]+dps[i-2]
    }
    return dps[n]
}
