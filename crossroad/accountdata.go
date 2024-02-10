package noxworld

import (
	"fmt"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/player"
)

type MyAccountData struct {
	Character struct {
		// General
		Registered bool
		Name       string
		Health     int
		Mana       int
		Level      int
		Class      player.Class
		// Order
		FireKnight      bool
		OrderOfOblivion bool
	}
	Quest struct {
		// 1 == quest accept
		// 2 == quest active
		// 3 == quest completed
		General struct {
			// Osborn
			LostSpectacles int
			// Mana Mines
			TroubleAtTheManaMines int
			// Tomb of Valor
			ExploreTheTombsOfValor int
			// Ix
			// Mayor Theogrin
			MayorTheogrinNeedsHelp int

			// Dun Mir
			// Gearheart
			// Mech suit questline
			//ResearchingMechanicalGolems          bool
			//ResearchingMechanicalGolemsCompleted bool
			//AlternativeEnergySource              bool
			//AlternativeEnergySourceCompleted     bool
			//SuitUp                               bool
			//SuitUpCompleted                      bool
			//TestDrive                            bool
			//TestDriveCompleted                   bool
			//Upgrade                              bool
			//UpgradeCompleted                     bool
			// Example for template
			FollowUpQuestDialogue int
		}
		Warrior struct {
			// Warrior questline
			JoinTheFireKnights int
		}
		Conjurer struct {
			// Conjurer questline
			BecomeTheConjurerApprentice int
		}
		Wizard struct {
			// Wizard questline
			BecomeTheWizardApprentice        int
			TravelToTheApprenticeHouse       int
			FindApprentice                   int
			TellHorvathYouFoundTheApprentice int
			GoToHorvathHisOffice             int
		}
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
