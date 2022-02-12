package utils

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	"Zinx/ziface"
)

/*
	Store all global parameters.
*/

type GlobalObj struct {
	/*
		Server
		@TcpServer:		Server object for current Zinx global
		Host:			IP address listened by host server
		TcpPort:		Port listened by host server
		Name:			Name of server
	*/

	TcpServer		ziface.IServer
	Host			string
	TcpPort			int
	Name			string

	/*
		Zinx
		Version:		version of current zinx framework
		MaxConn:		The maximum number of connection
		MaxPackageSize:	The maximum number of bytes processed by the current server
	*/

	Version			string
	MaxConn			int
	MaxPackageSize	uint32
}

var GlobalObject *GlobalObj

// Load config file from 'conf/zinx.json'
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("./conf/zinx.json")
	if err != nil {
		fmt.Println("Load config file error")
		panic(err)
	}
	//Extract data of the config file into global object.
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

// Initialize global object.
func init() {
	// Default
	GlobalObject = &GlobalObj{
		TcpServer:nil,
		Name:"ZinxServerApp",
		Version:"v0.4",
		TcpPort:8999,
		Host:"0.0.0.0",
		MaxConn:1000,
		MaxPackageSize:4096,
	}

	// Try to load config file from 'conf/zinx.json'
	GlobalObject.Reload()
}
