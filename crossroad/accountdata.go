package noxworld

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v4"
)

type MyAccountData struct {
	Character struct {
		// General
		Registered bool
		Name       string
		Health     int
		Mana       int
		Level      int
		// Class
		Warrior  bool
		Conjurer bool
		Wizard   bool
		// Order
		FireKnight bool
	}
	Quest struct {
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
		// Tomb of Valor
		ExploreTheTombsOfValor          bool
		ExploreTheTombsOfValorCompleted bool
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
	}
}

func loadMyQuestData(pl ns.Player) MyAccountData {
	var data MyAccountData
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Get("my-quest-name", &data)
	if err != nil {
		fmt.Println("cannot read quest data:", err)
	}
	return data
}

func saveMyQuestData(pl ns.Player, data MyAccountData) {
	err := pl.Store(ns.Persistent{Name: "noxworld"}).Set("my-quest-name", &data)
	if err != nil {
		fmt.Println("cannot save quest data:", err)
	}
}

func updateMyQuestData(pl ns.Player, fnc func(data *MyAccountData)) {
	data := loadMyQuestData(pl)
	fnc(&data)
	saveMyQuestData(pl, data)
}
