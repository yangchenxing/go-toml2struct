package toml2struct

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type TestEmbedTypeB struct {
	I int
}

type TestTypeA struct {
	IntA      int
	IntB      uint32
	IntC      int64
	IntHex    uint32
	IntOct    uint16
	FloatA    float32
	BoolA     bool
	BoolB     bool
	DurationA time.Duration
	TimeA     time.Time
	EmbedA    TestEmbedTypeB
}

const (
	test1TomlContent = `
include = [
    "test2.toml",
]
IntA = 1
IntB = 2
IntC = "3"
IntHex = "0x0fffffff"
IntOct = "0755"
FloatA = 1.0
BoolA = false
BoolB = "true"
`
	test2TomlContent = `
DurationA = "1h"
TimeA = "2006-01-02:15:04:05"
[EmbedA]
I = "99"
`
)

func TestLoadInclude(t *testing.T) {
	if err := ioutil.WriteFile("test1.toml", []byte(test1TomlContent), 0755); err != nil {
		t.Error("save test1.toml fail:", err.Error())
		return
	}
	defer os.Remove("test1.toml")
	if err := ioutil.WriteFile("test2.toml", []byte(test2TomlContent), 0755); err != nil {
		t.Error("save test2.toml fail:", err.Error())
		return
	}
	defer os.Remove("test2.toml")
	expectTime, _ := time.Parse("2006-01-02:15:04:05", "2006-01-02:15:04:05")
	expect := TestTypeA{
		IntA:      1,
		IntB:      2,
		IntC:      3,
		IntHex:    0x0fffffff,
		IntOct:    0755,
		FloatA:    1.0,
		BoolA:     false,
		BoolB:     true,
		DurationA: time.Hour,
		TimeA:     expectTime,
		EmbedA: TestEmbedTypeB{
			I: 99,
		},
	}
	var actual TestTypeA
	if err := Load("test1.toml", "include", &actual); err != nil {
		t.Error("load test.toml fail:", err.Error())
		return
	}
	if expect != actual {
		t.Errorf("load test.toml fail: expected=%s, actual=%s", jsonify(expect), jsonify(actual))
		return
	}
}

func TestLoadNonInclude(t *testing.T) {
	if err := ioutil.WriteFile("test1.toml", []byte(test1TomlContent), 0755); err != nil {
		t.Error("save test1.toml fail:", err.Error())
		return
	}
	defer os.Remove("test1.toml")
	if err := ioutil.WriteFile("test2.toml", []byte(test2TomlContent), 0755); err != nil {
		t.Error("save test2.toml fail:", err.Error())
		return
	}
	defer os.Remove("test2.toml")
	expect := TestTypeA{
		IntA:   1,
		IntB:   2,
		IntC:   3,
		IntHex: 0x0fffffff,
		IntOct: 0755,
		FloatA: 1.0,
		BoolA:  false,
		BoolB:  true,
	}
	var actual TestTypeA
	if err := Load("test1.toml", "", &actual); err != nil {
		t.Error("load test.toml fail:", err.Error())
		return
	}
	if expect != actual {
		t.Errorf("load test.toml fail: expected=%s, actual=%s", jsonify(expect), jsonify(actual))
		return
	}
}

func TestInvalidFile(t *testing.T) {
	if _, err := loadMap(".", ""); err == nil {
		t.Error("load invalid file success.")
		return
	}
}

func TestIncludeBadFile(t *testing.T) {
	if err := ioutil.WriteFile("test.toml", []byte("include=[\"\"]"), 0755); err != nil {
		t.Error("save test.toml fail:", err.Error())
		return
	}
	defer os.Remove("test.toml")
	var m map[string]interface{}
	if err := Load("test.toml", "include", &m); err == nil {
		t.Error("load test.toml success")
		return
	}
}

func jsonify(i interface{}) string {
	c, _ := json.Marshal(i)
	return string(c)
}
