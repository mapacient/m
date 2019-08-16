package main

import (
	"container/heap"
	"fmt"
	"sort"
)

const (
	vv = true
)

//[0] for value [1] for original index
type is2 [][2]int

//[0] for current index [1] for delta
type is3 [][2]int

func (is *is2) Len() int {
	return len(*is)
}

func (is *is2) Less(a, b int) bool {
	return (*is)[a][0] < (*is)[b][0]
}

func (is *is2) Swap(a, b int) {
	(*is)[a], (*is)[b] = (*is)[b], (*is)[a]
}

func (is *is3) Len() int {
	return len(*is)
}

func (is *is3) Less(a, b int) bool {
	return (*is)[a][1] > (*is)[b][1]
}

func (is *is3) Swap(a, b int) {
	(*is)[a], (*is)[b] = (*is)[b], (*is)[a]
}

func (is *is3) Push(i interface{}) {
	*is = append(*is, i.([2]int))
}

func (is *is3) Pop() interface{} {
	r := (*is)[len(*is)-1]
	*is = (*is)[:len(*is)-1]
	return r
}

func c(sl []int) (r []int) {
	sl_22 := make(is2, len(sl))
	for i := 0; i < len(sl); i++ {
		sl_22[i][0] = sl[i]
		sl_22[i][1] = i
	}

	sort.Sort(&sl_22)
	fmt.Print(sl_22, "\n")
	sl_33 := make(is3, len(sl))
	for i := 0; i < len(sl); i++ {
		delta := i - sl_22[i][1]
		if delta < 0 {
			delta = -delta
		}
		sl_33[i][0] = i
		sl_33[i][1] = delta
	}
	heap.Init(&sl_33)


	for len(sl_33) > 0 {
		//I am just playing

		cu_sl_33 := heap.Pop(&sl_33).([2]int)
		//fmt.Print("sl_333333333333333333333", "       ",sl_33,"\n")
		//cu_sl_33[0] for current index in sl_22
		//src for index in sl_22
		src := cu_sl_33[0]
		//target for index in sl_22
		target := sl_22[src][1]
		//swap with right
		if src == target {
			break
		}
		if src < target {

			if sl_22[src+1][1] >= src && sl_22[src+1][1]-src+1 > target-src {
				if vv {
					fmt.Print(" move right continue\n")
				}
				continue
			}
			sl_22[src], sl_22[src+1] = sl_22[src+1], sl_22[src]
			//I operate + till >= right
			sl_22[src+1][0] = sl_22[src][0]
			cu_sl_33[0]++
			if vv {
				fmt.Print(sl_22, "##move right\n\n")
			}

		} else {
			if sl_22[src-1][1] <= src && src-sl_22[src-1][1]+1 > src-target{
				fmt.Print("move left continue\n")
				continue
			}
			sl_22[src], sl_22[src-1] = sl_22[src-1], sl_22[src]
			//I operate - till -= right
			sl_22[src-1][0] = sl_22[src][0]
			cu_sl_33[0]--
			if vv {
				fmt.Print(sl_22, "##move left\n\n")

			}
			cu_sl_33[1]--
		}

		//	heap.Push(&sl_33,cu_sl_33)

		sl_33 = sl_33[:0]
		for j := 0; j < len(sl_22); j++ {
			sl_33 = append(sl_33, [2]int{0, 0})
			//current index j and orginal index saved in [1]
			delta := j - sl_22[j][1]
			if delta < 0 {
				delta = -delta
			}
			sl_33[j][0] = j
			sl_33[j][1] = delta
		}
		heap.Init(&sl_33)

	}

	for _, v := range sl_22 {
		r = append(r, v[0])
	}

	return
}
