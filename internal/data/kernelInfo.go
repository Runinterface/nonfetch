package data

import (
	"fmt"
	"strings"
	"os/exec"
)



func kernelVerion() string {

	kernelInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/uname -sr").Output()
        if err != nil {
                fmt.Println("(!) Error: ", err)
        }
        kernelTrim := strings.Trim(string(kernelInfo), "\n")
        return kernelTrim

}
