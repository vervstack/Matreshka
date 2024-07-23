package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ExtractUpdateValue(t *testing.T) {
	t.Run("VALUE", func(t *testing.T) {
		valIn := "string"
		valOut, upsert, del := extractValue(valIn)
		require.Equal(t, valIn, valOut)
		require.True(t, upsert)
		require.False(t, del)
	})
	t.Run("VALUE_REF", func(t *testing.T) {
		valIn := "string"
		valOut, upsert, del := extractValue(&valIn)
		require.Equal(t, valIn, valOut)
		require.True(t, upsert)
		require.False(t, del)
	})

	t.Run("VALUE_NIL", func(t *testing.T) {
		var valIn *string
		valOut, upsert, del := extractValue(valIn)
		require.Nil(t, valOut)
		require.False(t, upsert)
		require.True(t, del)
	})
}
