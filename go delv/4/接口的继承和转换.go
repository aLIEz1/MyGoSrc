package main

import "fmt"

type Stu struct {
	name string
	id int
}
type Person interface {
	SayHi()
}
type Student interface {
	Person
	sing(irc string)
}

func(s *Stu) SayHi()  {
	fmt.Printf("student[%s %d]SayHiToYou\n",s.name,s.id)
}
func(s *Stu) sing(irc string)  {
	fmt.Printf("student[%s %d]Sing:%s\n",s.name,s.id,irc)
}
func main()  {
	var i Student
	s:=&Stu{"mike",666}
	i=s
	i.SayHi()
	i.sing("NiceToMeetYou")
	var iPro Student=&Stu{"mike",666}
	var i2 Person
	i2=iPro//超集可以转换为子集，反过来不可以
	println("i2=",i2)

	
}
