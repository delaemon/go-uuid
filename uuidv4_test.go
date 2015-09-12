package uuidv4

import (
	"encoding/hex"
	"regexp"
	"strings"
	"testing"
)

const UuidRegex = "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}"

func TestGenerateV4(t *testing.T) {
	uuid, _ := Generate()
	matched, err := regexp.MatchString(UuidRegex, uuid)
	if err != nil {
		t.Errorf("format check error. %s", err)
	}
	if !matched {
		t.Errorf("invalid format. uuid: %s", uuid)
	}

	b, err := hex.DecodeString(strings.Replace(uuid, "-", "", -1))
	if err != nil {
		t.Errorf("%s", err)
	}

	version := b[6] >> 4
	if int(version) != 4 {
		t.Errorf("invalid version. uuid: %s, version: %d", uuid, version)
	}

	if b[8] < VarRFC {
		t.Errorf("invalid variant. uuid: %s, variant: %d", uuid, b[8])
	}
}
