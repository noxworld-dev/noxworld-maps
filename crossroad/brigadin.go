package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var brigadin ns.Obj
var brigadinSpawn ns.Pointf

func initBrigadin() {
	if ns.Object("Brigadin") != nil {
		brigadin = ns.Object("Brigadin")
	} else {
		// Fix spawn
		brigadin = ns.CreateObject("NPC", ns.GetHost())
	}
	brigadinSpawn = brigadin.Pos()
	ns.StoryPic(brigadin, "Warrior4Pic")
	ns.SetDialog(brigadin, ns.DialogNormal, brigadinDialogueStart, brigadinDialogueEnd)
}

func brigadinDialogueStart() {
	brigadin.LookAtObject(ns.GetCaller())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			ns.TellStory(audio.SwordsmanHurt, "War03a:DunMirGuard3")
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			return
		}
	}
}

func brigadinDialogueEnd() {
	if ns.GetAnswer(brigadin) == 0 { // Goodbye
	}
	if ns.GetAnswer(brigadin) == 1 { // Yes
	}
	if ns.GetAnswer(brigadin) == 2 { // No
	}
}

func resetBrigadinDialogue() {
	return
}
