package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
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
	// Warrior dialogue.
	if data.Character.Warrior {
		ns.TellStory(audio.Guard1Hurt, "War03a:GalavaGuardWarn") // This is the checkpoint for travelers who wish to go to Castle Galava. Warriors are not allowed beyond these gates.
		return
	}
	// Conjurer dialogue.
	if data.Character.Conjurer {
		// If quest is active.
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard1") // The Mana Mines are to the west of the Crossroads. Just return along the main path.
			return
		}
		return
	}
	// Wizard dialogue.
	if data.Character.Wizard {
		// If quest is active.
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard1") // The Mana Mines are to the west of the Crossroads. Just return along the main path.
			return
		}
		return
	}
}

func rasturDialogueEnd() {
	if ns.GetAnswer(rastur) == 0 { // Goodbye
	}
	if ns.GetAnswer(rastur) == 1 { // Yes
	}
	if ns.GetAnswer(rastur) == 2 { // No
	}
}

func resetRasturDialogue() {
	ns.SetDialog(rastur, ns.DialogNormal, rasturDialogueStart, rasturDialogueEnd)
	return
}
