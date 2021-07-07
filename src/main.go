package main

import (
	nvme "./nvme"
	//drivedb "./drivedb"
	"fmt"
	"os"
)

func main() {
	d := nvme.NewNVMeDevice("/dev/nvme0")
	err := d.Open()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer d.Close()

	if err := d.PrintSMART(os.Stdout); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
