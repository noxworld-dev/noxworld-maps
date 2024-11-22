package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {
		ns.StoryPic(ns.Object("Henrick"), "HenrickPic")
		henrickInit()
		ns.StoryPic(ns.Object("Lydia"), "SnobbyGirlPic")
		lydiaInit()
		ns.StoryPic(ns.Object("Joyce"), "MaidenPic")
		joyceInit()
		ns.StoryPic(ns.Object("Jacob"), "WardenPic")
		jailerInit()
		ns.StoryPic(ns.Object("Tanya"), "MaidenPic3")
		tanyaInit()
		ns.StoryPic(ns.Object("Bryan"), "MalePic4")
		bryanInit()
		ns.StoryPic(ns.Object("Clyde"), "MalePic11")
		clydeInit()
		ns.StoryPic(ns.Object("Gretchen"), "MaidenPic4")
		gretchenInit()
	})
}

func gretchenInit() {
	ns.SetDialog(ns.Object("Gretchen"), ns.DialogNormal, gretchenDialogueStart, gretchenDialogueEnd)
}

func gretchenDialogueEnd() {
	ns.Object("Gretchen").Wander()
}

func gretchenDialogueStart() {
	ns.Object("Gretchen").Idle()
	ns.Object("Gretchen").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	data := loadMyNoxWorldData(p)
	switch data.Character.Class {
	case 0:
		ns.Object("Gretchen").ChatStr("Hello there, Warrior!")
	case 1:
		ns.Object("Gretchen").ChatStr("Hello there, Wizard!")
	case 2:
		ns.Object("Gretchen").ChatStr("Good day, Conjurer")
	}
	ns.NewTimer(ns.Seconds(3), func() {
		ns.Object("Gretchen").Wander()
	})
}

func clydeInit() {
	ns.SetDialog(ns.Object("Clyde"), ns.DialogNormal, clydeDialogueStart, clydeDialogueEnd)
}

func clydeDialogueEnd() {
	ns.Object("Clyde").Wander()
}

func clydeDialogueStart() {
	ns.Object("Clyde").Idle()
	ns.Object("Clyde").LookAtObject(ns.GetCaller())
	rnd := ns.Random(1, 2)
	switch rnd {
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman13Talk02")
	case 2:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman13Talk03")
	}
}

func bryanInit() {
	ns.SetDialog(ns.Object("Bryan"), ns.DialogNormal, bryanDialogueStart, bryanDialogueEnd)
}

func bryanDialogueEnd() {
	ns.Object("Bryan").Wander()
}

func bryanDialogueStart() {
	ns.Object("Bryan").Idle()
	ns.Object("Bryan").LookAtObject(ns.GetCaller())
	rnd := ns.Random(1, 3)
	switch rnd {
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman10Talk01")
	case 2:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman10Talk02")
	case 3:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:Townsman10Talk03")

	}
}

func tanyaInit() {
	ns.SetDialog(ns.Object("Tanya"), ns.DialogNormal, tanyaDialogueStart, tanyaDialogueEnd)
}

func tanyaDialogueEnd() {
	ns.Object("Tanya").Wander()
}

func tanyaDialogueStart() {
	ns.Object("Tanya").Idle()
	ns.Object("Tanya").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.HumanMaleEatFood, "Con02A:Maiden3Talk01")
}

func jailerInit() {
	ns.SetDialog(ns.Object("Jacob"), ns.DialogNormal, jailerDialogueStart, jailerDialogueEnd)
}

func jailerDialogueEnd() {
}

func jailerDialogueStart() {
	ns.Object("Jacob").LookAtObject(ns.GetCaller())
	rnd := ns.Random(1, 3)
	switch rnd {
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:JailerTalk01")
	case 2:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:JailerTalk02")
	case 3:
		ns.TellStory(audio.HumanMaleEatFood, "Con02A:JailerTalk03")
	}

}

func joyceInit() {
	ns.SetDialog(ns.Object("Joyce"), ns.DialogNormal, joyceDialogueStart, joyceDialogueEnd)
}

func joyceDialogueEnd() {
}

func joyceDialogueStart() {
	ns.Object("Joyce").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.HumanMaleEatFood, "Con02a:BarMaiden")
}

func lydiaInit() {
	ns.SetDialog(ns.Object("Lydia"), ns.DialogNormal, lydiaDialogueStart, lydiaDialogueEnd)
}

func lydiaDialogueEnd() {
}

func lydiaDialogueStart() {
	ns.Object("Lydia").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.HumanMaleEatFood, "Con02a:BarMaiden2")
}

func henrickInit() {
	ns.SetDialog(ns.Object("Henrick"), ns.DialogYesNo, henrickDialogueStart, henrickDialogueEnd)
}

func henrickDialogueStart() {
	ns.Object("Henrick").LookAtObject(ns.GetCaller())
	p := ns.GetCaller().Player()
	data := loadMyNoxWorldData(p)
	switch data.Character.WolfCompanion {
	case 0:
		ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickSalesPitchA")
	case 1:
		ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickSalesPitchB")
	case 2:
		ns.SetDialog(ns.Object("Henrick"), ns.DialogNormal, henrickDialogueStart, henrickDialogueEnd)
		ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickNoMoreWolves")
		henrickInit()
	}
}

func henrickDialogueEnd() {
	switch ns.GetAnswer(ns.Object("Henrick")) {
	case ns.AnswerGoodbye:
		henrickInit()
	case ns.AnswerYes:
		if ns.GetCaller().GetGold() < 200 {
			ns.SetDialog(ns.Object("Henrick"), ns.DialogNormal, henrickDialogueStart, henrickDialogueEnd)
			ns.TellStory(audio.HumanMaleEatFood, "war08b:HenrickNotEnoughGold")
			henrickInit()
		} else {
			ns.SetDialog(ns.Object("Henrick"), ns.DialogNormal, henrickDialogueStart, henrickDialogueEnd)
			updateNoxWorldData(ns.GetCaller().Player(), func(data *NoxWorldData) {
				data.Character.WolfCompanion++
			})
			ns.TellStory(audio.HumanMaleEatFood, "War08b:HenrickSaleSuccessful")
			ns.GetCaller().ChangeGold(-200)
			wolf := ns.CreateObject("Wolf", ns.Waypoint("WolfCreateWP"))
			wolf.SetOwner(ns.GetCaller())
			wolf.Follow(ns.GetCaller())
			henrickInit()
			wolf.OnEvent(ns.EventDeath, func() {})
		}
	case ns.AnswerNo:
		henrickInit()
	}
}
