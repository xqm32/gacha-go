package gacha

import (
	"math/rand"
)

var (
	U5C_W = []int{
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60,
		60, 660, 1260, 1860, 2460, 3060, 3660, 4260, 4860, 5460, 6060, 6660, 7260, 7860, 8460, 9060,
		9660, 10260,
	}
	U4C_W = []int{510, 510, 510, 510, 510, 510, 510, 510, 5610, 10710}
	U5W_W = []int{
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 770, 1470, 2170, 2870, 3570, 4270,
		4970, 5670, 6370, 7070, 7770, 8470, 9170, 9870, 10570, 11270, 11970, 12670,
	}
	U4W_W = []int{600, 600, 600, 600, 600, 600, 600, 6600, 12600, 18600}
)

type Gacha struct {
	// records
	Pulls     int // Intertwined fates
	Stars     int // Starglitters
	CharsUp   int // 5 star characters up
	CharsDown int // 5 star characters down
	WeapsUp   int // 5 star weapons up
	WeapsDown int // 5 star weapons down
	// states
	U5cPity int // Up 5 star character pity
	U5cGuar int // Up 5 star up character guarantees
	U4cPity int // Up 4 star character pity
	U5wPity int // Up 5 star weapon pity
	U5wGuar int // Up 5 star up weapon guarantees
	U4wPity int // Up 4 star weapon pity
	// events
	OnCharUp        func(g *Gacha) // 5 star character up event
	OnCharDown      func(g *Gacha) // 5 star character down event
	OnCharLight     func(g *Gacha) // 5 star character light event
	OnWeapUp        func(g *Gacha) // 5 star weapon up event
	OnAnotherWeapUp func(g *Gacha) // Another 5 star weapon up event
	OnWeapDown      func(g *Gacha) // 5 star weapon down event
}

func (g *Gacha) PullChar() *Gacha {
	g.Pulls, g.U5cPity, g.U4cPity = g.Pulls+1, g.U5cPity+1, g.U4cPity+1
	if (rand.Intn(9999) + 1) <= U5C_W[g.U5cPity-1] {
		n := rand.Intn(9999) + 1
		switch {
		case (g.U5cGuar == 0 && 1 <= n && n <= 5000) || (g.U5cGuar == 1):
			if g.OnCharUp != nil {
				g.OnCharUp(g)
			}
			g.CharsUp, g.U5cGuar = g.CharsUp+1, 0
		case (g.U5cGuar == 0 && 5001 <= n && n <= 5500):
			if g.OnCharLight != nil {
				g.OnCharLight(g)
			}
			g.CharsUp, g.U5cGuar = g.CharsUp+1, g.U5cGuar+1
		default:
			if g.OnCharDown != nil {
				g.OnCharDown(g)
			}
			g.CharsDown, g.U5cGuar = g.CharsDown+1, g.U5cGuar+1
		}
		g.Stars, g.U5cPity = g.Stars+5, 0
	} else if g.U4cPity >= 10 || (rand.Intn(9999)+1) <= U4C_W[g.U4cPity-1] {
		g.Stars, g.U4cPity = g.Stars+2, 0
	}
	return g
}

func (g *Gacha) PullWeap() *Gacha {
	g.Pulls, g.U5wPity, g.U4wPity = g.Pulls+1, g.U5wPity+1, g.U4wPity+1
	if (rand.Intn(9999) + 1) <= U5W_W[g.U5wPity-1] {
		n := rand.Intn(9999) + 1
		switch {
		case (g.U5wGuar == 0 && 1 <= n && n <= 3750) || g.U5wGuar == 1:
			if g.OnWeapUp != nil {
				g.OnWeapUp(g)
			}
			g.WeapsUp, g.U5wGuar = g.WeapsUp+1, 0
		case (g.U5wGuar == 0 && 3751 <= n && n <= 7500):
			if g.OnAnotherWeapUp != nil {
				g.OnAnotherWeapUp(g)
			}
			g.WeapsDown, g.U5wGuar = g.WeapsDown+1, g.U5wGuar+1
		default:
			if g.OnWeapDown != nil {
				g.OnWeapDown(g)
			}
			g.WeapsDown, g.U5wGuar = g.WeapsDown+1, g.U5wGuar+1
		}
		g.Stars, g.U5wPity = g.Stars+5, 0
	} else if g.U4wPity >= 10 || (rand.Intn(9999)+1) <= U4W_W[g.U4wPity-1] {
		g.Stars, g.U4wPity = g.Stars+2, 0
	}
	return g
}

func (g *Gacha) PullChars(pulls int) *Gacha {
	pulls += g.Pulls
	for g.Pulls < pulls {
		g.PullChar()
	}
	return g
}

func (g *Gacha) PullWeaps(pulls int) *Gacha {
	pulls += g.Pulls
	for g.Pulls < pulls {
		g.PullWeap()
	}
	return g
}

func (g *Gacha) PullCharsUp(ups int) *Gacha {
	ups += g.CharsUp
	for g.CharsUp < ups {
		g.PullChar()
	}
	return g
}

func (g *Gacha) PullWeapsUp(ups int) *Gacha {
	ups += g.WeapsUp
	for g.WeapsUp < ups {
		g.PullWeap()
	}
	return g
}

func (g *Gacha) PullUp(charsUp int, weapsUp int) *Gacha {
	return g.PullCharsUp(charsUp).PullWeapsUp(weapsUp)
}
