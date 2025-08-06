//2 types to declare variables

// var name type
package main

import "fmt"


func main(){
var name = "Hung" 
nameLover := "Ngan"
var height = 1.6
age := 25
ageLover := 22
iAmMale :=true
loverIsMale := false
var float32Num float32 = 4.1234567891011
var float64Num float64 = 4.12345678901234567815161718

// I love this
  var (
     a int = 12
     b int = 1
     c string = " hello"
   )


var video byte = 23 //ASCII , raw data, decode, encode	(uint8)
var text rune = '🤣' //not ASCII, unicode, utf-8 (int32)


fmt.Println(nameLover)
fmt.Println(age)
fmt.Println(height)
fmt.Println(ageLover)
fmt.Println(float32Num)
fmt.Println(float64Num)
fmt.Println(name)
fmt.Println(iAmMale)
fmt.Println(loverIsMale)
fmt.Println(video)
fmt.Println(text)
fmt.Print(a, b, c)
}
