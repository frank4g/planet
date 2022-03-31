package keeper

import (
	"github.com/deep2essence/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
