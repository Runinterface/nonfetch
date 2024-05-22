package data

import ( 
	"os/exec"
	"fmt"
	"strings"
)

func shellName() string {
	shellInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/echo $0").Output()
        if err != nil {
                 fmt.Println("(!) Error: ", err)
         }
        return strings.Trim(string(shellInfo), "\n")
}




