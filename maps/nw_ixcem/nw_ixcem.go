package nw_ixcem

import (
	"github.com/noxworld-dev/noxscript/ns/v4"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {

	})
}

func IxCemeteryToIx() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapIx, nw.GoToMapOptions{
		Exit: "IxCemeteryToIx",
	})
}
