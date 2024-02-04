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
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			ns.TellStory(audio.HumanMaleEatFood, "War03a:GalavaGuardEnd") // Begone Warrior before I blast you.
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			if data.TroubleAtTheManaMines && !data.TroubleAtTheManaMinesCompleted {
				ns.TellStory(audio.FireKnight1Hurt, "Con03A.scr:GalavaGuard2") // Halt, Conjurer! You're supposed to go to the mines. The Mana shipment will be delayed if you don't get up there.
				return
			}
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

func kirikDialogueEnd() {
	if ns.GetAnswer(kirik) == 0 { // Goodbye
	}
	if ns.GetAnswer(kirik) == 1 { // Yes
	}
	if ns.GetAnswer(kirik) == 2 { // No
	}
}

func resetKirikDialogue() {
	ns.SetDialog(kirik, ns.DialogNormal, kirikDialogueStart, kirikDialogueEnd)
	return
}
