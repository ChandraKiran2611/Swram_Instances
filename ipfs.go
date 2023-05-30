package main

import (
	"fmt"
	"os/exec"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"os"
	
)

var (
	port string
	fileName string
	ipfsPath string 
)

const (
	muladdr  ="/ip4/127.0.0.1/tcp/" 
	envPath  = "IPFS_PATH="
	cmd		 = "ipfs"
	apiaddr  = "Addresses.API"
	gateaddr = "Addresses.Gateway"
	config   = "config"
	daemon   = "daemon"
)


func ExceCmd(cmd string , args ...string) error {
	command :=	exec.Command(cmd,args...)
	command.Env = append(command.Env,envPath+ipfsPath)

	return command.Run()
	
}

func Read(port string) string {
	 fmt.Scan(&port)

	 return port
}

func ReadPort() string {
	fmt.Println("Enter the API/Gatway Port : ")

	return Read(port)
}

func ReadFileName() string {
	fmt.Println("Enter the fileName  : ")
	
	return Read(fileName)
}

func CheckPorts() string {
	maAddr,_ := ma.NewMultiaddr(muladdr+ReadPort())
	_,err:= manet.Listen(maAddr)
	if err!=nil{
		os.RemoveAll(ipfsPath)
		return ""
	}

	addr:= fmt.Sprintf("%s",maAddr)

	return addr
}
