package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var g string
	fmt.Println("its the goat", g)

	var u1 uint32      //declare a variable and init with 0
	u1 = 32            //assign its value
	var u2 uint32 = 32 //declare a variable and assign its value at once
	//declare a new variable with defining data type:
	u3 := uint32(32)        //inside the function block this is equal to: var u3 uint32 = 32
	fmt.Println(u1, u2, u3) //32 32 32
	//u3 := 20//err: no new variables on left side of :=
	u3 = 20
	fmt.Println(u1, u2, u3)       //32 32 20
	u3, str4 := 100, "str"        // at least one new var
	fmt.Println(u1, u2, u3, str4) //32 32 100 str

}
