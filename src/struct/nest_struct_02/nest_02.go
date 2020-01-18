package main
import (
	"fmt"
)

type Beego struct {
	Data map[string]interface{}
	num  int
	Tpl  string
}
type MainController struct {
	Beego
}

// func(c MainController) Get(){
// 	c.num = 10
// 	c.Tpl = "index.php"
// }

func (c *MainController) Get() {
	c.num = 10
	c.Tpl = "index.php"
	// c.Data["name"] = "lmm"
	// c.Data["age"] = 32
	c.Data = map[string]interface{}{
		"name":"",
		"age":23,
	}
}

func main() {
	var c MainController
	fmt.Println(c)
	c.Get()
	fmt.Println(c)
	fmt.Println(c.Data["name"],c.Data["age"])
}
