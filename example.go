// +build ignore

package main

import (
	"log"

	rfkill "github.com/vbatts/go-rfkill"
)

func main() {

	devs, err := rfkill.ListAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, dev := range devs {
		if dev.Type == rfkill.WlanType && dev.IsBlocked() {
			if err := dev.Unblock(); err != nil {
				log.Fatal(err)
			}
		}
	}

}
