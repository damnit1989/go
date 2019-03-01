package package_one

import "fmt"

const(
	LENGHT =  10
	WEIGHT = 20
)

var A int = 10

type PackageOne struct{
	Name string
	Age int
}

func(p PackageOne) GetName() string{
	return p.Name
}

func(p PackageOne) GetAge() int{
	return p.Age
}

func Info() int{
	return 2
}

func Main(){
    fmt.Printf("xxx\n")
}
