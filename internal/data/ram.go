package data

import (
	"syscall"
	"strconv"
)


func ramInfo() string {
	var result string
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		panic(err)
	}
	
	usedMem := (in.Totalram - in.Freeram - in.Totalswap - in.Sharedram - in.Bufferram)
	result = strconv.FormatUint((usedMem / 1024) / 1024,10) +" / "+ strconv.FormatUint((in.Totalram / 1024) / 1024, 10) + " MB"
	return result
}

