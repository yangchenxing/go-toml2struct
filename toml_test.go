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
	testTomlContent = `
IntA = 1
IntB = 2
IntC = "3"
IntHex = "0x0fffffff"
IntOct = "0755"
FloatA = 1.0
BoolA = false
BoolB = "true"
DurationA = "1h"
TimeA = "2006-01-02:15:04:05"

[EmbedA]
I = "99"
`
)

func TestLoad(t *testing.T) {
	if err := ioutil.WriteFile("test.toml", []byte(testTomlContent), 0755); err != nil {
		t.Error("save test.toml fail:", err.Error())
		return
	}
	defer os.Remove("test.toml")
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
	if err := Load("test.toml", "", &actual); err != nil {
		t.Error("load test.toml fail:", err.Error())
		return
	}
	if expect != actual {
		t.Errorf("load test.toml fail: expected=%s, actual=%s", jsonify(expect), jsonify(actual))
		return
	}
}

func jsonify(i interface{}) string {
	c, _ := json.Marshal(i)
	return string(c)
}
