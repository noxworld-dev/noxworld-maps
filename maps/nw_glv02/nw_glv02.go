package nw_glv02

import (
	"github.com/noxworld-dev/noxscript/ns/v4"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func GalavaToBeachHouse() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapBeachHouse, nw.GoToMapOptions{
		Exit: "GalavaToBeachHouse",
	})
}

func GalavaToTowerOfIllusion01() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapTowerOfIllusion01, nw.GoToMapOptions{
		Exit: "GalavaToTowerOfIllusion01",
	})
}
