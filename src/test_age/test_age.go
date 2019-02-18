package main

import (
	"fmt"
	"time"
)

func GetTimeFromStrDate(date string)(year int){
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil{
		fmt.Println("出生日期解析错误!")
		return 0
	}
	year = d.Year()
	//month = int(d.Month())
	//day = d.Day()
	return  year
}

func GetZodiac(year int) (zodiac string) {
	if year <= 0{
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x== -11{
		zodiac = "鼠"
	}
	if x == 0{
		zodiac = "牛"
	}
	if x == 11 || x == -1{
		zodiac = "虎"
	}
	if x == 10 || x == -2{
		zodiac = "兔"
	}
	return 
}

func GetAge(year int) (age int){
	if year <= 0{
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

func main(){
	y := GetTimeFromStrDate("1986-06-18")
	fmt.Println(GetAge(y))
	fmt.Println(GetZodiac(y))
}