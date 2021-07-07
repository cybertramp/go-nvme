package main

import (
	nvme "./nvme"
	drivedb "./drivedb"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	d := nvme.NewNVMeDevice("/dev/nvme0")

	defer d.Close()

	db, err := drivedb.OpenDriveDb("drivedb.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := d.PrintSMART(&db, os.Stdout); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
