package api

import "github.com/google/wire"

var APISet = wire.NewSet(
	GreetSet,
) // end
