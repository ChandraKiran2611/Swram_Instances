package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func InitializeDirectory() error {
	usr,err := user.Current()
	if err!=nil{
		return err
	}

	//for windows '\' needed to be used. Remaining can be done with "/"
	ipfsPath = fmt.Sprintf("%s\\%s",usr.HomeDir,ReadFileName())

	_,err = os.Stat(ipfsPath)
	if err==nil{
		if err:=os.RemoveAll(ipfsPath);err!=nil{
			return err
		}
		
		if err:=ConfigurePorts();err!=nil{
			return err
		}
	}

	fmt.Println("filePAth:",ipfsPath)
	return ExceCmd(cmd,"init")
}

func ConfigurePorts() error {
	
	if err:= ExceCmd(cmd,config,apiaddr,CheckPorts());err!=nil{ 
		return err
	}
	return ExceCmd(cmd,config,gateaddr,CheckPorts()) 
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