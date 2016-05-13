package toml2struct

import (
	"github.com/BurntSushi/toml"
	"github.com/yangchenxing/go-map2struct"
	"path/filepath"
)

func Load(path, includeKey string, dest interface{}) error {
	data, err := loadMap(path, includeKey)
	if err != nil {
		return err
	}
	return map2struct.UnmarshalMap(dest, data)
}

func loadMap(path, includeKey string) (map[string]interface{}, error) {
	// load file
	data := make(map[string]interface{})
	_, err := toml.DecodeFile(path, &data)
	if err != nil {
		return nil, err
	}
	// check include
	if includeKey == "" {
		return data, nil
	}
	includes, ok := data[includeKey].([]interface{})
	if !ok {
		return data, nil
	}
	// load include
	for _, include := range includes {
		includePath, ok := include.(string)
		if !ok {
			continue
		}
		if !filepath.IsAbs(includePath) {
			includePath = filepath.Join(filepath.Dir(path), includePath)
		}
		inc, err := loadMap(includePath, includeKey)
		if err != nil {
			return nil, err
		}
		for key, value := range inc {
			data[key] = value
		}
	}
	// clean include key
	delete(data, includeKey)
	return data, nil
}
