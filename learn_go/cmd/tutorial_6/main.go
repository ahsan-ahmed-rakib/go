package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type owner struct{
	name string
}

func main(){
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons) // 0, 0 shown as default
	
	var myEngine1 gasEngine = gasEngine{10, 20, owner{"Rakib"}}
	fmt.Println(myEngine1.mpg, myEngine1.gallons, myEngine1.ownerInfo.name) 
	myEngine1.gallons = 30
	fmt.Println(myEngine1.mpg, myEngine1.gallons) 
	
	var myEngine2 = struct{
		mpg uint8
		gallons uint8
		}{25, 35}
	fmt.Println(myEngine2.mpg, myEngine2.gallons) 
}
