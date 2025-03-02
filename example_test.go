// SPDX-FileCopyrightText: 2022-2025 Rafael G. Martins <rafael@rafaelmartins.eng.br>
// SPDX-License-Identifier: BSD-3-Clause

package octokeyz_test

import (
	"fmt"
	"log"
	"time"

	"rafaelmartins.com/p/octokeyz"
)

func ExampleDevice() {
	dev, err := octokeyz.GetDevice("")
	if err != nil {
		log.Fatal(err)
	}

	if err := dev.Open(); err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	for i := 0; i < 3; i++ {
		dev.Led(octokeyz.LedFlash)
		time.Sleep(100 * time.Millisecond)
	}

	dev.AddHandler(octokeyz.BUTTON_1, func(b *octokeyz.Button) error {
		fmt.Println("pressed")
		duration := b.WaitForRelease()
		fmt.Printf("released. pressed for %s\n", duration)
		return nil
	})

	if err := dev.Listen(nil); err != nil {
		log.Fatal(err)
	}
}
