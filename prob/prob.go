package main

import (
	"fmt"
	"sort"
)

var die = []int{1, 2, 3, 4, 5, 6}

func main() {
	for _, reroll := range []int{1, 2, 3, 4, 5, 6, 7} {
		getResults(reroll)
	}
}

func getResults(reroll int) {
	histograms := map[string]map[int]int{
		"non-professional":   {},
		"professional":       {},
		"human-professional": {},
	}
	// var results [][]int
	count := 0
	for _, a := range die {
		for _, b := range die {
			for _, c := range die {
				for _, d := range die {
					count++
					dice := []int{a, b, c, d}
					// results = append(results)
					n, p, m := prob(dice, reroll)
					histograms["non-professional"][n]++
					histograms["professional"][p]++
					histograms["human-professional"][m]++
				}
			}
		}
	}

	for k, v := range histograms {
		fmt.Println(k, "( reroll <", reroll, ")")
		cum := 0
		fmt.Printf("roll %9s   %6s   %6s\n", "chance", "prcnt", "cum")
		for i := 12; i > 1; i-- {
			cum += v[i]
			prcnt := float64(100*v[i]) / float64(count)
			cumPrcnt := float64(100*cum) / float64(count)
			fmt.Printf("%02d   %4d/%d   %6.2f   %6.2f\n", i, v[i], count, prcnt, cumPrcnt)
		}
		fmt.Println()
	}
}

func prob(results []int, reroll int) (int, int, int) {
	raw := results[0] + results[1]
	// fmt.Println("non-pro (sum of ", results[0], results[1], ") is ", results[0]+results[1])
	pro := pro(results[0], results[1], results[2], reroll)
	manPro := manPro(results[0], results[1], results[2], results[3], reroll)
	return raw, pro, manPro
}

func pro(a, b, c, reroll int) int {
	x := []int{a, b, c}
	sort.Ints(x[:2])
	// fmt.Println("sort (of ", a, b, c, ") is ", x[0], x[1])
	if x[0] < reroll {
		// fmt.Println("pro (best 2 of ", a, b, c, ") is ", x[0]+c)
		return x[1] + c
	}
	// fmt.Println("pro (best 2 of ", a, b, c, ") is ", a+b)
	return a + b
}

func manPro(a, b, c, d, reroll int) int {
	x := a
	if a < reroll {
		x = b
	}
	y := c
	if c < reroll {
		y = d
	}
	return x + y
}
