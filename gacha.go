package gacha

import "math/rand/v2"

var (
	Char5Prob = []int{
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 660, 1260, 1860, 2460, 3060, 3660, 4260, 4860, 5460, 6060, 6660, 7260, 7860, 8460, 9060,
		9660, 10260,
	}
	Char4Prob = []int{510, 510, 510, 510, 510, 510, 510, 510, 5610, 10710}
)

type CharResult int

const (
	Char5Up CharResult = iota
	Char5Spec
	Char5
	Char4
	Char3
)

type CharWish struct {
	Pity5 int
	Guar  int
	Spec  int
	Pity4 int
}

func (w *CharWish) Pull() CharResult {
	return w.PullR(rand.N(9999)+1, rand.N(9999)+1)
}

func (w *CharWish) PullN(n int) []CharResult {
	res := make([]CharResult, n)
	for i := 0; i < n; i++ {
		res[i] = w.Pull()
	}
	return res
}

func (w *CharWish) PullR(r1, r2 int) CharResult {
	if r1 <= Char5Prob[w.Pity5] {
		if (w.Guar == 0 && 1 <= r2 && r2 <= 5000) || (w.Guar == 1) {
			w.Pity5, w.Guar, w.Spec, w.Pity4 = 0, 0, 0, w.Pity4+1
			return Char5Up
		} else if w.Guar == 0 && 5001 <= r2 && r2 <= 5500 {
			w.Pity5, w.Guar, w.Spec, w.Pity4 = 0, 0, 1, w.Pity4+1
			return Char5Spec
		} else {
			w.Pity5, w.Guar, w.Spec, w.Pity4 = 0, 1, 0, w.Pity4+1
			return Char5
		}
	} else if w.Pity4 >= 10 || r2 <= Char4Prob[w.Pity4] {
		w.Pity5, w.Spec, w.Pity4 = w.Pity5+1, 0, 0
		return Char4
	} else {
		w.Pity5, w.Spec, w.Pity4 = w.Pity5+1, 0, w.Pity4+1
		return Char3
	}
}

var (
	Weap5Prob = []int{
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 770, 1470, 2170, 2870, 3570, 4270,
		4970, 5670, 6370, 7070, 7770, 8470, 9170, 9870, 10570, 11270, 11970, 12670,
	}
	Weap4Prob = []int{600, 600, 600, 600, 600, 600, 600, 6600, 12600, 18600}
)

type WeapResult int

const (
	Weap5Up WeapResult = iota
	Weap5Spec
	Weap5
	Weap4
	Weap3
)

type WeapWish struct {
	Pity5 int
	Guar  int
	Spec  int
	Pity4 int
}

func (w *WeapWish) Pull() WeapResult {
	return w.PullR(rand.N(9999)+1, rand.N(9999)+1)
}

func (w *WeapWish) PullN(n int) []WeapResult {
	res := make([]WeapResult, n)
	for i := 0; i < n; i++ {
		res[i] = w.Pull()
	}
	return res
}

func (w *WeapWish) PullR(r1, r2 int) WeapResult {
	if r1 <= Weap5Prob[w.Pity5] {
		if (w.Guar == 0 && 1 <= r2 && r2 <= 3750) || (w.Guar == 1) {
			w.Pity5, w.Guar, w.Spec, w.Pity4 = 0, 0, 0, w.Pity4+1
			return Weap5Up
		} else if w.Guar == 0 && 3751 <= r2 && r2 <= 7500 {
			w.Pity5, w.Guar, w.Spec, w.Pity4 = 0, 1, 1, w.Pity4+1
			return Weap5Spec
		} else {
			w.Pity5, w.Guar, w.Spec, w.Pity4 = 0, 1, 0, w.Pity4+1
			return Weap5
		}
	} else if w.Pity4 >= 10 || r2 <= Weap4Prob[w.Pity4] {
		w.Pity5, w.Spec, w.Pity4 = w.Pity5+1, 0, 0
		return Weap4
	} else {
		w.Pity5, w.Spec, w.Pity4 = w.Pity5+1, 0, w.Pity4+1
		return Weap3
	}
}
