package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/object"
)

var priest ns.Obj
var priestSpawn ns.Pointf

func init() {
	OnLateInit(func() {
		priest = ns.Object("Priest")
		ns.StoryPic(priest, "GalavaPriestPic")
		ns.SetDialog(priest, ns.DialogYesNo, priestDialogueStart, priestDialogueEnd)
	})
}

func priestDialogueStart() {
	ns.AudioEvent(audio.Wizard1Talkable, priest)
	priest.LookAtObject(ns.GetCaller())
	ns.TellStory(audio.HumanMaleEatFood, "Would you like to make a donation of 50.000 gold?")
}

func priestDialogueEnd() {
	switch ns.GetAnswer(priest) {
	case ns.AnswerYes:
		gold := ns.GetCaller().GetGold()
		if gold >= 50000 {
			ns.GetCaller().ChangeGold(-50000)
			Ankh := ns.CreateObject("Ankh", ns.GetCaller().Pos())
			Ankh.FlagsEnable(object.FlagNoPushCharacters)
			Ankh.DeleteAfter(ns.Frames(1))
		} else {
			ns.SetDialog(priest, ns.DialogNormal, endPriestDialogue, resetPriestDialogue)
			ns.TellStory(audio.SwordsmanHurt, "I'm sorry. You don't have enough gold. Come back when you've got enough.")
		}
	case ns.AnswerNo:
		// No
	}
}

func endPriestDialogue() {
}

func resetPriestDialogue() {
	ns.SetDialog(priest, ns.DialogYesNo, priestDialogueStart, priestDialogueEnd)
}

// Priest end.
