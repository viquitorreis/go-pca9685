package main

import (
	"log"
	"time"

	i2c "github.com/googolgl/go-i2c"
	pca9685 "github.com/googolgl/go-pca9685"
)

func main() {
	i2c, err := i2c.New(0x62, "/dev/i2c-1")
	if err != nil {
		log.Printf("Error initializing I2C: %v", err)
		return
	}
	defer i2c.Close()

	pca, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Printf("Error initializing PCA9685: %v", err)
		return
	}

	if err := pca.SetFreq(1000); err != nil {
		log.Println("error setting frequency")
		return
	}

	// Blink LED on channel 0
	for {
		log.Println("blinking bwm...")
		// Full brightness (LED ON)
		err = pca.SetChannel(0, 0, 4095) // Channel 0, ON=0, OFF=4095 (100% duty cycle)
		if err != nil {
			log.Fatal("SetPWM failed:", err)
		}
		time.Sleep(1 * time.Second)

		// Turn LED OFF
		err = pca.SetChannel(0, 0, 0) // OFF=0 (0% duty cycle)
		if err != nil {
			log.Fatal("SetPWM failed:", err)
		}
		time.Sleep(1 * time.Second)
	}
}
