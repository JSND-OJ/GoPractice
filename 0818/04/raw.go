package main

import "fmt"

type DishesOfSashimi interface {
	String() string
}

type SashimiSlice interface {
	GetOneDish() DishesOfSashimi
}

type FreshSashimi struct {
	val string
}

func (b *FreshSashimi) PutSashimiSlice(sashimislice SashimiSlice) {
	dishes := sashimislice.GetOneDish()
	b.val += dishes.String()
}

func (b *FreshSashimi) String() string {
	return "GoBlockSashimi " + b.val
}

type MaguroSashimi struct {
}

func (j *MaguroSashimi) GetOneDish() DishesOfSashimi {
	return &DishesofMaguroSashimi{}
}

type SalmonSashimi struct {
}

func (j *SalmonSashimi) GetOneDish() DishesOfSashimi {
	return &DishOfSalmonSashimi{}
}

type CalamariSashimi struct {
}

func (j *CalamariSashimi) GetOneDish() DishesOfSashimi {
	return &DishOfCalamariSashimi{}
}

type DishesofMaguroSashimi struct {
}

func (s *DishesofMaguroSashimi) String() string {
	return "+ MaguroPlate"
}

type DishOfSalmonSashimi struct {
}

func (s *DishOfSalmonSashimi) String() string {
	return "+ SalmonPlate"
}

type DishOfCalamariSashimi struct {
}

func (s *DishOfCalamariSashimi) String() string {
	return "+ CalamariPlate"
}

func main() {
	FreshSashimi := &FreshSashimi{}
	SashimiSlice := &SalmonSashimi{}
	FreshSashimi.PutSashimiSlice(SashimiSlice)

	fmt.Println(FreshSashimi)
}
