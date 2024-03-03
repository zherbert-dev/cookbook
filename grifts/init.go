package grifts

import (
	"cookbook/actions"
	
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
