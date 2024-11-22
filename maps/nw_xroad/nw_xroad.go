package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {
		// StoryPic(Gvar18,"IxGuard1Pic") Janero
		//StoryPic(Gvar19,"IxGuard2Pic") Horst
	})
}

func CrossRoadsToIx() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "CrossRoadsToIx"
		ns.LoadMap("nw_ix", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}

func CrossRoadsToMines() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "CrossRoadsToMines"
		ns.LoadMap("nw_mines", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}

func CrossRoadsToDunMir() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "CrossRoadsToDunMir"
		ns.LoadMap("nw_dunmir", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}

func CrossRoadsToGalava01() {
	updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
		data.Character.LastExitUsed = "CrossRoadsToGalava01"
		ns.LoadMap("nw_galava01", &ns.LoadMapOptions{HideTitleScreen: true})
	})
}
