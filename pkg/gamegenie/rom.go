package gamegenie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

// ROM is a structure that represents a loaded SEGA Genesis ROM
type ROM struct {
	raw []byte
}

// ReadROM reads a ROM from disk
func ReadROM(filePath string) (*ROM, error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &ROM{
		raw: raw,
	}, nil
}

// ApplyPatch is a method for applying a patch to the ROM
func (r *ROM) ApplyPatch(pi PatchInfo) error {
	if pi.Address > uint32(len(r.raw)-binary.Size(pi.Value)) {
		return fmt.Errorf("Patch address is out of range: %x", pi.Address)
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, &pi.Value); err != nil {
		return err
	}

	valueBytes := buf.Bytes()
	copy(r.raw[pi.Address:pi.Address+uint32(len(valueBytes))], valueBytes)
	return nil
}

// Save is a method for writing a ROM to disk
func (r ROM) Save(filePath string) error {
	return ioutil.WriteFile(filePath, r.raw, 0644)
}
