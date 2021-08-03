package category_test

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/category"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidCategories(t *testing.T) {
	var categories = []struct {
		expected string
		wanted   error
	}{
		{"MBMR", nil},
		{"NBBR", nil},
		{"XKDX", nil},
		{"JBDX", nil},
	}

	for _, c := range categories {
		cat, err := category.FromValue(c.expected)

		assert.Nil(t, err)
		assert.Equal(t, c.wanted, err)
		assert.Equal(t, c.expected, cat.GetValue())
	}
}

func TestShouldReturnErrorOnInvalidCategories(t *testing.T) {
	var categories = []struct {
		expected string
		wanted   error
	}{
		{"ABBB", custom_error.InvalidArgument("Invalid Category ABBB, the number should a valid ACRISS code")},
		{"MBM", custom_error.InvalidArgument("Invalid Category MBM, the number should a valid ACRISS code")},
		{"", custom_error.InvalidArgument("Invalid Category , the number should a valid ACRISS code")},
		{" ", custom_error.InvalidArgument("Invalid Category  , the number should a valid ACRISS code")},
		{"1BBB", custom_error.InvalidArgument("Invalid Category 1BBB, the number should a valid ACRISS code")},
	}

	for _, c := range categories {
		_, err := category.FromValue(c.expected)

		assert.NotNil(t, err)
		assert.Equal(t, c.wanted.Error(),err.Error())
	}
}
