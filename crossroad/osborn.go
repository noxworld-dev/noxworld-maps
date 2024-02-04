package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
)

var osborn ns.Obj
var osbornSpawn ns.Pointf

func dialogExampleOsborn() {
}

func initOsborn() {
	if ns.Object("Osborn") != nil {
		osborn = ns.Object("Osborn")
	} else {
		// Fix spawn
		osborn = ns.CreateObject("NPC", ns.GetHost())
	}
	osbornSpawn = osborn.Pos()
	ns.StoryPic(osborn, "OsbornPic")
	ns.SetDialog(osborn, ns.DialogNormal, osbornDialogueStart, osbornDialogueEnd)
}

func osbornDialogueStart() {
	osborn.LookAtObject(ns.GetCaller())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
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

func osbornDialogueSpectaclesQuest() {
	// Con03A.scr:HermitHappy	My spectacles! You brought them back! May all that is great bless you! And please, take this scroll. It contains all I have learned about bats. It would be invaluable to any conjurer.
	// Con03A.scr:HermitMeet01	Gahhhhhh! No! Don't kill me! Oh. A young man?! I can't see well at all. But I know you're not one of those infernal bandits who stole my spectacles! I'm almost blind without them. If you could get them back, you'd save my life and I'd be eternally grateful.
	// Con03A.scr:HermitMeet02	Have you recovered my spectacles?! Oh.... Well, the rogues who took them have a hideout in the woods nearby.
}

func osbornDialogueEnd() {
	if ns.GetAnswer(osborn) == 0 { // Goodbye
	}
	if ns.GetAnswer(osborn) == 1 { // Yes
	}
	if ns.GetAnswer(osborn) == 2 { // No
	}
}

func resetOsbornDialogue() {
	return
}
