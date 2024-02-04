package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var millard ns.Obj
var millardSpawn ns.Pointf

func dialogExampleMillard() {
}

func initMillard() {
	if ns.Object("Millard") != nil {
		millard = ns.Object("Millard")
	} else {
		// Fix spawn
		millard = ns.CreateObject("NPC", ns.GetHost())
	}
	millardSpawn = millard.Pos()
	ns.StoryPic(millard, "MalePic2")
	ns.SetDialog(millard, ns.DialogNormal, millardDialogueStart, millardDialogueEnd)
	ns.Object("MineDoor1").Lock(true)
	ns.Object("MineDoor2").Lock(true)
}

func millardDialogueStart() {
	millard.LookAtObject(ns.GetCaller())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			ns.TellStory(audio.SwordsmanHurt, "War03a:MineGuard") // Sorry, the Mana Mines are closed.
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			ns.TellStory(audio.SwordsmanHurt, "War03a:MineGuard") // Sorry, the Mana Mines are closed.
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			ns.TellStory(audio.SwordsmanHurt, "War03a:MineGuard") // Sorry, the Mana Mines are closed.
			return
		}
	}
}

func millardDialogueManaMinesQuest() {
	// Start
	// Con03A.scr:MineGuardA	You have arrived! Lives are at stake! Please hurry to the mines, just west, down the path. The Foreman is waiting for you outside the Miner's Lodge. He'll direct you from there.
	// Con03A.scr:MineGuardB	The foreman will tell you the details of our crisis. He waits for you by the entrance of the mines.
	// End
	// Con03A.scr:MineGuardC	Your valor knows no peer. We all thank you for saving our men.
}

func millardDialogueEnd() {
	if ns.GetAnswer(millard) == 0 { // Goodbye
	}
	if ns.GetAnswer(millard) == 1 { // Yes
	}
	if ns.GetAnswer(millard) == 2 { // No
	}
}

func resetMillardDialogue() {
	return
}
