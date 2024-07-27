package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var seatingCapacity = 2
var arriveRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool, 1)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	shop.addBarber("Frank")
	shop.addBarber("Tomas")
	shop.addBarber("Neal")
	shop.addBarber("Alan")
	shop.addBarber("Kelly")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arriveRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Duration(randomMilliseconds) * time.Millisecond):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed
}
