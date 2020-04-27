package logger

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPANIC(t *testing.T) {
	t.Run("should not return panic", func(t *testing.T) {
		assert.NotPanics(t, func() { PANIC("Example error", nil) }, "The code did not panic")
	})
	t.Run("should return panic", func(t *testing.T) {
		assert.Panics(t, func() { PANIC("Example error", errors.New("error when start service")) }, "The code contains panic")
	})
}
