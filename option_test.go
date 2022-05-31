package reflecthelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	t.Run("init option", func(t *testing.T) {
		opt := NewOption()
		assert.Equal(t, byte(0), opt.FloatFormat)
	})
	t.Run("functional options", func(t *testing.T) {
		opt := NewOption().Assign(
			WithBaseSystem(10),
			WithBitSize(64, 128),
			WithFloatConfiguration(-1, 't'),
			WithTimeLayouts("hello"),
			WithCustomAssigner(nil, true),
		)
		assert.Equal(t, byte('g'), opt.FloatFormat)
		opt = NewOption().Assign(
			WithBaseSystem(10),
			WithBitSize(64, 128),
			WithFloatConfiguration(-1, 'f'),
			WithTimeLayouts("hello"),
			WithCustomAssigner(nil, true),
		)
		assert.Equal(t, byte('f'), opt.FloatFormat)
	})
}

func TestNewDefaultOption(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "new default option",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultOption()
			assert.NotNil(t, got)
		})
	}
}
