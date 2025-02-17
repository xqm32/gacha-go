package gacha

import "math/rand/v2"

type Gacha struct {
	CharWish  CharWish
	CharPulls int
	CharAt    int
	CherAt    int
	WeapWish  WeapWish
	WeapPulls int
	WeapAt    int
	WeepAt    int

	OnCharUp   func(g *Gacha)
	OnChar     func(g *Gacha)
	OnCharSpec func(g *Gacha)
	OnCher     func(g *Gacha)
	OnWeapUp   func(g *Gacha)
	OnWeap     func(g *Gacha)
	OnWeapSpec func(g *Gacha)
	OnWeep     func(g *Gacha)
}

func (g *Gacha) SetPity(cp, wp int) *Gacha {
	g.CharWish.Pity, g.WeapWish.Pity = cp, wp
	return g
}

func (g *Gacha) PullCharsUp(n int) *Gacha {
	for n > 0 {
		g.CharAt = g.CharWish.Pity + 1
		g.CherAt = g.CharWish.Poty + 1
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
		} else if g.CharWish.Poty == 0 && g.OnCher != nil {
			g.OnCher(g)
		}
	}
	return g
}

func (g *Gacha) PullWeapsUp(n int) *Gacha {
	for n > 0 {
		g.WeapAt = g.WeapWish.Pity + 1
		g.WeepAt = g.WeapWish.Poty + 1
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
		} else if g.WeapWish.Poty == 0 && g.OnWeep != nil {
			g.OnWeep(g)
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
