package nw_caveix

import (
	"github.com/noxworld-dev/noxscript/ns/v4"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {})
}

func ConStartCaveIxToIx() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapIx, nw.GoToMapOptions{
		Exit: "ConStartCaveIxToIx",
	})
}

func ConStartCaveIxToCastle() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapG_Castle, nw.GoToMapOptions{
		Exit: "ConStartCaveIxToCastle",
	})
}
