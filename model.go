package api

// CharacterItems represents JSON response's items field from Battle.net API
type CharacterItems struct {
	AverageItemLevel         *int  `json:"averageItemLevel"`
	AverageItemLevelEquipped *int  `json:"averageItemLevelEquipped"`
	Head                     *Item `json:"head"`
	Neck                     *Item `json:"neck"`
	Shoulder                 *Item `json:"shoulder"`
	Back                     *Item `json:"back"`
	Chest                    *Item `json:"chest"`
	Shirt                    *Item `json:"shirt"`
	Tabard                   *Item `json:"tabard"`
	Wrist                    *Item `json:"wrist"`
	Hands                    *Item `json:"hands"`
	Waist                    *Item `json:"waist"`
	Legs                     *Item `json:"legs"`
	Feet                     *Item `json:"feet"`
	Finger1                  *Item `json:"finger1"`
	Finger2                  *Item `json:"finger2"`
	Trinket1                 *Item `json:"trinket1"`
	Trinket2                 *Item `json:"trinket2"`
	MainHand                 *Item `json:"mainHand"`
	OffHand                  *Item `json:"offHand"`
}

// Character represents JSON response from the Battle.net API
type Character struct {
	LastModified        *int    `json:"lastModified"`
	Name                *string `json:"name"`
	Realm               *string `json:"realm"`
	Battlegroup         *string `json:"battlegroup"`
	ClassID             *int    `json:"class"`
	RaceID              *int    `json:"race"`
	GenderID            *int    `json:"gender"`
	Level               *int    `json:"level"`
	AchievementPoints   *int    `json:"achievementPoints"`
	Thumbnail           *string `json:"thumbnail"`
	CalcClass           *string `json:"calcClass"`
	FactionID           *int    `json:"faction"`
	TotalHonorableKills *int    `json:"totalHonorableKills"`
	// Used when unmarshaling JSON response into a struct in Feed resolver
	TempFeed    []map[string]interface{} `json:"feed"`
	Pets        *CharacterPets           `json:"pets"`
	Mounts      *CharacterMounts         `json:"mounts"`
	Items       *CharacterItems          `json:"items"`
	Professions *CharacterProfessions    `json:"professions"`
	Reputation  []*Reputation            `json:"reputation"`
	Stats       *CharacterStats          `json:"stats"`
}
