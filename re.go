package main

import "fmt"

func Reverse(){
	a:= [9]int{0,1,2,3,4,5,6,7,8}
	fmt.Println("原始数组",a);
	for i:=0;i< len(a)/2;i++  {
		temp:=a[i]
		a[i]=a[len(a)-1-i]
		a[len(a)-1-i]=temp
	}
	fmt.Println("置换数组",a);

}