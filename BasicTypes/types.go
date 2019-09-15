package main

import(
	"fmt"
	"unsafe"
	"strconv"
	"log"

)
func main(){
	//bool can either be true or false
	var b bool= true 
	fmt.Printf("b is : '%v'\n", b)
	b = false
	fmt.Printf("b is : '%v'\n", b)
	// Let's try to save bool value = 5
	// b=5 // in this line error given it's not type of bool
	

	// now assign no value that means null just decleard bool value and print
	var b2 bool
	fmt. Printf("None value of b2 is : '%v'\n", b2) // if we print that it's value will be false becasuse of null value

	// bool size of varriable
	b3 := true
	fmt.Printf("size of bool : %d\n", unsafe.Sizeof(b3)) // NOTE: U must import the unsafe

	var i1 int = -48
	fmt.Printf("value of i1 : %s\n",  i1) // if we print it's says like %!s(int = -48)
	// if we want to convert this int to string we can use strconv.Itoa 
	// NOTE: U must import the strconv 
	fmt.Printf("Converted value of i1 : %s\n", strconv.Itoa(i1))

	// if we Printf int32/int64 value 
	var i2 int32 = 111
	var i3 int64 = 112
	fmt.Printf("value of int32 and int 64 is : %d %d \n", i2, i3) // it's normally printed the vaule
	// Now convert this 
	fmt.Printf("Converted value of int32 is : %s \n", strconv.Itoa(int(i2))) // Remember that is strconv can convert only one varriable if it's syntex like strconv.Itoa (int(then value)) 
	
	// Now convert int to string this into Sprintf
	var a1 int = -56
	var a2 int32 =67
	fmt.Sprintf("%d",a1)// if we normally do that nothing will happand
	fmt.Sprintf("%d",a2)
	// now if we store the value into a string then printf then it's show the value 
	s1:=fmt.Sprintf("%d",a1)
	s2:=fmt.Sprintf("%d",a2)
	fmt.Printf("%s %s\n", s1,s2)

	// Convert string to int strconv.Atoi
	s:="-48"
	fmt.Printf("%s\n",s) // it's print -48 it's a string type not int
	//fmt.Printf("%s\n",strconv.Atoi(string(s))) // when it is int then it will run but if string it's will show error
	c1, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("strconv.Atoi() is faild %s \n", err) //NOTE : MUST IMPORT LOG
	}
	fmt.Printf("%d\n",c1)

	// convert now with Sscanf
	s9 := "56"
	var c2 int
	_, err =fmt.Sscanf(s9,"%d", &c2)
	if err != nil {
		log.Fatalf("fmt.Sscanf is faild %s \n", err) //NOTE : MUST IMPORT LOG
	}
	fmt.Printf("%d\n",c2)

	// Now Float 
	// In go float32 is our c float and float64 is our double float types
	//int to string
	var f32 float32 = 4.5
	bitSize := 32
	s5 := strconv.FormatFloat(float64(f32), 'E', -1 , bitSize)
	fmt.Printf("%s\n", s5)
	var f64 float64 = 5.5
	bitSize = 64
	s6 := strconv.FormatFloat(float64(f64), 'E', -1 , bitSize)
	fmt.Printf("%s\n", s6)

	var sf64 float64 = 6.5
	s7 := fmt.Sprintf("%f", sf64)
	fmt.Printf("%s\n",s7)
	
	// Strconv is faster then Sprintf

	// string to float using ParshFloat and Sccanf
	s1f64 := "7.5"
	s2f, err := strconv.ParseFloat(s1f64,64)
	if err != nil {
		log.Fatalf("strconv.ParseFloat() is fail %s\n", err)
	}
	fmt.Printf("%f\n", s2f)

	var s3f float64
	_, err = fmt.Sscanf(s1f64,"%f",&s3f)
	if err != nil {
		log.Fatalf("fmt.Sscanf is fail %s\n", err)
	}
	fmt.Printf("%f\n", s3f)

	// Character is 2 types byte and rune 
	// byte is unit8 type and rune is int32 type (byte = 1 byte value and rune = unicode (encode UTF-8 and 1-4 bytes))

	cs := "zonieed"
	for i:=0; i<len(cs); i++{
		c := cs[i]
		fmt.Printf("byte of index %d is '%c' (0x%x)\n", i, c, c)
	}

	sc := "zonieed"
	for i,runeChar := range sc {
		fmt.Printf("Rune of the byte position %d is %#U\n",i, runeChar)
	}

	 var z string
	 z1 := "zonieed is a good boy "
	 z2 := "zonieed is bad boy"
	 fmt.Printf("%s\n",z+z1+z2)

	 var ar1 = [5]byte{1,2,3,4,5}
	ar2 := []byte{1,2}
	fmt.Printf("Length  of array %d %d\n",ar1, ar2)
	fmt.Printf("Length  of array %d %d\n",len(ar1), len(ar2))
	m := make(map[string]int)
	m ["number3"] = 3
	ck := func(k string){
		if n, ok := m[k]; ok{
			fmt.Printf("value of key %s is  %d\n", k, n)
		}else{
			fmt.Printf("key doesn't %s exist\n", k)
		}
	} 
	ck("number3")
	ck("number4")
	delete(m, "number3")
	fmt.Printf("key deleted\n")
	ck("number3")
	
	var p int= 4
	pa := &p
	fmt.Printf("Here is the pointer address %p and the value %d\n", pa, p)

	// channel is used for goroutine to communicate with each other in a thread safe manner 
	ch := make(chan int)

	go func() { ch <- 3}()

	n := <-ch
	fmt.Printf("passed: %d\n", n)

	// zero values
	var zb bool
	fmt.Printf("bool: %v\n",zb)
	var zi int
	fmt.Printf("int: %v\n",zi)
	var zs string
	fmt.Printf("string: %v\n",zs)
	var zf float64
	fmt.Printf("floact64: %v\n",zf)
	var zss []uint32
	fmt.Printf("slice: %v\n",zss)
	var zp *int
	fmt.Printf("pointer: %v\n",zp)
	var zm map[string]int
	fmt.Printf("map: %v\n",zm)
	var zin interface{}
	fmt.Printf("interface: %v\n",zin)
	var zc chan bool
	fmt.Printf("chan: %v\n",zc)
	var za [5]int
	fmt.Printf("array: %v\n",za)
	type struc struct{
		a int
		b string
	}
	var zst struc
	fmt.Printf("struct: %v\n",zst)
	var zff func(bool)
	fmt.Printf("func: %v\n",zff)

}	