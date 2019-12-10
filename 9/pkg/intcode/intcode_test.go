package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("test position mode", func(tt *testing.T) {
		source, err := Read("../../testdata/jumpposition.txt")
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
		source, err := Read("../../testdata/jumpimmediate.txt")
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
		source, err := Read("../../testdata/complicate.txt")
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

	t.Run("test equal numbers", func(tt *testing.T) {
		source, err := Read("../../testdata/equal.txt")
		assert.NoError(tt, err)
		p := New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(8)
			}
			if p.HasOutput() {
				assert.Equal(tt, 1, p.Output())
			}
		}

		p = New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(7)
			}
			if p.HasOutput() {
				assert.Equal(tt, 0, p.Output())
			}
		}
	})

	t.Run("test equal position", func(tt *testing.T) {
		source, err := Read("../../testdata/equalposition.txt")
		assert.NoError(tt, err)
		p := New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(8)
			}
			if p.HasOutput() {
				assert.Equal(tt, 1, p.Output())
			}
		}

		p = New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(7)
			}
			if p.HasOutput() {
				assert.Equal(tt, 0, p.Output())
			}
		}
	})

	t.Run("test less than", func(tt *testing.T) {
		source, err := Read("../../testdata/lessthan.txt")
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

		p = New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(9)
			}
			if p.HasOutput() {
				assert.Equal(tt, 0, p.Output())
			}
		}
	})

	t.Run("test less than position mode", func(tt *testing.T) {
		source, err := Read("../../testdata/lessthanposition.txt")
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

		p = New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(9)
			}
			if p.HasOutput() {
				assert.Equal(tt, 0, p.Output())
			}
		}
	})

	t.Run("test large numbers", func(tt *testing.T) {
		source, err := Read("../../testdata/largenum.txt")
		assert.NoError(tt, err)
		p := New(source)

		for !p.Halted() {
			err = p.Run()
			assert.NoError(tt, err)
			if p.RequiresInput() {
				p.Input(8)
			}
			if p.HasOutput() {
				assert.Equal(tt, 1187721666102244, p.Output())
			}
		}
	})

}
