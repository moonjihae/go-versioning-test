package x

import "fmt"

type TestType struct {
	Name string
}

func (t *TestType) PrintName() {
	fmt.Println(t.Name)
}
