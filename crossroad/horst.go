package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
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
	switch data.Character.Class {
	case player.Warrior:
		// if mayor needs help
		if !data.Quest.MayorTheogrinNeedsHelp && !data.Quest.MayorTheogrinNeedsHelpCompleted {
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2Intro") //	Good day, Warrior! Welcome to the Village of Ix! The Mayor is expecting you.
			return
		}
	case player.Conjurer:
		// if mayor needs help
		if !data.Quest.MayorTheogrinNeedsHelp && !data.Quest.MayorTheogrinNeedsHelpCompleted {
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2End") // The gates are unlocked so you may enter the Village. Delay no further. The Mayor needs your help!
			return
		}
	case player.Wizard:
		// if mayor needs help
		if !data.Quest.MayorTheogrinNeedsHelp && !data.Quest.MayorTheogrinNeedsHelpCompleted {
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard2End") // The gates are unlocked so you may enter the Village. Delay no further. The Mayor needs your help!
			return
		}
	}
}

func horstDialogueManaMinesQuest() {
	ns.TellStory(audio.HumanMaleEatApple, "Con03A.scr:IxGuard2") // Good luck in the mines, boy. Watch out for bandits on the way!
}

func horstDialogueEnd() {
	switch ns.GetAnswer(horst) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetHorstDialogue() {
	ns.SetDialog(horst, ns.DialogNormal, horstDialogueStart, horstDialogueEnd)
}
