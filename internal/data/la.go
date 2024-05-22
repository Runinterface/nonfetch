package data

import ( 
	"os/exec"
	"fmt"
	"strings"
)

func laInfo() string {

	laInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/cat /proc/loadavg").Output()
        if err != nil {
                fmt.Println("(!) Error: ", err)
        }

        strLaArray := strings.Split(string(laInfo), " ")
        var laResult string
        for i:= 0; i < 3; i++ {
                laResult += strLaArray[i]+ " "
        }
        return laResult

}




