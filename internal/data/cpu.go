package data

import (
	"os/exec"
	"strings"
	"fmt"
)

func cpuName() string {

	cpuInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/cat /proc/cpuinfo").Output()
        if err != nil {
                fmt.Println("(!) Error: ", err)
        }
	
	var result string
        strCPUArray := strings.Split(string(cpuInfo), "\n")
        for i := 0; i < len(strCPUArray); i++ {
                if strings.Contains(strCPUArray[i], "model name") {
                        a := strings.Split(strCPUArray[i], ":")
                        result = strings.Trim(string(a[1]), " ")
                        break;
                }
        }

	return result

}
