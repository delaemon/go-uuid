package uuidv4

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestGenerateV4(t *testing.T) {
	uuid, _ := Generate()
	Validate(uuid)

	uuid, _ = Generate()
	uuid = strings.Replace(uuid, "-", "=", 1)
	Validate(uuid)

	uuid, _ = Generate()
	b, _ := hex.DecodeString(strings.Replace(uuid, "-", "", -1))
	b[6] = 0
	uuid = format(b)
	Validate(uuid)

	uuid, _ = Generate()
	b, _ = hex.DecodeString(strings.Replace(uuid, "-", "", -1))
	b[8] = 0
	uuid = format(b)
	Validate(uuid)
}
