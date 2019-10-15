//原接口：安卓充电线
//目标：苹果接口
//适配方式：安卓-ios转换器

package main

import "fmt"

type Target interface{
	Power()
}

type Andro struct{
	Name string
}
func (a Andro) Diff() string{
	return a.Name
}

type Ios struct{
	Name  string
}

func (i Ios) Power(){
	fmt.Println(i.Name)
}

type Adapt struct{
	Andro
}

func (a Adapt) Power(){
	fmt.Println(a.Diff())
}

func info(t Target){
	t.Power()
}
func main(){
	ios := Ios{"one"}
	info(ios)

	andro := Andro{"two"}
	ad := Adapt{andro}
	info(ad)
}