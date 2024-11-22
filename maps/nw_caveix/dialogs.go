package con01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar8        int
	gvar9        = 1
	captainState int
	captain      ns.ObjectID
	gatekeep2    ns.ObjectID
	gatekeep3    ns.ObjectID
)

func InitializeDialogs() {
	initDialogs()
}

func initDialogs() {
	captain = ns.Object("Airship_Captain")
	gatekeep2 = ns.Object("Ryan")
	gatekeep3 = ns.Object("Ed")

	ns.StoryPic(captain, "AirshipCaptainPic")
	ns.StoryPic(gatekeep2, "IxGuard2Pic")
	ns.StoryPic(gatekeep3, "IxGuard1Pic")
	ns.SetDialog(gatekeep2, ns.NORMAL, Gatekeeper2DialogStart, Gatekeeper2DialogEnd)
	ns.SetDialog(gatekeep3, ns.NORMAL, Gatekeeper3DialogStart, Gatekeeper3DialogEnd)
}

func OwnObjects() {
	ns.SetOwner(ns.GetHost(), captain)
	ns.SetOwner(ns.GetHost(), gatekeep2)
	ns.SetOwner(ns.GetHost(), gatekeep3)
}

func StartCaptainConversation() {
	ns.SetDialog(captain, ns.NORMAL, CaptainDialogStart, CaptainDialogEnd)
	ns.StartDialog(captain, ns.GetHost())
}

func CaptainDialogStart() {
	ns.LookAtObject(captain, ns.GetHost())
	switch captainState {
	case 0:
		ns.TellStory(ns.SwordsmanHurt, "Con01a:CaptainGreet")
	case 1:
		ns.TellStory(ns.SwordsmanHurt, "Con01a:CaptainProd")
	}
}

func CaptainDialogEnd() {
	switch captainState {
	case 0:
		ns.JournalEntry(ns.GetHost(), "LocateAldwyn", 2)
		ns.PrintToAll("Con01a:NewJournalEntry")
		captainState = 1
	case 1:
		// nop
	}
}

func Gatekeeper2DialogStart() {
	ns.LookAtObject(ns.GetTrigger(), ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:Gatekeeper2Greet")
}

func Gatekeeper2DialogEnd() {
}

func Gatekeeper3DialogStart() {
	ns.LookAtObject(ns.GetTrigger(), ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:Gatekeeper3Greet")
}

func Gatekeeper3DialogEnd() {
}
