package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var horst ns.Obj
var horstSpawn ns.Pointf

func dialogExampleHorst() {
}

func initHorst() {
	if ns.Object("Horst") != nil {
		horst = ns.Object("Horst")
	} else {
		// Fix spawn
		horst = ns.CreateObject("NPC", ns.GetHost())
	}
	horstSpawn = horst.Pos()
	ns.StoryPic(horst, "IxGuard2Pic")
	ns.SetDialog(horst, ns.DialogNormal, horstDialogueStart, horstDialogueEnd)
}

func horstDialogueStart() {
	horst.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			// if mayor needs help
			if !data.MayorTheogrinNeedsHelp && !data.MayorTheogrinNeedsHelpCompleted {
				ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2Intro") //	Good day, Warrior! Welcome to the Village of Ix! The Mayor is expecting you.
				return
			}
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			// if mayor needs help
			if !data.MayorTheogrinNeedsHelp && !data.MayorTheogrinNeedsHelpCompleted {
				ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2End") // The gates are unlocked so you may enter the Village. Delay no further. The Mayor needs your help!
				return
			}
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			// if mayor needs help
			if !data.MayorTheogrinNeedsHelp && !data.MayorTheogrinNeedsHelpCompleted {
				ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2End") // The gates are unlocked so you may enter the Village. Delay no further. The Mayor needs your help!
				return
			}
			return
		}
	}
}

func horstDialogueManaMinesQuest() {
	ns.TellStory(audio.HumanMaleEatApple, "Con03A.scr:IxGuard2") // Good luck in the mines, boy. Watch out for bandits on the way!
}

func horstDialogueEnd() {
	if ns.GetAnswer(horst) == 0 { // Goodbye
	}
	if ns.GetAnswer(horst) == 1 { // Yes
	}
	if ns.GetAnswer(horst) == 2 { // No
	}
}

func resetHorstDialogue() {
	ns.SetDialog(horst, ns.DialogNormal, horstDialogueStart, horstDialogueEnd)
	return
}
