package grifts

import (
	"github.com/constellation-project/goshortie/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
