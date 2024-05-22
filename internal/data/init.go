package data

import ( 
	"os/exec"
	"fmt"
	"strings"
)

func initSystem() string {
	
	initInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/readlink /sbin/init").Output()
        if err != nil  {
                fmt.Println("(!) Error: ", err)
        }
        return strings.Trim(string(initInfo), "\n")

}




