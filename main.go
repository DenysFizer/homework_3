package main

import "fmt"

type animal interface {
	nameGetter
	eatGetter
}
type nameGetter interface {
	getName() string
}
type eatGetter interface {
	getFeed() float64
}

type cat struct {
	weight int
}

func (c cat) getFeed() float64 {
	return float64(c.weight * 7)
}

func (c cat) getName() string {
	return "cat"
}

type dog struct {
	weight int
}

func (d dog) getFeed() float64 {
	absolute := d.weight / 5
	return float64(absolute * 10)
}

func (d dog) getName() string {
	return "dog"
}

type cow struct {
	weight int
}

func (cow cow) getFeed() float64 {
	return float64(cow.weight * 25)
}
func (cow cow) getName() string {
	return "cow"
}

var WeightallFeed float64

func eatAll(feed eatGetter) float64 {
	korm := feed.getFeed()
	WeightallFeed += korm
	return korm
}
func main() {
	animals := []animal{
		cat{weight: 15},
		dog{weight: 25},
		cow{weight: 100},
		dog{weight: 40},
		cow{weight: 145},
		cat{weight: 55},
		dog{weight: 75},
	}
	for _, t := range animals {
		eat := eatAll(t)
		fmt.Printf(" %s eat :%.f kg \n", t.getName(), eat)
	}
	fmt.Printf(" Weight for all Farm: %.f kg\n", WeightallFeed)
}
