package main

import (
	"fmt"
	"time"
)

func main() {
	var digits = [11][5]string{
		{
			"|||",
			"| |",
			"| |",
			"| |",
			"|||",
		},
		{

			"|| ",
			" | ",
			" | ",
			" | ",
			"|||",
		},
		{

			"|||",
			"  |",
			"|||",
			"|  ",
			"|||",
		},
		{

			"|||",
			"  |",
			"|||",
			"  |",
			"|||",
		},
		{

			"| |",
			"| |",
			"|||",
			"  |",
			"  |",
		},
		{

			"|||",
			"|  ",
			"|||",
			"  |",
			"|||",
		},
		{

			"|||",
			"|  ",
			"|||",
			"| |",
			"|||",
		},
		{

			"|||",
			"  |",
			"  |",
			"  |",
			"  |",
		},
		{

			"|||",
			"| |",
			"|||",
			"| |",
			"|||",
		},
		{

			"|||",
			"| |",
			"|||",
			"  |",
			"|||",
		},
		{
			"   ",
			" : ",
			" : ",
			"   ",
			"   ",
		},
	}
	for true {
		tim := time.Now()
		var clock [8][5]string
		clock[0] = digits[tim.Hour()/10]
		clock[1] = digits[tim.Hour()%10]
		clock[2] = digits[10]
		clock[3] = digits[tim.Minute()/10]
		clock[4] = digits[tim.Minute()%10]
		clock[5] = digits[10]
		clock[6] = digits[tim.Second()/10]
		clock[7] = digits[tim.Second()%10]
		// fmt.Print(clock)
		k := 0
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if k < 5 {
					fmt.Print(clock[j][k], " ")
				}
			}
			fmt.Println()
			k++
		}
		time.Sleep(time.Second)
		print("\033[H\033[2J")
	}
}
