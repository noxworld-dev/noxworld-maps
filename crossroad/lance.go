package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var lance ns.Obj
var lanceSpawn ns.Pointf

func initLance() {
	if ns.Object("Lance") != nil {
		lance = ns.Object("Lance")
	} else {
		// Fix spawn
		lance = ns.CreateObject("NPC", ns.GetHost())
	}
	lanceSpawn = lance.Pos()
	ns.StoryPic(lance, "Warrior2Pic")
	ns.SetDialog(lance, ns.DialogNormal, lanceDialogueStart, lanceDialogueEnd)
}

func lanceDialogueStart() {
	lance.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Warrior dialogue.
	if data.Character.Warrior {
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:DunMirGuard2") // The Mana Mines are to the west of the Crossroads. You should return there.return
			return
		} else if data.Character.FireKnight {
			ns.TellStory(audio.FireKnight1Hurt, "War03a:DunMirGuard2") // You should be proud to serve Horrendous and the Fire Knights of Dün Mir.
			return
		}
		return

	}
	// Conjurer dialogue.
	if data.Character.Conjurer {
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:DunMirGuard2") // The Mana Mines are to the west of the Crossroads. You should return there.return
		}
		return
	}
	// Wizard dialogue.
	if data.Character.Wizard {
		rnd := ns.Random(1, 3)
		if rnd == 1 {
			ns.AudioEvent(audio.TauntShakeFist, lance)
			lance.ChatStr("Get lost!")
		}
		if rnd == 2 {
			if lance.Pos().Sub(ns.GetCaller().Pos()).Len() <= 50 {
				lance.LookAtObject(ns.GetCaller())
				lance.HitMelee(lance.Pos())
				ns.GetCaller().PushTo(lance, 20)
				ns.NewTimer(ns.Frames(15), func() {
					lance.Guard(lanceSpawn, lanceSpawn, 300)
				})
			} else {
				ns.AudioEvent(audio.TauntShakeFist, lance)
				lance.ChatStr("You rat!")
			}
		}
		if rnd == 3 {
			ns.AudioEvent(audio.TauntShakeFist, lance)
			lance.ChatStr("No Wizards allowed!")
		}
		return
	}
}

func lanceDialogueEnd() {
	if ns.GetAnswer(lance) == 1 { // Yes
	}
	if ns.GetAnswer(lance) == 2 { // No
	}
}

func resetLanceDialogue() {
	ns.SetDialog(lance, ns.DialogNormal, lanceDialogueStart, lanceDialogueEnd)
	return
}
