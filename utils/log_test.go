package utils

import "testing"

func TestLog(t *testing.T) {
	Log().Info("hahah")

	Log().Info("hahah1")
	Log().Info("hahah2")
	Log().Info("hahah3")
	Log().Clean()
}
