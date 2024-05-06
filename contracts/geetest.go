package contracts

import "github.com/goravel-kit/geetest"

type Geetest interface {
	Verify(ticket geetest.Ticket) (bool, error)
}
