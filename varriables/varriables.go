package main

import "fmt"


func mR()(int, int){
	return 5,6
}
func mR2()(a int, b int){
	a=3
	b=2
	return
}
func main(){
	fmt.Println("hello Varriables")	

	var a int
	var b float32
	var c float64
	var h string
	a =3
	d := 5
	f := int32(6)
	var g int64=100
	fmt.Printf("%d %f %f %d %d %d %s\n",a,b,c,d,f,g,h)


	var a1,a2,a3 int
	a1 , a2 , a3 = 1,2,3
	fmt.Printf("%d\n",a1+a2+a3)
	var b1, b2,b3 string
	b1, b2, b3 = "zonieed" , "Hossain", "Badhon"
	fmt.Printf("%s\n", b1+b2+b3)
	c1 , c2 := 4, "zonieed"
	fmt.Printf("%d %s\n",c1,c2)

	x1,y1 := mR()
	x2,y2 := mR2()
	fmt.Printf("print all the value %d %d %d %d\n",x1,x2,y1,y2)
	fmt.Printf("multiply %d\n",x1*x2)
}