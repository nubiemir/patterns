package main

import (
	// "fmt"
	// cr "patterns/creational"
    bh "patterns/behavioral"
)


func main() {
     lfu := &bh.Lfu{}
     cache := bh.Init(lfu)

     cache.Add("a", "1")
     cache.Add("b", "2")

     cache.Add("c", "3")

     lru := &bh.Lru{}
     cache.SetEvictionAlgo(lru)

     cache.Add("d", "4")

     fifo := &bh.Fifo{}
     cache.SetEvictionAlgo(fifo)

     cache.Add("e", "5")
}
