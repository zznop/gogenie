package gamegenie

import (
	"errors"
	"fmt"
	"strings"
)

const gameGenieCharacters = "AaBbCcDdEeFfGgHhJjKkLlMmNnPpRrSsTtVvWwXxYyZz0O1I2233445566778899"

// PatchInfo is a struct that represents patch information encoded in Game Genie codes
type PatchInfo struct {
	Address uint32
	Value   uint16
}

// NewPatchInfo decodes the start address and value embedded in a Game Genie code and returns
// an initialized PatchInfo struct
func NewPatchInfo(code []string) (*PatchInfo, error) {
	if err := verifyCode(code); err != nil {
		return nil, err
	}

	address := uint32(0)
	value := uint16(0)

	// Character 0
	n := strings.Index(gameGenieCharacters, code[0]) >> 1
	value |= uint16(n << 3)

	// Character 1
	n = strings.Index(gameGenieCharacters, code[1]) >> 1
	value |= uint16(n >> 2)
	address |= uint32((n & 3) << 14)

	// Character 2
	n = strings.Index(gameGenieCharacters, code[2]) >> 1
	address |= uint32(n << 9)

	// Character 3
	n = strings.Index(gameGenieCharacters, code[3]) >> 1
	address |= uint32(((n & 0x0f) << 20) | ((n >> 4) << 8))

	// Character 4: '-'

	// Character 5
	n = strings.Index(gameGenieCharacters, code[5]) >> 1
	address |= uint32((n >> 1) << 16)
	value |= uint16((n & 1) << 12)

	// Character 6
	n = strings.Index(gameGenieCharacters, code[6]) >> 1
	value |= uint16(((n & 1) << 15) | ((n >> 1) << 8))

	// Character 7
	n = strings.Index(gameGenieCharacters, code[7]) >> 1
	address |= uint32((n & 7) << 5)
	value |= uint16((n >> 3) << 13)

	// Character 8
	n = strings.Index(gameGenieCharacters, code[8]) >> 1
	address |= uint32(n)

	return &PatchInfo{
		Address: address,
		Value:   value,
	}, nil
}

func verifyCode(code []string) error {
	if len(code) != 9 || code[4] != "-" {
		return errors.New("code must be 9 characters in length and contain a hyphen (RFAA-A6VR)")
	}

	for i := 0; i < len(code); i++ {
		if i == 4 {
			continue // skip the hyphen
		}

		if !strings.Contains(gameGenieCharacters, code[i]) {
			return fmt.Errorf("code contains an invalid character: '%v'", code[i])
		}
	}

	return nil
}
