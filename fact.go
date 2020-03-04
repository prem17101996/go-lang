package main
import "fmt"
  


func factorial(a int)  {   
    if(a < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{ 
	var factVal int = 1
	var i int       
        for i:=1; i<=a; i++ {
            factVal =factVal* (i) 
        }
         
    }    
    return factVal 
}
 
func main(){  
	var n int  
    fmt.Print("Enter a positive integer between 0 - 50 : ")
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",factorial(n))
}