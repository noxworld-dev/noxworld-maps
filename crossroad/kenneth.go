package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

var kenneth ns.Obj
var kennethSpawn ns.Pointf
var lockDoorDunMir bool

func initKenneth() {
	if ns.Object("Kenneth") != nil {
		kenneth = ns.Object("Kenneth")
	} else {
		// Fix spawn
		kenneth = ns.CreateObject("NPC", ns.GetHost())
	}
	kennethSpawn = kenneth.Pos()
	ns.StoryPic(kenneth, "Warrior3Pic")
	ns.SetDialog(kenneth, ns.DialogNormal, kennethDialogueStart, kennethDialogueEnd)
	kennethManageDoorLock()
}

func kennethManageDoorLock() {
	search := ns.FindClosestObject(kenneth, ns.HasTypeName{"NewPlayer"}, ns.InCirclef{Center: kenneth, R: 300})
	if search != nil {
		for i := 0; i < len(wizardClass); i++ {
			if !lockDoorDunMir {
				if kenneth.Pos().Sub(wizardClass[i].Pos()).Len() <= 300 {
					lockDoorDunMir = true
					ns.Object("DunMirDoor1").Lock(true)
					ns.Object("DunMirDoor2").Lock(true)
				}
			}
		}
	} else {
		if lockDoorDunMir {
			lockDoorDunMir = false
			ns.Object("DunMirDoor1").Lock(false)
			ns.Object("DunMirDoor2").Lock(false)
		}
	}
	ns.NewTimer(ns.Seconds(1), func() {
		kennethManageDoorLock()
	})
}

func kennethDialogueStart() {
	kenneth.LookAtObject(ns.GetCaller())
	data := loadMyQuestData(ns.GetCaller().Player())
	// Warrior dialogue.
	for i := 0; i < len(warriorClass); i++ {
		if ns.GetCaller() == warriorClass[i] {
			ns.TellStory(audio.ArcherHurt, "War03a:DunMirGuard1") //  The Fortress of Dün Mir is home of the legendary Fire Knights.
			return
		}
	}
	// Conjurer dialogue.
	for i := 0; i < len(ConjurerClass); i++ {
		if ns.GetCaller() == ConjurerClass[i] {
			ns.TellStory(audio.ArcherHurt, "War03a:DunMirGuard1") //  The Fortress of Dün Mir is home of the legendary Fire Knights.
			return
		}
	}
	// Wizard dialogue.
	for i := 0; i < len(wizardClass); i++ {
		if ns.GetCaller() == wizardClass[i] {
			rnd := ns.Random(1, 3)
			if rnd == 1 {
				ns.AudioEvent(audio.TauntShakeFist, kenneth)
				kenneth.ChatStr("Go away Wizard!")
			}
			if rnd == 2 {
				if kenneth.Pos().Sub(wizardClass[i].Pos()).Len() <= 50 {
					kenneth.LookAtObject(wizardClass[i])
					kenneth.HitMelee(kenneth.Pos())
					ns.GetCaller().PushTo(kenneth, 20)
					ns.NewTimer(ns.Frames(15), func() {
						kenneth.Guard(kennethSpawn, kennethSpawn, 300)
					})
				} else {
					ns.AudioEvent(audio.TauntShakeFist, kenneth)
					kenneth.ChatStr("I smell rat!")
				}
			}
			if rnd == 3 {
				ns.AudioEvent(audio.TauntShakeFist, kenneth)
				kenneth.ChatStr("Beat it!")
			}
			return
		}
	}
}

func kennethDialogueEnd() {
	if ns.GetAnswer(kenneth) == 1 { // Yes
	}
	if ns.GetAnswer(kenneth) == 2 { // No
	}
}

func resetKennethDialogue() {
	ns.SetDialog(kenneth, ns.DialogNormal, kennethDialogueStart, kennethDialogueEnd)
	return
}
