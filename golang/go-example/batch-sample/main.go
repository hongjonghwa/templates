package main
import (
    "fmt"
    "os"
)
func main() {
    argsWithProg := os.Args
    
    if len(argsWithProg) == 1{
        fmt.Println("usage: git <command> [<args>]")
        // os.Exit(1)
        return    
    }
	
	argsWithoutProg := os.Args[1:]
	len1 := len(argsWithProg)
	len2 := len(argsWithoutProg)
	fmt.Println(len1)
	fmt.Println(len2)

    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
	
}