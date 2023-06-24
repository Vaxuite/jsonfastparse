package jsonfastparse

import "fmt"

func ParseStringMap(bytes []byte) (map[string]string, error) {
	parsedMap := map[string]string{}
	keyFirstIndex := -1
	keySecondIndex := -1
	valueFirstIndex := -1
	for x := 1; x < len(bytes); x++ {
		if bytes[x] == ':' && bytes[x+1] != '"' {
			return nil, fmt.Errorf("not map[string]string")
		}
		if bytes[x] != '"' || bytes[x-1] == '\\' {
			continue
		}
		if keyFirstIndex == -1 {
			keyFirstIndex = x
		} else if keySecondIndex == -1 {
			keySecondIndex = x
		} else if valueFirstIndex == -1 {
			valueFirstIndex = x
		} else {
			parsedMap[string(bytes[keyFirstIndex+1:keySecondIndex])] = string(bytes[valueFirstIndex+1 : x])
			keyFirstIndex = -1
			keySecondIndex = -1
			valueFirstIndex = -1
		}
	}
	return parsedMap, nil
}
