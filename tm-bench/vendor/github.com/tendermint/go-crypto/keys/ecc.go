package keys

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"hash/crc64"
)

// ECC is used for anything that calculates an error-correcting code
type ECC interface {
	// AddECC calculates an error-correcting code for the input
	// returns an output with the code appended
	AddECC([]byte) []byte

	// CheckECC verifies if the ECC is proper on the input and returns
	// the data with the code removed, or an error
	CheckECC([]byte) ([]byte, error)
}

// NoECC is a no-op placeholder, kind of useless... except for tests
type NoECC struct{}

var _ ECC = NoECC{}

func (_ NoECC) AddECC(input []byte) []byte            { return input }
func (_ NoECC) CheckECC(input []byte) ([]byte, error) { return input, nil }

// CRC32 does the ieee crc32 polynomial check
type CRC32 struct {
	Poly  uint32
	table *crc32.Table
}

var _ ECC = &CRC32{}

func NewIEEECRC32() *CRC32 {
	return &CRC32{Poly: crc32.IEEE}
}

func NewCastagnoliCRC32() *CRC32 {
	return &CRC32{Poly: crc32.Castagnoli}
}

func NewKoopmanCRC32() *CRC32 {
	return &CRC32{Poly: crc32.Koopman}
}

func (c *CRC32) AddECC(input []byte) []byte {
	table := c.getTable()

	// get crc and convert to some bytes...
	crc := crc32.Checksum(input, table)
	check := make([]byte, crc32.Size)
	binary.BigEndian.PutUint32(check, crc)

	// append it to the input
	output := append(input, check...)
	return output
}

func (c *CRC32) CheckECC(input []byte) ([]byte, error) {
	table := c.getTable()

	if len(input) <= crc32.Size {
		return nil, errors.New("input too short, no checksum present")
	}
	cut := len(input) - crc32.Size
	data, check := input[:cut], input[cut:]
	crc := binary.BigEndian.Uint32(check)
	calc := crc32.Checksum(data, table)
	if crc != calc {
		return nil, errors.New("Checksum does not match")
	}
	return data, nil
}

func (c *CRC32) getTable() *crc32.Table {
	if c.table == nil {
		if c.Poly == 0 {
			c.Poly = crc32.IEEE
		}
		c.table = crc32.MakeTable(c.Poly)
	}
	return c.table
}

// CRC64 does the ieee crc64 polynomial check
type CRC64 struct {
	Poly  uint64
	table *crc64.Table
}

var _ ECC = &CRC64{}

func NewISOCRC64() *CRC64 {
	return &CRC64{Poly: crc64.ISO}
}

func NewECMACRC64() *CRC64 {
	return &CRC64{Poly: crc64.ECMA}
}

func (c *CRC64) AddECC(input []byte) []byte {
	table := c.getTable()

	// get crc and convert to some bytes...
	crc := crc64.Checksum(input, table)
	check := make([]byte, crc64.Size)
	binary.BigEndian.PutUint64(check, crc)

	// append it to the input
	output := append(input, check...)
	return output
}

func (c *CRC64) CheckECC(input []byte) ([]byte, error) {
	table := c.getTable()

	if len(input) <= crc64.Size {
		return nil, errors.New("input too short, no checksum present")
	}
	cut := len(input) - crc64.Size
	data, check := input[:cut], input[cut:]
	crc := binary.BigEndian.Uint64(check)
	calc := crc64.Checksum(data, table)
	if crc != calc {
		return nil, errors.New("Checksum does not match")
	}
	return data, nil
}

func (c *CRC64) getTable() *crc64.Table {
	if c.table == nil {
		if c.Poly == 0 {
			c.Poly = crc64.ISO
		}
		c.table = crc64.MakeTable(c.Poly)
	}
	return c.table
}
