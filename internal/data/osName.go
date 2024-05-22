package data

import (
        "fmt"
        "os/exec"
        "strings"
)


//=== Function wich return os name ===

func osName() string {

	getInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/cat /etc/os-release").Output()
	if err != nil {
		fmt.Println("(!) Error: ", err)
	}
	strOSArray := strings.Split(string(getInfo), "\n")
 	var result string
	for i := 0; i < len(strOSArray); i++ {
		if strings.Contains(strOSArray[i], "PRETTY_NAME") {
			a := strings.Split(strOSArray[i], "\"")
			result =  string(a[1])
			break;
		}
	}
	return result
}


