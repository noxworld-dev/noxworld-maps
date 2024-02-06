package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var kirik ns.Obj
var kirikSpawn ns.Pointf
var lockGateGalava bool

func initKirik() {
	if ns.Object("Kirik") != nil {
		kirik = ns.Object("Kirik")
	} else {
		// Fix spawn
		kirik = ns.CreateObject("NPC", ns.GetHost())
	}
	kirikSpawn = kirik.Pos()
	ns.StoryPic(kirik, "WizardGuard2Pic")
	ns.SetDialog(kirik, ns.DialogNormal, kirikDialogueStart, kirikDialogueEnd)
	kirikManageDoorLock()
}

func kirikManageDoorLock() {
	search := ns.FindClosestObject(kenneth, ns.HasTypeName{"NewPlayer"}, ns.InCirclef{Center: kirik, R: 300})
	if search != nil {
		for i := 0; i < len(warriorClass); i++ {
			if !lockGateGalava {
				if kirik.Pos().Sub(warriorClass[i].Pos()).Len() <= 300 {
					lockGateGalava = true
					ns.Object("GalavaGate1").Lock(true)
					ns.Object("GalavaGate2").Lock(true)
				}
			}
		}
	} else {
		if lockGateGalava {
			lockGateGalava = false
			ns.Object("GalavaGate1").Lock(false)
			ns.Object("GalavaGate2").Lock(false)
		}
	}
	ns.NewTimer(ns.Seconds(1), func() {
		kirikManageDoorLock()
	})
}

func kirikDialogueStart() {
	kirik.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Warrior dialogue.
	if data.Character.Warrior {
		ns.TellStory(audio.HumanMaleEatFood, "War03a:GalavaGuardEnd") // Begone Warrior before I blast you.
		return
	}
	// Conjurer dialogue.
	if data.Character.Conjurer {
		if data.Quest.TroubleAtTheManaMines && !data.Quest.TroubleAtTheManaMinesCompleted {
			ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard2") // Halt, Conjurer! You're supposed to go to the mines. The Mana shipment will be delayed if you don't get up there.
			return
		}
		return
	}
	// Wizard dialogue.
	if data.Character.Wizard {
		return
	}
}

func kirikDialogueEnd() {
	switch ns.GetAnswer(kirik) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetKirikDialogue() {
	ns.SetDialog(kirik, ns.DialogNormal, kirikDialogueStart, kirikDialogueEnd)
}
