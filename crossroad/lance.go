package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var lance ns.Obj
var lanceSpawn ns.Pointf

func init() {
	OnLateInit(func() {
		if ns.Object("Lance") != nil {
			lance = ns.Object("Lance")
		} else {
			// Fix spawn
			lance = ns.CreateObject("NPC", ns.GetHost())
		}
		lanceSpawn = lance.Pos()
		ns.StoryPic(lance, "Warrior2Pic")
		ns.SetDialog(lance, ns.DialogNormal, lanceDialogueStart, lanceDialogueEnd)
	})
}

func lanceDialogueStart() {
	ns.AudioEvent(audio.FireKnight1Talkable, lance)
	lance.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
		if data.Character.FireKnight {
			ns.TellStory(audio.FireKnight1Hurt, "War03a:DunMirGuard2") // You should be proud to serve Horrendous and the Fire Knights of Dün Mir.
		}
	case player.Wizard:
		switch ns.Random(1, 3) {
		case 1:
			ns.AudioEvent(audio.TauntShakeFist, lance)
			lance.ChatStr("Get lost!")
		case 2:
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
		case 3:
			ns.AudioEvent(audio.TauntShakeFist, lance)
			lance.ChatStr("No Wizards allowed!")
		}
	}
	lance_TroubleAtTheManaMines()
}

func lance_TroubleAtTheManaMines() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Conjurer, player.Warrior:
		switch data.Quest.General.TroubleAtTheManaMines {
		case QuestInactive, QuestComplete:
		default:
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:DunMirGuard2") // The Mana Mines are to the west of the Crossroads. You should return there.return
		}
	}
}

func lanceDialogueEnd() {
	switch ns.GetAnswer(lance) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetLanceDialogue() {
	ns.SetDialog(lance, ns.DialogNormal, lanceDialogueStart, lanceDialogueEnd)
}
