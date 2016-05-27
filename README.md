# go-toml2struct

[![Go Report Card](https://goreportcard.com/badge/github.com/yangchenxing/go-toml2struct)](https://goreportcard.com/report/github.com/yangchenxing/go-toml2struct)
[![Build Status](https://travis-ci.org/yangchenxing/go-toml2struct.svg?branch=master)](https://travis-ci.org/yangchenxing/go-toml2struct)
[![GoDoc](http://godoc.org/github.com/yangchenxing/go-toml2struct?status.svg)](http://godoc.org/github.com/yangchenxing/go-toml2struct)
[![Coverage Status](https://coveralls.io/repos/github/yangchenxing/go-toml2struct/badge.svg?branch=master)](https://coveralls.io/github/yangchenxing/go-toml2struct?branch=master)

go-toml2struct unmarshal toml file to struct.

## Example

    type Foo struct {
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
      EmbedA    Bar
    }
    
    type Bar struct {
      I int
    }
    
    var foo Foo
    toml2struct.Load("test.toml", "", &foo)
    
    // test.toml
    // 
    // IntA = 1
    // IntB = 2
    // IntC = "3"
    // IntHex = "0x0fffffff"
    // IntOct = "0755"
    // FloatA = 1.0
    // BoolA = false
    // BoolB = "true"
    // DurationA = "1h"
    // TimeA = "2006-01-02:15:04:05"
    // 
    // [EmbedA]
    // I = "99"
