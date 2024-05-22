package data

import ( 
	"os/exec"
	"fmt"
	"strings"
	"strconv"
)

func uptimeInfo()  string { 
        uptimeInfo, err := exec.Command("/bin/bash", "-c", "/usr/bin/cat /proc/uptime").Output()
        if err != nil {
                fmt.Println("(!) Error: ", err)
        }
	strUpArray := strings.Split(string(uptimeInfo), " ")
	getTime := strings.Split(strUpArray[0], ".")
	sec, err := strconv.ParseInt(string(getTime[0]), 10, 32)
	if err != nil {
		panic(err)
	}
	min := sec / 60
	h := min / 60
	rmin := min - (h * 60)
	if rmin > 59 {
		rmin = 0
	}
	d := h / 24
	rh := h - (d * 24)
	rsec := sec - (min * 60)
	
	return strconv.FormatInt(d, 10) + "d:" + strconv.FormatInt(rh, 10) + "h:"  + strconv.FormatInt(rmin, 10) + "m:"  + strconv.FormatInt(rsec, 10) + "." +  getTime[1] + "sec. "

}




