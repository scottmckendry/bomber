// Package ai contains functionality to render output using GenAI
package ai

import (
	"testing"

	"github.com/devops-kung-fu/common/util"
	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/bomber/models"
)

func TestRenderer_Render(t *testing.T) {
	output := util.CaptureOutput(func() {
		renderer := Renderer{}
		_ = renderer.Render(models.NewResults([]models.Package{}, models.Summary{}, []models.ScannedFile{}, []string{"GPL"}, "0.0.0", "test", ""))
	})
	assert.NotNil(t, output)
}
