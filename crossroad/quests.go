package noxworld

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v4"
)

type MyQuestData struct {
	// Warrior questline
	JoinTheFireKnights          bool
	JoinTheFireKnightsCompleted bool
	// Conjurer questline
	// Wizard questline
	BecomeTheWizardApprentice          bool
	BecomeTheWizardApprenticeCompleted bool
	// Crossroads
	LostSpectacles          bool
	LostSpectaclesCompleted bool
	// Mana Mines
	TroubleAtTheManaMines          bool
	TroubleAtTheManaMinesCompleted bool
	// Ix
	// Mayor Theogrin
	MayorTheogrinNeedsHelp          bool
	MayorTheogrinNeedsHelpCompleted bool
	// Dun Mir
	// Gearheart
	ResearchingMechanicalGolems          bool
	ResearchingMechanicalGolemsCompleted bool
	AlternativeEnergySource              bool
	AlternativeEnergySourceCompleted     bool
	SuitUp                               bool
	SuitUpCompleted                      bool
	TestDrive                            bool
	TestDriveCompleted                   bool
	Upgrade                              bool
	UpgradeCompleted                     bool
	// Example for template
	FollowUpQuestDialogue          bool
	FollowUpQuestDialogueCompleted bool

	// class data
	fireKnight bool
}

func loadMyQuestData(pl ns.Player) MyQuestData {
	var data MyQuestData
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Get("my-quest-name", &data)
	if err != nil {
		fmt.Println("cannot read quest data:", err)
	}
	return data
}

func saveMyQuestData(pl ns.Player, data MyQuestData) {
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Set("my-quest-name", &data)
	if err != nil {
		fmt.Println("cannot save quest data:", err)
	}
}

func updateMyQuestData(pl ns.Player, fnc func(data *MyQuestData)) {
	data := loadMyQuestData(pl)
	fnc(&data)
	saveMyQuestData(pl, data)
}
