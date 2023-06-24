package jsonfastparse

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("single key value", func(t *testing.T) {
		singleKey := map[string]string{
			"key": "value",
		}
		bytes, err := json.Marshal(singleKey)
		require.NoError(t, err)
		parsed, err := ParseStringMap(bytes)
		require.NoError(t, err)
		assert.Equal(t, singleKey, parsed)
	})

	t.Run("many key value", func(t *testing.T) {
		singleKey := map[string]string{
			"key":  "value",
			"key2": "value2",
		}
		bytes, err := json.Marshal(singleKey)
		require.NoError(t, err)
		parsed, err := ParseStringMap(bytes)
		require.NoError(t, err)
		assert.Equal(t, singleKey, parsed)
	})

	t.Run("map string int", func(t *testing.T) {
		singleKey := map[string]int{
			"key": 1,
		}
		bytes, err := json.Marshal(singleKey)
		require.NoError(t, err)
		fmt.Println(string(bytes))
		parsed, err := ParseStringMap(bytes)
		require.Errorf(t, err, "not map[string]string")
		assert.Nil(t, parsed)
	})
}

func BenchmarkParseStringMap(b *testing.B) {
	singleKey := map[string]string{
		"key":  "value",
		"key2": "value2",
	}
	bytes, _ := json.Marshal(singleKey)
	b.Run("custom parser", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ParseStringMap(bytes)
		}
	})

	b.Run("std lib parser", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var m map[string]string
			json.Unmarshal(bytes, &m)
		}
	})

}
