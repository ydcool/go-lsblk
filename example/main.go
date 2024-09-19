package main

import (
	"encoding/json"
	"fmt"
	golsblk "github.com/ydcool/go-lsblk"
	"log"
)

func main() {
	devices, err := golsblk.ListBlockDevice()
	if err != nil {
		log.Fatalln(err)
	}
	out, err := json.MarshalIndent(devices, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))
}
