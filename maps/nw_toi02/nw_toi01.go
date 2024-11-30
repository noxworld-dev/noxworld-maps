package nw_toi02

import (
	"github.com/noxworld-dev/noxscript/ns/v4"

	nw "github.com/noxworld-dev/noxworld-maps/noxworld"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {
	})
}

func TowerOfIllusion02ToTowerOfIllusion01() {
	nw.GoToMap(ns.GetHost().Player(), nw.MapTowerOfIllusion01, nw.GoToMapOptions{
		Exit: "TowerOfIllusion02ToTowerOfIllusion01",
	})
}
