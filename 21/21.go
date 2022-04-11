package main

//Реализовать паттерн «адаптер» на любом примере.

import "fmt"

type duck interface {
	qrya()
}
type mallardDuck struct {
}

func (*mallardDuck) qrya() {
	fmt.Println("Mallard qrya")
}

type madagascarDuck struct {
}

func (*madagascarDuck) qrya() {
	fmt.Println("Доложите обстановку, Ковальский")
}

type mechaDuck struct {
}

func (*mechaDuck) mecha_qrya() {
	fmt.Println([]rune("Кря-Кря"))
}

type adaptMecha struct {
	roboDuck *mallardDuck
}

func (a *adaptMecha) likeMallard() {
	a.roboDuck.qrya()
}
func main() {
	malDuck := &mallardDuck{}
	malDuck.qrya()

	shkiper := &madagascarDuck{}
	shkiper.qrya()

	mecha := &mechaDuck{}
	mecha.mecha_qrya()

	mechaAdapt := &adaptMecha{roboDuck: malDuck}
	mechaAdapt.likeMallard()
}
