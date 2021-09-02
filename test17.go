package main

import "fmt"

type Monkey struct {
	Name string
}

func (m Monkey) climbing()  {
	fmt.Println(m.Name, "生来会爬树...")
}

//type BirdAble interface {
//	Flying()
//}
//
//type FishAble interface {
//	Swimming()
//}

type LittleMonkey struct {
	Monkey
}

func (lm LittleMonkey) Flying() {
	fmt.Println(lm.Name, "通过学习会飞翔...")
}

func (lm LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "通过学习会游泳...")
}


func main() {
	littleMonkey := LittleMonkey{
		Monkey{"悟空"},
	}
	littleMonkey.climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()
}