package behavioral

import "fmt"


type Subject interface {
    Register(observer Observer);
    Remove(observer Observer);
    Notify();
}

type Observer interface{
    update(string);
    getID() string;
}


type Item struct {
    observerList    []Observer
    Name            string;
    InStock         bool;
}

type Customer struct {
    id string
}

func (c *Customer) update(itemName string) {
    fmt.Printf("Sending email to customer %s for item %s: \n", c.id, itemName)
}

func (c *Customer) getID() string {
    return c.id
}

func (c *Customer) SetID(id string) {
    c.id = id
}

func NewItem (name string) *Item {
    return &Item{
        Name: name,
    }
}


func (i *Item) UpdataAvailability() {
    fmt.Printf("Item %s is now available\n", i.Name);
    i.InStock = true;
    i.notify();
}

func (i *Item) Register(observer Observer) {
    i.observerList = append(i.observerList, observer)
}

func (i *Item) Remove(observer Observer) {
    i.observerList = removeFromslice(i.observerList, observer)
}

func (i *Item) notify() {
    for _, observer := range i.observerList {
        observer.update(i.Name)
    }
}


func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
    observerListLength := len(observerList)
    for i, observer := range observerList {
        if observerToRemove.getID() == observer.getID() {
            observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
            return observerList[:observerListLength-1]
        }
    }
    return observerList
}
