package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Run("should found the target", func(t *testing.T) {
		slice := make([]int64, 0, 5)
		slice = append(slice, 1, 2, 3, 4, 5)
		result, err := BinarySearch(slice, 3)

		assert.NoError(t, err)
		assert.Equal(t, result, int64(3))
	})
}
