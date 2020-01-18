package main

type Servicer interface{
	Log()
	Start()
}

type Logger struct{

}

func(l *Logger) Log(){
	println("this is log")
}

type GameServer struct{
	Logger
}

func(g *GameServer) Start(){
	println("this is start")
}

func main(){
	var s Servicer
	l := new(Logger)
	g := new(GameServer)
	l.Log()
	g.Start()

	//cannot use g (type *GameServer) as type Servicer in assignment:
	//*GameServer does not implement Servicer (missing Log method)
	s = g

	s.Log()
	s.Start()
}