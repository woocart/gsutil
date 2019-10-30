package pipeline

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalFile(t *testing.T) {
	ctx := context.Background()
	p, err := NewPipeline(ctx, "/local/file")

	assert.NoError(t, err)
	assert.False(t, p.IsCloud())
	assert.False(t, p.IsStdio())
	assert.True(t, p.IsLocal())
	assert.Equal(t, "/local/file", p.Path)
	assert.Equal(t, Local, p.Type)
	assert.Equal(t, "/local/file", p.String())

}
func TestStdio(t *testing.T) {
	ctx := context.Background()
	p, err := NewPipeline(ctx, "stdio")

	assert.NoError(t, err)
	assert.False(t, p.IsCloud())
	assert.True(t, p.IsStdio())
	assert.False(t, p.IsLocal())
	assert.Equal(t, "", p.Path)
	assert.Equal(t, STDIO, p.Type)
	assert.Equal(t, "", p.String())
}
