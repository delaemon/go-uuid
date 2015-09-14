//Package to generate a UUID version 4 according to the RFC 4122 standard.
package uuidv4

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

const (
	// 00001111 Clears all bits of version byte with AND
	ClearVer = 15
	// 00111111  Clears all relevant bits of variant byte with AND
	ClearVar = 63
	// 10000000	 The RFC 4122 variant (this variant)
	VarRFC = 128
	// 01000000
	Version4 = 64

	UuidRegex = "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}"
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
func Generate() (string, error) {
	c := 16
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	b[8] = b[8]&ClearVar | VarRFC
	b[6] = b[6]&ClearVer | Version4
	uuid := format(b)
	return uuid, nil
}

func Validate(uuid string) bool {
	matched, err := regexp.MatchString(UuidRegex, uuid)
	if err != nil {
		return false
	}

	if !matched {
		return false
	}

	b, err := hex.DecodeString(strings.Replace(uuid, "-", "", -1))
	if err != nil {
		return false
	}

	version := b[6] >> 4
	if int(version) != 4 {
		return false
	}

	if b[8] < VarRFC {
		return false
	}
	return true
}
