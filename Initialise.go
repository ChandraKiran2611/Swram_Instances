package main

import (
	"os/user"
	"fmt"
	
)

func InitializeDirectory() error {
	usr,err := user.Current()
	if err!=nil{
		return err
	}

	ipfsPath = fmt.Sprintf("%s/%s",usr.HomeDir,ReadFileName())
	return ExceCmd(cmd,"init")
	
}

func ConfigurePorts() error {
	ExceCmd(cmd,config,apiaddr,CheckPorts())
	ExceCmd(cmd,config,gateaddr,CheckPorts())
	return nil
}

func StartDaemon() bool {
	err:= ExceCmd(cmd ,daemon)
	fmt.Println("Daemon Running")
	if err!=nil{
		return false
	}
	return true
}
