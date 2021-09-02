package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		p := &Parser{code: `"Amirreza"`}
		n, err := p.Parse()
		assert.NoError(t, err)
		assert.Equal(t, &Node{
			Type:  NodeType_String,
			Value: `Amirreza`,
		}, n)
	})

	t.Run("numbers", func(t *testing.T) {
		p := &Parser{code: `12121`}
		n, err := p.Parse()
		assert.NoError(t, err)
		assert.Equal(t, &Node{
			Type:  NodeType_Number,
			Value: 12121,
		}, n)
	})

	t.Run("mixed numbers and strings", func(t *testing.T) {
		p := &Parser{code: `12121"Amirreza"`}
		n, err := p.Parse()
		assert.NoError(t, err)
		assert.Equal(t, &Node{
			Type:  NodeType_String,
			Value: "Amirreza",
		}, n)
	})

	t.Run("lists", func(t *testing.T) {
		p := &Parser{code: `("a" 1)`}
		n, err := p.Parse()
		assert.NoError(t, err)
		assert.Equal(t, &Node{
			Type:  NodeType_List,
			Value: &List{},
		}, n)

	})

}
