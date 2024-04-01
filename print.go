package x

import "fmt"

type TestType struct {
	Name string
}

func (t *TestType) PrintName() {
	fmt.Printf("%s in version v1.3.0", t.Name)
}
