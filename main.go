package main

import (
	// "fmt"
	// cr "patterns/creational"
	"fmt"
	bh "patterns/behavioral"
)


func main() {
     // lfu := &bh.Lfu{}
     // cache := bh.Init(lfu)
     //
     // cache.Add("a", "1")
     // cache.Add("b", "2")
     //
     // cache.Add("c", "3")
     //
     // lru := &bh.Lru{}
     // cache.SetEvictionAlgo(lru)
     //
     // cache.Add("d", "4")
     //
     // fifo := &bh.Fifo{}
     // cache.SetEvictionAlgo(fifo)
     //
     // cache.Add("e", "5")

     // shirtItem := bh.NewItem("Nike shirt")
     // cus1 := &bh.Customer{}
     // cus2 := &bh.Customer{}
     // cus1.SetID("123")
     // cus2.SetID("456")
     //
     // shirtItem.Register(cus1)
     // shirtItem.Register(cus2)
     //
     //
     // shirtItem.UpdataAvailability()
     pizza := &bh.Margaretta{}
     pizzaWithCheese := &bh.CheeseTopping {
         Pizza: pizza,
     }

     pizzaWithVeggie := &bh.VeggieTopping {
         Pizza: pizza,
     }

     both := &bh.VeggieTopping{Pizza: pizzaWithCheese}

     fmt.Printf("with cheese only %d\n", pizzaWithCheese.GetPrice())
     fmt.Printf("with veggie only %d\n", pizzaWithVeggie.GetPrice())
     fmt.Printf("with both only %d\n", both.GetPrice())
}
