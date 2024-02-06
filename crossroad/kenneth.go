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
	if data.Character.Warrior {
		ns.TellStory(audio.ArcherHurt, "War03a:DunMirGuard1") //  The Fortress of Dün Mir is home of the legendary Fire Knights.
		return
	}
	// Conjurer dialogue.
	if data.Character.Conjurer {
		ns.TellStory(audio.ArcherHurt, "War03a:DunMirGuard1") //  The Fortress of Dün Mir is home of the legendary Fire Knights.
		return
	}
	// Wizard dialogue.
	if data.Character.Wizard {
		switch ns.Random(1, 3) {
		case 1:
			ns.AudioEvent(audio.TauntShakeFist, kenneth)
			kenneth.ChatStr("Go away Wizard!")
		case 2:
			if kenneth.Pos().Sub(ns.GetCaller().Pos()).Len() <= 50 {
				kenneth.LookAtObject(ns.GetCaller())
				kenneth.HitMelee(kenneth.Pos())
				ns.GetCaller().PushTo(kenneth, 20)
				ns.NewTimer(ns.Frames(15), func() {
					kenneth.Guard(kennethSpawn, kennethSpawn, 300)
				})
			} else {
				ns.AudioEvent(audio.TauntShakeFist, kenneth)
				kenneth.ChatStr("I smell rat!")
			}
		case 3:
			ns.AudioEvent(audio.TauntShakeFist, kenneth)
			kenneth.ChatStr("Beat it!")
		}
		return
	}
}

func kennethDialogueEnd() {
	switch ns.GetAnswer(kenneth) {
	case ns.AnswerGoodbye:
		// Goodbye
	case ns.AnswerYes:
		// Yes
	case ns.AnswerNo:
		// No
	}
}

func resetKennethDialogue() {
	ns.SetDialog(kenneth, ns.DialogNormal, kennethDialogueStart, kennethDialogueEnd)
	return
}
