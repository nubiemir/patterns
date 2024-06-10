package behavioral

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

type Lru  struct {
}

type Lfu  struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo")
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru")
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu")
}

type Cache struct {
    Storage      map[string]string
    EvictionAlgo EvictionAlgo
    Capacity     int
    MaxCapacity  int
}


func Init (e EvictionAlgo) *Cache {
    storage := make(map[string]string)
    return &Cache{
        Storage: storage,
        EvictionAlgo: e,
        Capacity: 0,
        MaxCapacity: 2,
    }
}


func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
    c.EvictionAlgo = e
}

func (c *Cache) get(key string) {
    delete(c.Storage, key)
}

func (c *Cache) Evict() {
    c.EvictionAlgo.evict(c)
    c.Capacity--
}

func (c *Cache) Add(key, val string) {
    if c.Capacity == c.MaxCapacity {
        c.Evict()
    }
    c.Capacity++
    c.Storage[key] = val
}
