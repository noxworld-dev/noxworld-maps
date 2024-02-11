package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var rastur ns.Obj
var rasturSpawn ns.Pointf

func init() {
	OnLateInit(func() {
		if ns.Object("Rastur") != nil {
			rastur = ns.Object("Rastur")
		} else {
			// Fix spawn
			rastur = ns.CreateObject("NPC", ns.GetHost())
		}
		rasturSpawn = rastur.Pos()
		ns.StoryPic(rastur, "WizardGuard1Pic")
		ns.SetDialog(rastur, ns.DialogNormal, rasturDialogueStart, rasturDialogueEnd)
	})
}

func rasturDialogueStart() {
	ns.AudioEvent(audio.Guard1Talkable, rastur)
	rastur.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
		ns.TellStory(audio.Guard1Hurt, "War03a:GalavaGuardWarn") // This is the checkpoint for travelers who wish to go to Castle Galava. Warriors are not allowed beyond these gates.
	}
	rastur_TroubleAtTheManaMines()
}

func rastur_TroubleAtTheManaMines() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.General.TroubleAtTheManaMines {
	case QuestInactive, QuestComplete:
	default:
		switch data.Character.Class {
		case player.Wizard, player.Conjurer:
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard1") // The Mana Mines are to the west of the Crossroads. Just return along the main path.
		}
	}
}

func rasturDialogueEnd() {
	switch ns.GetAnswer(rastur) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetRasturDialogue() {
	ns.SetDialog(rastur, ns.DialogNormal, rasturDialogueStart, rasturDialogueEnd)
}
