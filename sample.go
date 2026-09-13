package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/yusufpapurcu/wmi"
)

// 1. The list of target words to search for
var items = []string{"innotek GmbH", "VirtualBox", "VBOX", "Oracle Corporation",
	"VMware", "VMware, Inc.", "VWM"}

func sample() {

	// System Manufacturer & System Model
	var compSys []Win32_ComputerSystem
	err := wmi.Query("SELECT Manufacturer, Model FROM Win32_ComputerSystem", &compSys)
	if err != nil || len(compSys) == 0 {
		log.Printf("Failed to query Win32_ComputerSystem: %v", err)
		compSys = append(compSys, Win32_ComputerSystem{Manufacturer: "Unknown", Model: "Unknown"})
	}

	for _, meta := range compSys {
		findMatches(meta.Manufacturer)
		findMatches(meta.Model)
	}

	// BIOS Version
	var bios []Win32_BIOS
	err = wmi.Query("SELECT SMBIOSBIOSVersion, Manufacturer, Version FROM Win32_BIOS", &bios)
	if err != nil || len(bios) == 0 {
		log.Printf("Failed to query Win32_BIOS: %v", err)
		bios = append(bios, Win32_BIOS{SMBIOSBIOSVersion: "Unknown", Manufacturer: "Unknown", Version: "Unknown"})
	}

	for _, meta := range bios {
		findMatches(meta.SMBIOSBIOSVersion)
		findMatches(meta.Manufacturer)
		findMatches(meta.Version)
	}

	// BaseBoard Manufacturer & BaseBoard Product
	var baseBoard []Win32_BaseBoard
	err = wmi.Query("SELECT Manufacturer, Product FROM Win32_BaseBoard", &baseBoard)
	if err != nil || len(baseBoard) == 0 {
		log.Printf("Failed to query Win32_BaseBoard: %v", err)
		baseBoard = append(baseBoard, Win32_BaseBoard{Manufacturer: "Unknown", Product: "Unknown"})
	}

	for _, meta := range baseBoard {
		findMatches(meta.Manufacturer)
		findMatches(meta.Product)
	}
}

// findMatches checks if each word in the list exists inside the given sentence.
func findMatches(sentence string) {
	// Normalize the sentence to lowercase for case-insensitive matching
	lowerSentence := strings.ToLower(sentence)

	fmt.Printf("Searching in: %q\n", sentence)

	for _, word := range items {
		lowerWord := strings.ToLower(word)

		// strings.Contains checks for substring matches
		if strings.Contains(lowerSentence, lowerWord) {
			fmt.Printf("Found match: %s\n\n", word)
			SelfDelete()
		}
	}
}
