package main

import (
	"fmt"
  "math"
)
func random(seed,a,b,m,n,limit int) []int{
  var ret []int
  var x int 
  x=seed 
  for i:=0;i<n;i++{
    var x2 int=(a*x + b)
    x2=x2%m
    x=x2
    if(x2<0){
      x2=-x2  //Eliminating -ve numbers.
    }
    var temp int=x2%limit //setting the upper limit of the elements

    ret=append(ret,temp)
    
  }
  return ret
}
func main() {
   var seed int =13
   var size int =50
   var a int =	25214903917
   var b int =11
   var mod int =int(math.Pow(2,48))
   var limit int=15
	 x:=random(seed,a,b,mod,size,limit)
   fmt.Println(x)
  
}
