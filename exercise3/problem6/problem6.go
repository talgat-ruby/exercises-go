package problem6

import "fmt"

type Animal struct {
    name    string
    legsNum int
}

type Insect struct {
    name    string
    legsNum int
}

func (a *Animal) GetLegsNum() int {
    return a.legsNum
}

func (i *Insect) GetLegsNum() int {
    return i.legsNum
}


func sumOfAllLegsNum(animalsAndInsects ...interface{}) int {
    totalLegs := 0
    for _, item := range animalsAndInsects {
        switch v := item.(type) {
        case *Animal:
            totalLegs += v.GetLegsNum()
        case *Insect:
            totalLegs += v.GetLegsNum()
        }
    }
    return totalLegs
}

func main() {
    horse := &Animal{name: "horse", legsNum: 4}
    kangaroo := &Animal{name: "kangaroo", legsNum: 2}
    ant := &Insect{name: "ant", legsNum: 6}
    spider := &Insect{name: "spider", legsNum: 8}

    totalLegs := sumOfAllLegsNum(horse, kangaroo, ant, spider)
    fmt.Println("Total number of legs:", totalLegs) // Вывод: 20
}
