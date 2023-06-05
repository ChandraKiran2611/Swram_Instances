package main



import (

	"os"
	"fmt"

)

func main(){
	fileName:=os.Getenv("FILENAME")
    apiPort := os.Getenv("API_PORT")
    gatewayPort := os.Getenv("GATEWAY_PORT")

	//calling the functions 
	InitializeDirectory(fileName)
	ConfigurePorts(apiPort,gatewayPort)

	if err:= StartDaemon();err!=nil{
		fmt.Println("error",err)
		os.Exit(1)
	}
	
}