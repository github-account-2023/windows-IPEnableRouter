package main

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Tcpip\Parameters`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer func(k registry.Key) {
		err := k.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(k)

	err = k.SetDWordValue("IPEnableRouter", 1)
	if err != nil {
		log.Fatal(err)
	}
}
