package noxworld

import "github.com/noxworld-dev/noxscript/ns/v4"

func checkExits() {
	if !exitUsed {
		checkIfNearbyExit()
	}

}

func checkIfNearbyExit() {
	if ns.Object("IxToCrossRoads") != nil {
		if (ns.InCirclef{Center: ns.Object("IxToCrossRoads"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "IxExitToCrossRoads"
				ns.LoadMap("nw_xroad", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	if ns.Object("IxCemToIx") != nil {
		if (ns.InCirclef{Center: ns.Object("IxCemToIx"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "IxCemExitToIx"
				ns.LoadMap("nw_ix", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	if ns.Object("DunMirToCrossRoads") != nil {
		if (ns.InCirclef{Center: ns.Object("DunMirToCrossRoads"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "DunMirExitToCrossRoads"
				ns.LoadMap("nw_xroad", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	if ns.Object("CrossRoadsToIx") != nil {
		if (ns.InCirclef{Center: ns.Object("CrossRoadsToIx"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "CrossRoadsExitToIx"
				ns.LoadMap("nw_ix", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	if ns.Object("CrossRoadsToMines") != nil {
		if (ns.InCirclef{Center: ns.Object("CrossRoadsToMines"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "CrossRoadsExitToMines"
				ns.LoadMap("nw_mines", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	if ns.Object("IxToTemple01") != nil {
		if (ns.InCirclef{Center: ns.Object("IxToTemple01"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "IxExitToTemple01"
				ns.LoadMap("nw_temple01", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	if ns.Object("Temple01ToIx") != nil {
		if (ns.InCirclef{Center: ns.Object("Temple01ToIx"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "Temple01ExitToIx"
				ns.LoadMap("nw_ix", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
	// Quest Maps
	if ns.Object("IxToQTemple") != nil {
		if (ns.InCirclef{Center: ns.Object("IxToQTemple"), R: 50}).Matches(ns.GetHost()) {
			exitUsed = true
			updateNoxWorldData(ns.GetHost().Player(), func(data *NoxWorldData) {
				data.Character.LastExitUsed = "IxExitToQTemple"
				ns.LoadMap("nw_qtemple", &ns.LoadMapOptions{HideTitleScreen: true})
			})
		}
	}
}

func moveToEntry() {
	data := loadMyNoxWorldData(ns.GetHost().Player())
	if data.Character.LastExitUsed == "IxExitToCrossRoads" {
		for i := 0; i < len(ns.Players()); i++ {
			ns.Players()[i].Unit().SetPos(ns.Waypoint("CrossRoadsEntryFromIx").Pos())
		}
	}
	if data.Character.LastExitUsed == "IxExitToIxCem" {
		for i := 0; i < len(ns.Players()); i++ {
			ns.Players()[i].Unit().SetPos(ns.Waypoint("IxCemEntryFromIx").Pos())
		}
	}
	if data.Character.LastExitUsed == "CrossRoadsExitToIx" {
		for i := 0; i < len(ns.Players()); i++ {
			ns.Players()[i].Unit().SetPos(ns.Waypoint("IxEntryFromCrossRoads").Pos())
		}
	}
	if data.Character.LastExitUsed == "DunMirExitToCrossRoads" {
		for i := 0; i < len(ns.Players()); i++ {
			ns.Players()[i].Unit().SetPos(ns.Waypoint("CrossRoadsEntryFromDunMir").Pos())
		}
	}
	if data.Character.LastExitUsed == "IxExitToTemple01" {
		for i := 0; i < len(ns.Players()); i++ {
			ns.Players()[i].Unit().SetPos(ns.Waypoint("Temple01EntryFromIx").Pos())
		}
	}
	if data.Character.LastExitUsed == "Temple01ExitToIx" {
		for i := 0; i < len(ns.Players()); i++ {
			ns.Players()[i].Unit().SetPos(ns.Waypoint("IxEntryFromTemple01").Pos())
		}
	}
}
