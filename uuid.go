package uuid
import (
	"fmt"
	"crypto/rand"
)

const (
	// 00001111 Clears all bits of version byte with AND
	ClearVer  =  15
	// 00111111  Clears all relevant bits of variant byte with AND
	ClearVar  =  63
	// 10000000	 The RFC 4122 variant (this variant)
	VarRFC    = 128
	// 01000000
	Version4  =  64
)

func format(b []byte) string {
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4],
		b[4:6],
		b[6:8],
		b[8:10],
		b[10:16])
	return uuid
}

// Generate a Version 4 UUID.
// These are derived solely from random numbers.
func GenerateV4() (string, error) {
	c := 16
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	b[8] = b[8] & ClearVar | VarRFC
	b[6] = b[6] & ClearVer | Version4
	uuid := format(b)
	return uuid, nil
}

