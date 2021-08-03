package brand_test

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/brand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidBrands(t *testing.T) {
	var brands = []struct {
		expected string
		wanted   error
	}{
		{"", nil},
		{"123", nil},
		{"1a1a", nil},
		{"abcde", nil},
	}

	for _, b := range brands {
		br, err := brand.FromValue(b.expected)
		assert.Nil(t, err)
		assert.Equal(t, b.wanted, err)
		assert.Equal(t, b.expected, br.GetValue())
	}
}
