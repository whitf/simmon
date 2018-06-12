//+build windows
package main

import (
	"log"
)

type MemoryThreshold struct {
	memoryFreePercent			float64
	swapFreePercent				float64
}

var mt MemoryThreshold

func memfree(t int, options []ServiceOption) int {

	log.Println("c NetConfig")
	log.Println(c)

	log.Println("t int")
	log.Println(t)

	log.Println("memfree for windows net yet implemented")

	return 1
}

func swapfree(t int, options []ServiceOption) int {

	log.Println("c NetConfig")
	log.Println(c)

	log.Println("t int")
	log.Println(t)

	log.Println("swapfree for windows net yet implemented")

	return 1
}