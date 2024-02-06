package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/player"
)

// Ideas:
// - add the smuggler's tunnel as a tunnel underneath the crossroads to circumvent the guards from all factions.

// General variables.
var playerCount []ns.Obj

var wizardClass []ns.Obj
var warriorClass []ns.Obj
var ConjurerClass []ns.Obj

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
	checkClass(ns.GetHost().Player())
	ns.PrintStrToAll("Welcome to the NoxWorld server.")
	ns.Runtime().OnPlayerJoin(playerJoin)

}

func playerJoin(p ns.Player) bool {
	ns.PrintStrToAll("NewPlayer joined")
	checkClass(p)
	return true
}

func checkClass(p ns.Player) {
	// check the character's class and add them into the array.
	switch p.Unit().GetClass() {
	case player.Wizard:
		wizardClass = append(wizardClass, p.Unit())
	case player.Conjurer:
		ConjurerClass = append(ConjurerClass, p.Unit())
	case player.Warrior:
		warriorClass = append(warriorClass, p.Unit())
	}
	updateMyQuestData(p, func(data *MyAccountData) {
		data.Character.Class = p.Unit().GetClass()
	})
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
