package noxworld

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

func init() {
	ns.NewTimer(ns.Frames(1), func() {
		ns.StoryPic(ns.Object("IxPriest"), "IxPriestPic")
		IxPriestInit()
	})
}

func IxPriestInit() {
	ns.SetDialog(ns.Object("IxPriest"), ns.DialogNext, IxPriestDialogueStart, IxPriestDialogueEnd)
}

func IxPriestDialogueStart() {
	ns.SetDialog(ns.Object("IxPriest"), ns.DialogNext, IxPriestDialoguePartTwo, IxPriestDialoguePartTwo)
	ns.Object("IxPriest").Idle()
	ns.Object("IxPriest").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.SwordsmanHurt, "Con08a:PriestGreet")
}

func IxPriestDialoguePartTwo() {
	ns.SetDialog(ns.Object("IxPriest"), ns.DialogNext, IxPriestDialoguePartTwo, IxPriestDialoguePartThree)
	ns.Object("IxPriest").Idle()
	ns.Object("IxPriest").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.SwordsmanHurt, "Con08a:PriestProd")
}

func IxPriestDialoguePartThree() {
	ns.SetDialog(ns.Object("IxPriest"), ns.DialogNext, IxPriestDialoguePartThree, IxPriestDialoguePartFour)
	ns.Object("IxPriest").Idle()
	ns.Object("IxPriest").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.SwordsmanHurt, "Con08a:PriestProd2")
}

func IxPriestDialoguePartFour() {
	ns.SetDialog(ns.Object("IxPriest"), ns.DialogNormal, IxPriestDialoguePartFour, IxPriestDialogueEnd)
	ns.Object("IxPriest").Idle()
	ns.Object("IxPriest").LookAtObject(ns.GetCaller())
	ns.TellStory(audio.SwordsmanHurt, "Con08a:PriestProd3")
}

func IxPriestDialogueEnd() {
	IxPriestInit()
}
