package facades

import (
	"log"

	"github.com/goravel-kit/geetest"
	"github.com/goravel-kit/geetest/contracts"
)

func Geetest() contracts.Geetest {
	instance, err := geetest.App.Make(geetest.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Geetest)
}
