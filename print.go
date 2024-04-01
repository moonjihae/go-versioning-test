package x

import "fmt"

type TestType struct {
	Name string
}

func (t *TestType) PrintName() {
	fmt.Printf("%s in version v2.0.0", t.Name)
}
