package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func InitializeDirectory(filename string) error {
	usr,err := user.Current()
	if err!=nil{
		return err
	}

	fmt.Println("file:",filename)
	ipfsPath = fmt.Sprintf("%s/%s",usr.HomeDir,filename)

	_,err = os.Stat(ipfsPath)
	if err==nil{
		StartDaemon()
		os.Exit(0)
	}

	fmt.Println("filePAth:",ipfsPath)
	return ExceCmd(cmd,"init")
}

func ConfigurePorts(apiport,gatewayport string) error {
	fmt.Println("api:",apiport)
	fmt.Println("gate:",gatewayport)
	if err:= ExceCmd(cmd,config,apiaddr,CheckPorts(apiport));err!=nil{ 
		return err
	}

	return ExceCmd(cmd,config,gateaddr,CheckPorts(gatewayport)) 
}

func StartDaemon() error {
	run:= exec.Command(cmd,daemon)
	run.Env = append(run.Env,envPath+ipfsPath)
	run.Stderr = os.Stderr
	run.Stdout = os.Stdout

	if err:= run.Start();err!=nil{
		fmt.Println("deamon is not running")
		return err
	}
	fmt.Println("Daemon Running")
	return nil
}