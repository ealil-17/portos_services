package main

import(
	"fmt"
	"log"
	"strings"
	"os/exec"
)
func main(){

    var i string;

	fmt.Printf("Enter the command to exe ");
	fmt.Scan(&i);
	//fmt.Printf("the string is %v",i)

	args := strings.Split(i," ");

	cmd := exec.Command(args[0],args[1:]...);

	out , err := cmd.CombinedOutput();

	if err!=nil{
	   log.Printf("%v",err);
    }
    log.Printf("\n%s" ,out);
}
