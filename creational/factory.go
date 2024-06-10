package creational

import "fmt"


type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}


type Gun struct {
	name string
	power int
}


type ak47 struct {
	Gun
}


type musket struct {
	Gun
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) SetPower(power int) {
	g.power = power 
}

func (g *Gun) GetName() string {
	return g.name 
}

func (g *Gun) GetPower() int {
	return g.power 
}

func NewAk47() IGun {
	return &ak47{
		Gun: Gun {
			name: "Ak47.....",
			power: 8,
		},
	}
}

func NewMusket() IGun {
	return &musket{
		Gun: Gun {
			name: "Musket.....",
			power: 8,
		},
	}
}


func NewGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil	
	}

	if gunType == "musket" {
		return NewMusket(), nil	
	}

	return nil, fmt.Errorf("Wrong gun passed: %s", gunType)
}