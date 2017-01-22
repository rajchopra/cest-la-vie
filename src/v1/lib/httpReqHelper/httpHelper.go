package main

 import (
 	"fmt"
 	"io/ioutil"
 	"net/http"
 	"os"
 )

 func main() {

 	resp, err := http.Get("http://ifconfig.co")

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	defer resp.Body.Close()

 	htmlData, err := ioutil.ReadAll(resp.Body)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	fmt.Println("My IP is : ", string(htmlData))


 }
