package main

import "fmt"

type TongsOfPatty interface {
	String() string
}

type Patty interface {
	GetOneTongs() TongsOfPatty
}

type BigBurger struct {
	val string
}

func (b *BigBurger) PutPatty(patty Patty) {
	tongs := patty.GetOneTongs()
	b.val += tongs.String()
}

func (b *BigBurger) String() string {
	return "GoBlockBurger " + b.val
}

type FattyPatty struct {
}

func (j *FattyPatty) GetOneTongs() TongsOfPatty {
	return &DishesofFattyPatty{}
}

type ShittyPatty struct {
}

func (j *ShittyPatty) GetOneTongs() TongsOfPatty {
	return &TongsOfShittyPatty{}
}

type CalamariSashimi struct {
}

func (j *CalamariSashimi) GetOneTongs() TongsOfPatty {
	return &DishOfCalamariSashimi{}
}

type DishesofFattyPatty struct {
}

func (s *DishesofFattyPatty) String() string {
	return "+ Maguroslice"
}

type DishOfCalamariSashimi struct {
}

func (s *DishOfCalamariSashimi) String() string {
	return "+ CalamaruSlice"
}

type TongsOfShittyPatty struct {
}

func (s *TongsOfShittyPatty) String() string {
	return "+ FattyPatty"
}

func main() {
	BigBurger := &BigBurger{}
	Patty := &ShittyPatty{}
	BigBurger.PutPatty(Patty)

	fmt.Println(BigBurger)
}
