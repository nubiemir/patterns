package behavioral

type IPizza interface {
    GetPrice() int
}


type Margaretta struct {
}

func (m *Margaretta) GetPrice() int {
    return 15
}


type CheeseTopping struct {
    Pizza IPizza
}

type VeggieTopping struct {
    Pizza IPizza
}



func (c *CheeseTopping) GetPrice() int {
    price := c.Pizza.GetPrice()
    return 5 * price
}

func (v *VeggieTopping) GetPrice() int {
    price := v.Pizza.GetPrice()
    return 10 * price
}
