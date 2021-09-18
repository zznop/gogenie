package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zznop/gogenie/pkg/gamegenie"
)

func displayUsage() {
	fmt.Printf(`Usage: %v <code> <in> <outt>

Required:
  code    SEGA Genesis Game Genie alphanumeric code
  in      File path to input ROM
  out     File path to write out patched ROM

example:
  %v RFAA-A6VR StreetsOfRage.bin StreetsOfRagePatched.bin
`, os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 4 {
		displayUsage()
	}

	// Get patch info from the Game Genie code
	code := strings.Split(os.Args[1], "")
	pi, err := gamegenie.NewPatchInfo(code)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Load the ROM from disk
	inputPath := os.Args[2]
	fmt.Printf("* Loading the ROM (%v)...\n", inputPath)
	rom, err := gamegenie.ReadROM(inputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Apply the patch
	fmt.Printf(`* Applying the patch...
  - Address : 0x%08x
  - Value   : %04x
`, pi.Address, pi.Value)
	err = rom.ApplyPatch(*pi)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Fixup the checksum
	fmt.Println("* Re-computing the checksum...")
	err = rom.FixChecksum()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Export the patched ROM
	outputPath := os.Args[3]
	fmt.Printf("* Saving the patched ROM (%v)...\n", outputPath)
	err = rom.Save(outputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
