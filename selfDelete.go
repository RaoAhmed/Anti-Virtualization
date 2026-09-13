package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func SelfDelete() {
	// 1. Get the path of the currently running executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error resolving executable path: %v\n", err)
		return
	}

	fmt.Printf("Preparing to delete: %s\n", exePath)

	// 2. Construct a command that waits for this process to exit, then removes the file.
	// /C executes the command string and terminates cmd.exe.

	cmd := exec.Command("cmd.exe", "/C", "del", "/f", "/q", exePath) // working command

	// 3. Detach the process completely and hide the window
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP | 0x08000000, // 0x08000000 = CREATE_NO_WINDOW
		HideWindow:    true,
	}

	// 4. Start the cleanup command asynchronously
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to spawn deletion process: %v\n", err)
		return
	}

	fmt.Println("\nCleanup process spawned. Exiting executable...")
	os.Exit(0)
}
