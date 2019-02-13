package main

func sum(a int)(sum int)  {
	sum=0
	for i:=1;i<=a ;i++  {
		sum=sum+i
	}
	return
}
func sum02(a int) int {
	if a==1 {
		return 1
	}
	return a+sum02(a-1)

}
func sum03(a int) int {
	if a==100 {
		return 100
	}
	return a+sum03(a+1)

}
func main()  {
	//a:=sum(100)
	//a:=sum02(100)
	a:=sum03(1)
	println("sum=",a)

}