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
	rastur.OnEvent(ns.EventEnemyHeard, func() { rasturCallBackup() })
	rastur.OnEvent(ns.EventEnemySighted, func() { rasturCallBackup() })
}

func rasturCallBackup() {

}

func rasturDialogueStart() {
	rastur.LookAtObject(ns.GetCaller())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			ns.TellStory(audio.Guard1Hurt, "War03a:GalavaGuardWarn") // This is the checkpoint for travelers who wish to go to Castle Galava. Warriors are not allowed beyond these gates.
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			// If mana mines quest active
			// Con03A.scr:GalavaGuard1
			//
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			// If mana mines quest active
			// Con03A.scr:GalavaGuard1
			//
			return
		}
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
	return
}
