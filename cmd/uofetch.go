package main

import (
	"fmt"
	"github.com/Runinterface/nonfetch/internal/data"
)


func main() {

        f := data.GetOsInfo()
	fmt.Println(f.OS)
	fmt.Println(f.Kernel)
	fmt.Println(f.CPU)
	fmt.Println(f.Shell)
	fmt.Println(f.Init)
	fmt.Println(f.LA)
	fmt.Println(f.Uptime)
	fmt.Println(f.RAM)
	fmt.Println(f.Logo)

}
