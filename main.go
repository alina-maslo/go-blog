package main

import (
	"os"
	"fmt"
	//"flag"
	//"strings"

	//"net/http"

	"github.com/labstack/echo"
)

func MsgPrintln(msg string) {
	fmt.Println("\033[1;92m" + msg + "\033[0m")
}

func main() {
	MsgPrintln("test test test")
	fmt.Println("%#v", os.Args)

	//server := fcgi.S

	e := echo.New()
	e.Start()
}
