package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/player"
)

var janero ns.Obj
var janeroSpawn ns.Pointf

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
}

func janeroDialogueStart() {
	ns.AudioEvent(audio.MaleNPC1Talkable, janero)
	janero.LookAtObject(ns.GetCaller())
	janero_MayorTheogrinNeedsHelp()
	janero_TroubleAtTheManaMines()
}

func janero_MayorTheogrinNeedsHelp() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Character.Class {
	case player.Warrior:
		switch data.Quest.General.MayorTheogrinNeedsHelp {
		case QuestInactive, QuestComplete:
		default:
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard1Intro") // Greetings! You must be the great Warrior Horrendous dispatched to aid our beloved Village of Ix! Mayor Theogrin is expecting you.
		}
	case player.Wizard, player.Conjurer:
		switch data.Quest.General.MayorTheogrinNeedsHelp {
		case QuestInactive, QuestComplete:
		default:
			ns.TellStory(audio.ArcherHurt, "War03a:IxGuard1End") // The gates are unlocked so you may enter the Village. Delay no longer! The Mayor needs your help!
		}
	}
}

func janero_TroubleAtTheManaMines() {
	data := loadMyQuestData(ns.GetCaller().Player())
	switch data.Quest.General.TroubleAtTheManaMines {
	case QuestInactive, QuestComplete:
	default:
		ns.TellStory(audio.ArcherHurt, "Con03A.scr:IxGuard1") // The Mana Mines are to the west of the Crossroads. Just follow this path south to the Crossroads and then head west.
	}
}

func janeroDialogueEnd() {
	switch ns.GetAnswer(janero) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetJaneroDialogue() {
	ns.SetDialog(janero, ns.DialogNormal, janeroDialogueStart, janeroDialogueEnd)
}
