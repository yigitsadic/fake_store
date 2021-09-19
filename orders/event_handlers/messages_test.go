package event_handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_unmarshalMessage(t *testing.T) {
	t.Run("it should parse successfully", func(t *testing.T) {
		got, err := unmarshalMessage(`{ "reference_id": "344" }`)

		assert.Nil(t, err)
		assert.Equal(t, "344", got)
	})

	t.Run("it should return an error if something happen", func(t *testing.T) {
		got, err := unmarshalMessage(``)

		assert.NotNil(t, err)
		assert.Equal(t, "", got)
	})

	t.Run("it should return an error if reference is empty", func(t *testing.T) {
		got, err := unmarshalMessage(`{ "message": "success" }`)

		assert.NotNil(t, err)
		assert.Equal(t, "", got)
	})
}
