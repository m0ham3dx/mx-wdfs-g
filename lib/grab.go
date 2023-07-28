/*
LinkGrabber and Writer
*/

package lib

import (
	"fmt"
	i "io"
	o "os"
	e "os/exec"
)

// Script Grabber Function
func GrabScript() {
	cmd := e.Command("curl", "https://snips.sh/f/uEf7aLoioN")
	cmd.Stdout = o.Stdout

	err := cmd.Run()
	if err != nil {
		Err("===ERROR===")
		fmt.Println("❌❌❌", "curl download failure", err)
		o.Exit(1)
	}
}

// Script Write Function
func GrabWrite() {

	cmd := e.Command("curl", "https://snips.sh/f/uEf7aLoioN")

	// Create file
	f, err := o.Create("🍌.fish")
	if err != nil {
		Err("===ERROR===")
		fmt.Println("❌❌❌", "file creation failure", err)
		panic(err)
	}
	defer f.Close()

	// Set command stdout to both file and os.Stdout
	ParaText2("\n" + "Downlading and Writing..." + "\n")
	cmd.Stdout = i.MultiWriter(f, o.Stdout)

	err = cmd.Run()
	if err != nil {
		Err("===ERROR===")
		fmt.Println("❌❌❌", "Downloading failure", err)
		panic(err)
	}

}

// File Display Function

func FileDisplay() {
	ParaText2("\n" + "ListFile" + "\n")

	cmd := e.Command("ls", "-al", "🍌.fish")
	cmd.Stdout = o.Stdout
	err := cmd.Run()
	if err != nil {
		Err("===ERROR===")
		fmt.Println("❌❌❌", "file not available...", err)
		o.Exit(1)
	}

	cmd2 := e.Command("file", "🍌.fish")
	cmd2.Stdout = o.Stdout
	err2 := cmd2.Run()
	if err2 != nil {
		Err("===ERROR===")
		fmt.Println("❌❌❌", "file not available", err)
		o.Exit(1)
	}

	cmd3 := e.Command("realpath", "🍌.fish")
	cmd3.Stdout = o.Stdout
	err3 := cmd3.Run()
	if err3 != nil {
		Err("===ERROR===")
		fmt.Println("❌❌❌", "file not available", err)
		o.Exit(1)
	}

	ParaText2("😽" + "\n"  + "1. rename file" + "\n" + "2. cp <file> ~/.config/fish/functions" + "\n" + "3. run wdv")

}
