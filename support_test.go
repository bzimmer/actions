package actions_test

import (
	"testing"

	"github.com/bzimmer/actions"
)

func TestSupport(t *testing.T) {
	switch actions.Support() {
	case false:
		t.Fail()
	default:
	}
}
