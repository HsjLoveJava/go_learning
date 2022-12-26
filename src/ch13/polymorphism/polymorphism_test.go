package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}

func WriteFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProgrammer := &GoProgrammer{}
	javaProgrammer := new(JavaProgrammer)
	WriteFirstProgrammer(goProgrammer)
	WriteFirstProgrammer(javaProgrammer)
}
