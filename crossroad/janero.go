package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var janero ns.Obj
var janeroSpawn ns.Pointf

func dialogExampleJanero() {
}

func initJanero() {
	if ns.Object("Janero") != nil {
		janero = ns.Object("Janero")
	} else {
		// Fix spawn
		janero = ns.CreateObject("NPC", ns.GetHost())
	}
	janeroSpawn = janero.Pos()
	ns.StoryPic(janero, "IxGuard1Pic")
	ns.SetDialog(janero, ns.DialogNormal, janeroDialogueStart, janeroDialogueEnd)
	janero.OnEvent(ns.EventEnemyHeard, func() { janeroCallBackup() })
	janero.OnEvent(ns.EventEnemySighted, func() { janeroCallBackup() })
}

func janeroCallBackup() {

}

func janeroDialogueStart() {
	janero.LookAtObject(ns.GetCaller())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			// if mayor needs help
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard1Intro") // Greetings! You must be the great Warrior Horrendous dispatched to aid our beloved Village of Ix! Mayor Theogrin is expecting you.
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			// if mayor needs help
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard1End") // The gates are unlocked so you may enter the Village. Delay no longer! The Mayor needs your help!
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			// if mayor needs help
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard1End") // The gates are unlocked so you may enter the Village. Delay no longer! The Mayor needs your help!
			return
		}
	}
}

func janeroDialogueManaMinesQuest() {
	ns.TellStory(audio.ArcherHurt, "Con03A.scr:IxGuard1") // The Mana Mines are to the west of the Crossroads. Just follow this path south to the Crossroads and then head west.
}

func janeroDialogueEnd() {
	if ns.GetAnswer(janero) == 0 { // Goodbye
	}
	if ns.GetAnswer(janero) == 1 { // Yes
	}
	if ns.GetAnswer(janero) == 2 { // No
	}
}

func resetJaneroDialogue() {
	return
}
