package q3

import (
	"math"
)

// GetPrimes 用于获取小于或者等于参数max的所有质数。
// 本函数是有的是Sieve of Eratosthenes算法。
// 该算法的存储数据结构最好使用双向链表。
func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	// 初始化一个bool slice用于标记下标所代表的自然数是否为质数
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}
