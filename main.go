package main

import (
	"delivery-much-challenge/conf"
	"fmt"
)

func main() {
	environmentConfiguration := conf.Env()
	databaseConnect(environmentConfiguration)
	createServer(environmentConfiguration).Listen(fmt.Sprintf(":%s", environmentConfiguration.Port))
}
