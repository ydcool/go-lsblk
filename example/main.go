package main

import (
	"encoding/json"
	"fmt"
	go_lsblk "github.com/ydcool/go-lsblk"
	"log"
)

func main() {
	devices, err := go_lsblk.ListBlockDevice()
	if err != nil {
		log.Fatalln(err)
	}
	out, err := json.MarshalIndent(devices, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))
}
