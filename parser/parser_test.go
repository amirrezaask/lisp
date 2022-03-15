package parser_test

import (
	"testing"

	"github.com/amirrezaask/lisp/parser"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		n, err := parser.Parse(`"Amirreza"`)
		assert.NoError(t, err)
		assert.Equal(t, &parser.Node{
			Type:  parser.NodeTypeString,
			Value: `Amirreza`,
		}, n)
	})

	t.Run("numbers", func(t *testing.T) {
		n, err := parser.Parse(`12121`)
		assert.NoError(t, err)
		assert.Equal(t, &parser.Node{
			Type:  parser.NodeTypeNumber,
			Value: 12121,
		}, n)
	})

	t.Run("mixed numbers and strings", func(t *testing.T) {
		n, err := parser.Parse(`12121"Amirreza"`)
		assert.Error(t, err)
		assert.Nil(t, n)
	})

	t.Run("lists", func(t *testing.T) {
		n, err := parser.Parse(`("a" 1)`)
		assert.NoError(t, err)
		assert.Equal(t, &parser.Node{
			Type: parser.NodeTypeList,
			Value: &parser.List{
				Value: []*parser.Node{
					{Type: parser.NodeTypeString, Value: "a"},
					{Type: parser.NodeTypeNumber, Value: 1},
				},
			},
		}, n)

	})

}
