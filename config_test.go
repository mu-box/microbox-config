// -*- mode: go; tab-width: 2; indent-tabs-mode: 1; st-rulers: [70] -*-
// vim: ts=4 sw=4 ft=lua noet
package config

import (
	"testing"
)

func TestParser(test *testing.T) {
	// this test really needs to be finished
	defaults := map[string]string{
		"key": "value"}
	Load(defaults, "")

}
