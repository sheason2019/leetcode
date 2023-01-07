package solution130

import (
	"math"
	"sort"
)

func candy(ratings []int) int {
	n := len(ratings)
	// 收集评分
	min := math.MaxInt
	ratingMap := make(map[int][]int)
	candies := make([]int, n)
	candiesSum := 0
	for i, rating := range ratings {
		if rating < min {
			min = rating
		}
		if ratingMap[rating] == nil {
			ratingMap[rating] = []int{}
		}
		ratingMap[rating] = append(ratingMap[rating], i)
	}
	ratingSort := make([]int, len(ratingMap))
	i := 0
	for rating := range ratingMap {
		ratingSort[i] = rating
		i++
	}
	sort.Ints(ratingSort)

	// 然后从最低分开始遍历
	for _, rating := range ratingSort {
		indexes := ratingMap[rating]
		for _, index := range indexes {
			if rating == min {
				candies[index] = 1
				candiesSum = candiesSum + 1
			} else {
				candyCount := 1
				if index != 0 {
					if candies[index-1]+1 > candyCount && ratings[index-1] < rating {
						candyCount = candies[index-1] + 1
					}
				}
				if index != n-1 {
					if candies[index+1]+1 > candyCount && ratings[index+1] < rating {
						candyCount = candies[index+1] + 1
					}
				}
				candies[index] = candyCount
				candiesSum = candiesSum + candyCount
			}
		}
	}

	// fmt.Println("candies", candies)

	return candiesSum
}
