package main



import (

	"os"
	"fmt"

)

func main(){


	//calling the functions 
	if err:=InitializeDirectory();err!=nil{
		fmt.Println("error",err)
	}

	
	ConfigurePorts()

	if err:= StartDaemon();err!=nil{
		fmt.Println("error",err)
		os.Exit(1)
	}
	
}