package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	err := order.Validate()
	if err == nil {
		t.Error("Expected Error, got Nil")
	}
}

func TestErrorIfIDIsBlankWithAssert(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}
