package main

type Flyer interface{
	Fly()
}

type Walker interface{
	Walk()
}

type Bird struct{

}

func(b *Bird) Fly(){
	println("i am bird,i can fly")
}
func(b *Bird) Walk(){
	println("i am bird,i can walk")
}

type Pag struct{

}
func(p *Pag)Walk(){
	println("i am pag,i can walk")
}
func main(){
	m := map[string]interface{}{
		"bird" : new(Bird),
		"pag" : new(Pag),
	}

	for _,v := range m{
		f,isFlyer := v.(Flyer)
		w,isWalk := v.(Walker)

		if isFlyer{
			f.Fly()
		}

		if isWalk{
			w.Walk()
		}
	}
}