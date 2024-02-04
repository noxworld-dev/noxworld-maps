package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
)

// Ideas:
// - add the smuggler's tunnel as a tunnel underneath the crossroads to circumvent the guards from all factions.

// General variables.
var playerCount []ns.Obj

var wizardClass []ns.Obj
var warriorClass []ns.Obj
var ConjurerClass []ns.Obj
var fireKnight []ns.Obj

// Initial server boot function.
func init() {
	ns.Music(22, 100)
	ns.OnChat(onCommand)
	ns.NewTimer(ns.Frames(5), func() {
		// Characters.
		initPriest()
		initBlacksmith()
		initCaptain()
		initOllie()
		initKenneth()
		initLance()
		initBrigadin()
		initJanero()
		initHorst()
		initKirik()
		initRastur()
		initMillard()
		initOsborn()
		initTest()
	})
	checkClass()
	ns.PrintStrToAll("Welcome to the NoxWorld server.")
	//ns.OnPlayerJoin(playerJoin)
	//ns.OnPlayerLeave(playerLeave)
	//ns.OnPlayerDeath(playerDeath)
}

func playerJoin() {

}

func playerLeave() {

}

func playerDeath() {

}

func checkClass() {
	// check the character's class and add them into the array.
	for i := 0; i < len(ns.Players()); i++ {
		mana := ns.Players()[i].Unit().MaxMana()
		if mana == 450 {
			wizardClass = append(wizardClass, ns.Players()[i].Unit())
		}
		if mana == 375 {
			ConjurerClass = append(ConjurerClass, ns.Players()[i].Unit())
		}
		if mana == 0 {
			warriorClass = append(warriorClass, ns.Players()[i].Unit())
		}
	}
}

// Server Commands.
func onCommand(t ns.Team, p ns.Player, obj ns.Obj, msg string) string {
	if p != nil {
		switch msg {
		case "login Ephreaym":
			if p.Name() == "Ephreaym" {
				p.Unit().DestroyChat()
			} else {
				ns.PrintStrToAll("Invalid username or password.")
			}
		}
	}
	return msg
}

func OnFrame() {
}

func endDialogue() {
	return
}

func forLoopExample() {
	//  for i := 0; i < len(arrayname); i++ {
	//		arrayname[i].Enchant(enchant.ANCHORED, ns.Frames(1))
	//	}
}

//var dataByObject = make(map[ns.Obj]*Data)
//
//type Data struct {
//   Level int
//}
//
//func getData(obj ns.Obj) *Data {
//   d := dataByObject[obj]
//   if d == nil {
//      d = new(Data)
//      dataByObject[obj] = d
//   }
//   return d
//}
