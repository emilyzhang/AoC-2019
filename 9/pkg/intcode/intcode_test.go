package intcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	t.Run("test position mode", func(tt *testing.T) {
		source, err := Read("jumpposition.txt")
		assert.NoError(tt, err)
		p := New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(7)
			}
			if p.HasOutput() {
				assert.Equal(tt, 1, p.Output())
			}
		}
	})

	t.Run("test immediate mode", func(tt *testing.T) {
		source, err := Read("jumpimmediate.txt")
		assert.NoError(tt, err)
		p := New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(7)
			}
			if p.HasOutput() {
				assert.Equal(tt, 1, p.Output())
			}
		}
	})

	t.Run("test complicated mode", func(tt *testing.T) {
		source, err := Read("complicate.txt")
		fmt.Println(source)
		assert.NoError(tt, err)
		p := New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(7)
			}
			if p.HasOutput() {
				assert.Equal(tt, 999, p.Output())
			}
		}
	})
}
