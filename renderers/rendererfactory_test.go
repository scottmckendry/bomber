package renderers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/bomber/renderers/ai"
	"github.com/devops-kung-fu/bomber/renderers/html"
	"github.com/devops-kung-fu/bomber/renderers/json"
	"github.com/devops-kung-fu/bomber/renderers/md"
	"github.com/devops-kung-fu/bomber/renderers/stdout"
)

func TestNewRenderer(t *testing.T) {
	renderer, err := NewRenderer("stdout")
	assert.NoError(t, err)
	assert.IsType(t, stdout.Renderer{}, renderer)

	renderer, err = NewRenderer("json")
	assert.NoError(t, err)
	assert.IsType(t, json.Renderer{}, renderer)

	renderer, err = NewRenderer("html")
	assert.NoError(t, err)
	assert.IsType(t, html.Renderer{}, renderer)

	renderer, err = NewRenderer("ai")
	assert.NoError(t, err)
	assert.IsType(t, ai.Renderer{}, renderer)

	renderer, err = NewRenderer("md")
	assert.NoError(t, err)
	assert.IsType(t, md.Renderer{}, renderer)

	_, err = NewRenderer("test")
	assert.Error(t, err)
}
