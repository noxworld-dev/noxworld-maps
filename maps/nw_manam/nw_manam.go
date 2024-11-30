package nw_manam

import (
	"github.com/noxworld-dev/noxscript/ns/v4"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {

	})
}

func ManaMinesToCrossRoads() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapCrossroads, nw.GoToMapOptions{
		Exit: "ManaMinesToCrossRoads",
	})
}
