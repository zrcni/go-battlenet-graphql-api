package battlenet

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch makes a http get request to the Battle.net API using provided url and token
func Fetch(url string, authToken string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("new request: %v", err)
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("do request: %v", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("read body: %v", err)
		return nil, err
	}

	return body, nil
}

// MapRaceIDToRaceName maps ID number to character race name
func MapRaceIDToRaceName(id int) string {
	switch id {
	case 1:
		fallthrough
	case 33:
		return "Human"
	case 2:
		return "Orc"
	case 3:
		return "Dwarf"
	case 4:
		return "Night Elf"
	case 5:
		return "Undead"
	case 6:
		return "Tauren"
	case 7:
		return "Gnome"
	case 8:
		return "Troll"
	case 9:
		return "Goblin"
	case 10:
		return "Blood Elf"
	case 11:
		return "Draenei"
	case 22:
		return "Worgen"
	case 23:
		return "Gilnean"
	case 24:
		fallthrough
	case 25:
		fallthrough
	case 26:
		return "Pandaren"
	case 27:
		return "Nightborne"
	case 28:
		return "Highmountain Tauren"
	case 29:
		return "Void Elf"
	case 30:
		return "Lightforged Draenei"
	case 31:
		return "Zandalari Troll"
	case 32:
		return "Kul Tiran"
	case 34:
		return "Dark Iron Dwarf"
	case 35:
		return "Mag'har Orc"
	default:
		return "Unknown Race"
	}
}

// MapClassIDToClassName maps class ID number to character class name
func MapClassIDToClassName(id int) string {
	switch id {
	case 1:
		return "Warrior"
	case 2:
		return "Paladin"
	case 3:
		return "Hunter"
	case 4:
		return "Rogue"
	case 5:
		return "Priest"
	case 6:
		return "Death Knight"
	case 7:
		return "Shaman"
	case 8:
		return "Mage"
	case 9:
		return "Warlock"
	case 10:
		return "Monk"
	case 11:
		return "Druid"
	case 12:
		return "Demon Hunter"
	default:
		return "Unknown Class"
	}
}

// MapGenderIDToGenderName maps gender ID to character gender name
func MapGenderIDToGenderName(id int) string {
	switch id {
	case 0:
		return "Male"
	case 1:
		return "Female"
	default:
		return "Unknown Gender"
	}
}

// MapFactionIDToFactionName maps faction ID to character faction name
func MapFactionIDToFactionName(id int) string {
	switch id {
	case 0:
		return "Alliance"
	case 1:
		return "Horde"
	case 2:
		return "Neutral"
	default:
		return "Unknown Faction"
	}
}
