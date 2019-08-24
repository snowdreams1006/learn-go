package extend

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Skill() {
	fmt.Println("能文能武的宠物")
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Navigate() {
	fmt.Println("自带导航汪汪汪")
}

func (d *Dog) Speak() {
	fmt.Println("Wang")
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}

type Cat struct {
	p *Pet
}

func (c *Cat) Catch() {
	fmt.Println("老鼠天敌喵喵喵")
}

func TestDog(t *testing.T) {
	d := new(Dog)
	d.SpeakTo("Chao")
}

func TestExtendInstance(t *testing.T) {
	p := new(Pet)

	d := new(Dog)
	d.p = p

	// 自带导航汪汪汪
	d.Navigate()
	// 能文能武的宠物
	d.p.Skill()

	fmt.Println()

	c := new(Cat)
	c.p = p

	// 老鼠天敌喵喵喵
	c.Catch()
	// 能文能武的宠物
	c.p.Skill()
}
