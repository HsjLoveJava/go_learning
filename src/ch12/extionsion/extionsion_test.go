package extionsion

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("...")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

//func (d *Dog) speak() {
//	d.p.speak()
//}
//
//func (d *Dog) speakTo(host string) {
//	d.p.speakTo(host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.speakTo("hushengjin")
}
