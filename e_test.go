package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_c(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sl := make([]int, 7)
	for j := 0; j < 7; j++ {
		for i := 0; i < 7; i++ {
			sl[i] = rand.Intn(7)
		}

		sl1:=c(sl)
		t.Log(sl)
		t.Log(sl1)
		r:=0
		for i:=0;i<7;i++{
			d:=sl[i]-sl1[i]
		if d>=0{
			r+=d
		}else{
			r-=d
		}
		}
		t.Log(r)
	}
}
