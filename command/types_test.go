package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackwardFirstCommand(t *testing.T) {
	renderResult := Backward.render(1)
	expected := []int{4150, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 1284, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 1284, 642, 18042}
	assert.Equal(t, expected, renderResult)
}

func TestBackWardSecondCommand(t *testing.T) {
	renderResult := Backward.render(2)
	expected := []int{4150, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 1284, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 642, 1284, 642, 642, 642, 18042}
	assert.Equal(t, expected, renderResult)
}
