package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var rastur ns.Obj
var rasturSpawn ns.Pointf

func initRastur() {
	if ns.Object("Rastur") != nil {
		rastur = ns.Object("Rastur")
	} else {
		// Fix spawn
		rastur = ns.CreateObject("NPC", ns.GetHost())
	}
	rasturSpawn = rastur.Pos()
	ns.StoryPic(rastur, "WizardGuard1Pic")
	ns.SetDialog(rastur, ns.DialogNormal, rasturDialogueStart, rasturDialogueEnd)
}

func rasturDialogueStart() {
	rastur.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
		ns.TellStory(audio.Guard1Hurt, "War03a:GalavaGuardWarn") // This is the checkpoint for travelers who wish to go to Castle Galava. Warriors are not allowed beyond these gates.
	case player.Conjurer:
		// If quest is active.
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard1") // The Mana Mines are to the west of the Crossroads. Just return along the main path.
			return
		}
	case player.Wizard:
		// If quest is active.
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard1") // The Mana Mines are to the west of the Crossroads. Just return along the main path.
			return
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
	return
}
