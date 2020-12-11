package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//GenUID -
func TestGenUID(t *testing.T) {
	id := GenUID()
	assert.True(t, id != "")
}
