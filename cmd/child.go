package main

import (
	"gostudy/rpcx"
)

func main() {
	rpcx.SynchronousCall()
	rpcx.AsynchronousCall()
	//rpcx.PrintWrr()
	//rpcx.PrintWrrNgx()
}