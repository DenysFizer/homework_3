package main

import "fmt"

type animal interface {
	eatGetter
	fmt.Stringer //fixer add stringer interface
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

func (c cat) String() string {
	return "cat"
}

type dog struct {
	weight int
}

func (d dog) getFeed() float64 {
	absolute := d.weight / 5
	return float64(absolute * 10)
}

func (d dog) String() string {
	return "dog"
}

type cow struct {
	weight int
}

func (c cow) getFeed() float64 {
	return float64(c.weight * 25)
}
func (c cow) String() string { //fixed reciever
	return "cow"
}

func eatAll(feed eatGetter) float64 {
	korm := feed.getFeed()
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
	var foodForAll float64 //fixed hided counter to the function
	for _, t := range animals {
		eat := eatAll(t)
		foodForAll += eat
		fmt.Printf(" %s eat :%.f kg \n", t.String(), eat)
	}
	fmt.Printf(" Weight for all Farm: %.f kg\n", foodForAll)
}
