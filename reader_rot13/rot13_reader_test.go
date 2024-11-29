package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestRot13Reader_Read(t *testing.T) {
	t.Run("decode ok", func(t *testing.T) {
		source := "Lbh penpxrq gur pbqr!"
		expected := "You cracked the code!"

		s := strings.NewReader(source)
		rot13 := rot13Reader{s}
		var res []byte
		res, err := io.ReadAll(rot13)

		assert.NoError(t, err, "не должно быть ошибки")

		assert.Equal(t, expected, string(res), "декодирование было верным")
	})
}
