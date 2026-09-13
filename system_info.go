package main

import (
	"fmt"
	"log"

	// go get github.com/yusufpapurcu/wmi
	"github.com/yusufpapurcu/wmi"
)

type Win32_ComputerSystem struct {
	Manufacturer string
	Model        string
}

type Win32_BIOS struct {
	SMBIOSBIOSVersion string
	Manufacturer      string
	Version           string
}

type Win32_BaseBoard struct {
	Manufacturer string
	Product      string
}

func sysInfo() {

	// 1. System Manufacturer & System Model
	var compSys []Win32_ComputerSystem
	err := wmi.Query("SELECT Manufacturer, Model FROM Win32_ComputerSystem", &compSys)
	if err != nil || len(compSys) == 0 {
		log.Printf("Failed to query Win32_ComputerSystem: %v", err)
		compSys = append(compSys, Win32_ComputerSystem{Manufacturer: "Unknown", Model: "Unknown"})
	}

	// 2. BIOS Version
	var bios []Win32_BIOS
	err = wmi.Query("SELECT SMBIOSBIOSVersion, Manufacturer, Version FROM Win32_BIOS", &bios)
	if err != nil || len(bios) == 0 {
		log.Printf("Failed to query Win32_BIOS: %v", err)
		bios = append(bios, Win32_BIOS{SMBIOSBIOSVersion: "Unknown", Manufacturer: "Unknown", Version: "Unknown"})
	}

	// 3. BaseBoard Manufacturer & BaseBoard Product
	var baseBoard []Win32_BaseBoard
	err = wmi.Query("SELECT Manufacturer, Product FROM Win32_BaseBoard", &baseBoard)
	if err != nil || len(baseBoard) == 0 {
		log.Printf("Failed to query Win32_BaseBoard: %v", err)
		baseBoard = append(baseBoard, Win32_BaseBoard{Manufacturer: "Unknown", Product: "Unknown"})
	}

	// Print Results
	fmt.Println("================ Hardware Details ================")
	fmt.Printf("System Manufacturer:     %s\n", compSys[0].Manufacturer)
	fmt.Printf("System Model:            %s\n", compSys[0].Model)
	fmt.Printf("Bios Version:            %s\n", bios[0].SMBIOSBIOSVersion)
	fmt.Printf("BIOS Manufacturer:       %s\n", bios[0].Manufacturer)
	fmt.Printf("BIOS Version:    %s\n", bios[0].Version)
	fmt.Printf("BaseBoard Manufacturer:  %s\n", baseBoard[0].Manufacturer)
	fmt.Printf("BaseBoard Product:       %s\n", baseBoard[0].Product)
	fmt.Println("==================================================")
	fmt.Println()
}
