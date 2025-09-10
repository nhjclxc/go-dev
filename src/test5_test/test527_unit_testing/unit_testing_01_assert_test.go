package test527_unit_testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// go get github.com/stretchr/testify/assert

func TestAssert01(t *testing.T) {
	res := 2 + 3
	res2 := 2 + 5

	assert.NotNil(t, res)
	assert.NotNil(t, res2)

	//assert.Equal(t, 5, res, "2+3 should be equal 5")
	assert.Equal(t, 5, res2, "2+3 should be equal 5")

}
