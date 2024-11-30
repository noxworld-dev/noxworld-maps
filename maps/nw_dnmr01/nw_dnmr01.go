package nw_dnmr01

import (
	"github.com/noxworld-dev/noxscript/ns/v4"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {
	})
}

func DunMir01ToCrossRoads() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapCrossroads, nw.GoToMapOptions{
		Exit: "DunMir01ToCrossRoads",
	})
}
