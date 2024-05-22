package data

import (
        "fmt"
        "runtime"
)


type osInfo struct {
        OS       string
	Logo	 string
        Kernel   string
        CPU      string
        RAM      string
        Shell    string
        Init     string
        LA       string
        Uptime   string
}

func GetOsInfo() osInfo {

	infoOS := osInfo{}

	rOS  := runtime.GOOS
	if rOS == "linux" {

		infoOS.OS = osName()
		infoOS.Logo = osLogo()
	        infoOS.Kernel = kernelVerion()
		infoOS.CPU = cpuName()
		infoOS.RAM = ramInfo()
		infoOS.Shell = shellName()
		infoOS.Init = initSystem()
		infoOS.LA = laInfo()
		infoOS.Uptime = uptimeInfo()

	} else if rOS == "darwin" {
		fmt.Println("Stay Tuned: Soon!")

	} else {
		fmt.Println("Error: Unsupported Operation System")
	}

	return infoOS

}
