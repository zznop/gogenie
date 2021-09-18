package gamegenie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

const checksumStart = 0x18e
const romStart = 0x200

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

// FixChecksum is a method for computing the ROM checksum and writing back to the Checksum field in the ROM header
func (r *ROM) FixChecksum() error {
	sum := uint16(0)
	for i := romStart; i < len(r.raw); i += 2 {
		sum += binary.BigEndian.Uint16(r.raw[i : i+2])
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, &sum); err != nil {
		return err
	}

	sumBytes := buf.Bytes()
	copy(r.raw[checksumStart:checksumStart+len(sumBytes)], sumBytes)
	return nil
}

// Save is a method for writing a ROM to disk
func (r ROM) Save(filePath string) error {
	return ioutil.WriteFile(filePath, r.raw, 0644)
}
