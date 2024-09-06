package gacha

import (
	"math/rand/v2"
)

var (
	CharProb = []int{
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 660, 1260, 1860, 2460, 3060, 3660, 4260, 4860, 5460, 6060, 6660, 7260, 7860, 8460, 9060,
		9660, 10260,
	}
	CherProb = []int{510, 510, 510, 510, 510, 510, 510, 510, 5610, 10710}
	WeapProb = []int{
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 770, 1470, 2170, 2870, 3570, 4270,
		4970, 5670, 6370, 7070, 7770, 8470, 9170, 9870, 10570, 11270, 11970, 12670,
	}
	WeepProb = []int{600, 600, 600, 600, 600, 600, 600, 6600, 12600, 18600}
)

type CharWish struct {
	Pity int
	Guar int
	Spec int
	Poty int
}

func (w *CharWish) Pull(r1, r2 int) {
	if r1 <= CharProb[w.Pity] {
		if (w.Guar == 0 && 1 <= r2 && r2 <= 5000) || (w.Guar == 1) {
			w.Pity, w.Guar, w.Spec, w.Poty = 0, 0, 0, w.Poty+1
		} else if w.Guar == 0 && 5001 <= r2 && r2 <= 5500 {
			w.Pity, w.Guar, w.Spec, w.Poty = 0, 0, 1, w.Poty+1
		} else {
			w.Pity, w.Guar, w.Spec, w.Poty = 0, 1, 0, w.Poty+1
		}
	} else if w.Poty >= 10 || r2 <= CherProb[w.Poty] {
		w.Pity, w.Spec, w.Poty = w.Pity+1, 0, 0
	} else {
		w.Pity, w.Spec, w.Poty = w.Pity+1, 0, w.Poty+1
	}
}

type WeapWish struct {
	Pity int
	Guar int
	Spec int
	Poty int
}

func (w *WeapWish) Pull(r1, r2 int) {
	if r1 <= WeapProb[w.Pity] {
		if (w.Guar == 0 && 1 <= r2 && r2 <= 3750) || (w.Guar == 1) {
			w.Pity, w.Guar, w.Spec, w.Poty = 0, 0, 0, w.Poty+1
		} else if w.Guar == 0 && 3751 <= r2 && r2 <= 7500 {
			w.Pity, w.Guar, w.Spec, w.Poty = 0, 1, 1, w.Poty+1
		} else {
			w.Pity, w.Guar, w.Spec, w.Poty = 0, 1, 0, w.Poty+1
		}
	} else if w.Poty >= 10 || r2 <= WeepProb[w.Poty] {
		w.Pity, w.Spec, w.Poty = w.Pity+1, 0, 0
	} else {
		w.Pity, w.Spec, w.Poty = w.Pity+1, 0, w.Poty+1
	}
}

type Gacha struct {
	CharWish  CharWish
	CharPulls int
	CharAt    int
	WeapWish  WeapWish
	WeapPulls int
	WeapAt    int

	OnCharUp   func(g *Gacha)
	OnChar     func(g *Gacha)
	OnCharSpec func(g *Gacha)
	OnWeapUp   func(g *Gacha)
	OnWeap     func(g *Gacha)
	OnWeapSpec func(g *Gacha)
}

func (g *Gacha) SetPity(cp, wp int) *Gacha {
	g.CharWish.Pity, g.WeapWish.Pity = cp, wp
	return g
}

func (g *Gacha) PullCharsUp(n int) *Gacha {
	for n > 0 {
		g.CharAt = g.CharWish.Pity + 1
		r1, r2 := rand.N(9999)+1, rand.N(9999)+1
		g.CharWish.Pull(r1, r2)
		g.CharPulls += 1
		if g.CharWish.Pity == 0 {
			if g.CharWish.Guar == 0 {
				n -= 1
				if g.CharWish.Spec == 0 && g.OnCharUp != nil {
					g.OnCharUp(g)
				} else if g.OnCharSpec != nil {
					g.OnCharSpec(g)
				}
			} else if g.OnChar != nil {
				g.OnChar(g)
			}
		}
	}
	return g
}

func (g *Gacha) PullWeapsUp(n int) *Gacha {
	for n > 0 {
		g.WeapAt = g.WeapWish.Pity + 1
		r1, r2 := rand.N(9999)+1, rand.N(9999)+1
		g.WeapWish.Pull(r1, r2)
		g.WeapPulls += 1
		if g.WeapWish.Pity == 0 {
			if g.WeapWish.Guar == 0 {
				n -= 1
				if g.OnWeapUp != nil {
					g.OnWeapUp(g)
				}
			} else if g.WeapWish.Spec == 0 && g.OnWeapSpec != nil {
				g.OnWeapSpec(g)
			} else if g.OnWeap != nil {
				g.OnWeap(g)
			}
		}
	}
	return g
}

func (g *Gacha) PullUp(cn, wn int) *Gacha {
	return g.PullCharsUp(cn).PullWeapsUp(wn)
}

func (g *Gacha) Pulls() int {
	return g.CharPulls + g.WeapPulls
}
