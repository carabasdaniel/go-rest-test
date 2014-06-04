// RestExec project main.go
package main

import (
	"code.google.com/p/gorest"	
	"net/http"		
)

func main() {
	gorest.RegisterService(new(CmdService)) //Register our service
	http.Handle("/", gorest.Handle())	
	http.ListenAndServe(":4050", nil)
}

//Service Definition
type CmdService struct {
	gorest.RestService `root:"/"`
	sayHello       gorest.EndPoint `method:"GET" path:"/Hello/{name:string}" output:"string"`
}

func (serv CmdService) SayHello(name string) string {
	return "Hello "+name
}
