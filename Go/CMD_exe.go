package main

import(
	"fmt"
	"log"
	"strings"
	"os/exec"
)
func main(){

	//getting input from the front-end
    var i string;
	fmt.Printf("Enter the command to exe ");
	fmt.Scan(&i);
	
	//calling the defined-function
	exe(i);
}

//  func that execute the terminal commands

func exe(var i string){


	args := strings.Split(i," ");

	cmd := exec.Command(args[0],args[1:]...);

	out , err := cmd.CombinedOutput();

 // checking for errors
 
	if err!=nil{
	   log.Printf("%v",err);
    }

// Displaying the logs

    log.Printf("\n%s" ,out);


}