// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"strconv"
)

type Achievement struct {
	ID          *int                   `json:"id"`
	Title       *string                `json:"title"`
	Points      *int                   `json:"points"`
	Description *string                `json:"description"`
	Icon        *string                `json:"icon"`
	Criteria    []*AchievementCriteria `json:"criteria"`
	AccountWide *bool                  `json:"accountWide"`
	FactionID   *int                   `json:"factionId"`
}

type AchievementCriteria struct {
	ID          *int    `json:"id"`
	Description *string `json:"description"`
	OrderIndex  *int    `json:"orderIndex"`
	Max         *int    `json:"max"`
}

type AzeriteEmpoweredItem struct {
	AzeritePowers []*AzeritePower `json:"azeritePowers"`
}

type AzeriteItem struct {
	AzeriteLevel               *int `json:"azeriteLevel"`
	AzeriteExperience          *int `json:"azeriteExperience"`
	AzeriteExperienceRemaining *int `json:"azeriteExperienceRemaining"`
}

type AzeritePower struct {
	ID          *int `json:"id"`
	Tier        *int `json:"tier"`
	SpellID     *int `json:"spellId"`
	BonusListID *int `json:"bonusListId"`
}

type CharacterFeedAchievement struct {
	Type           *CharacterFeedItemType `json:"type"`
	Timestamp      *int                   `json:"timestamp"`
	Achievement    *Achievement           `json:"achievement"`
	FeatOfStrength *bool                  `json:"featOfStrength"`
}

func (CharacterFeedAchievement) IsCharacterFeedItem() {}

type CharacterFeedBossKill struct {
	Type           *CharacterFeedItemType `json:"type"`
	Timestamp      *int                   `json:"timestamp"`
	Achievement    *Achievement           `json:"achievement"`
	FeatOfStrength *bool                  `json:"featOfStrength"`
	Criteria       *AchievementCriteria   `json:"criteria"`
	Quantity       *int                   `json:"quantity"`
	Name           *string                `json:"name"`
}

func (CharacterFeedBossKill) IsCharacterFeedItem() {}

type CharacterFeedCriteria struct {
	Type        *CharacterFeedItemType `json:"type"`
	Timestamp   *int                   `json:"timestamp"`
	Achievement *Achievement           `json:"achievement"`
}

func (CharacterFeedCriteria) IsCharacterFeedItem() {}

type CharacterFeedItem interface {
	IsCharacterFeedItem()
}

type CharacterFeedLoot struct {
	Type       *CharacterFeedItemType `json:"type"`
	Timestamp  *int                   `json:"timestamp"`
	ItemID     *int                   `json:"itemId"`
	Context    *string                `json:"context"`
	BonusLists []*int                 `json:"bonusLists"`
}

func (CharacterFeedLoot) IsCharacterFeedItem() {}

type CharacterMounts struct {
	NumCollected    *int   `json:"numCollected"`
	NumNotCollected *int   `json:"numNotCollected"`
	Collected       []*Pet `json:"collected"`
}

type CharacterPets struct {
	NumCollected    *int   `json:"numCollected"`
	NumNotCollected *int   `json:"numNotCollected"`
	Collected       []*Pet `json:"collected"`
}

type CharacterProfessions struct {
	Primary   []*Profession `json:"primary"`
	Secondary []*Profession `json:"secondary"`
}

type CharacterQueryInput struct {
	Name   string `json:"name"`
	Realm  string `json:"realm"`
	Region string `json:"region"`
}

type CharacterStats struct {
	Health                      *int     `json:"health"`
	PowerType                   *string  `json:"powerType"`
	Power                       *int     `json:"power"`
	Str                         *int     `json:"str"`
	Agi                         *int     `json:"agi"`
	Int                         *int     `json:"int"`
	Sta                         *int     `json:"sta"`
	SpeedRating                 *float64 `json:"speedRating"`
	SpeedRatingBonus            *float64 `json:"speedRatingBonus"`
	Crit                        *float64 `json:"crit"`
	CritRating                  *float64 `json:"critRating"`
	Haste                       *float64 `json:"haste"`
	HasteRating                 *float64 `json:"hasteRating"`
	HasteRatingPercent          *float64 `json:"hasteRatingPercent"`
	Mastery                     *float64 `json:"mastery"`
	MasteryRating               *float64 `json:"masteryRating"`
	Leech                       *float64 `json:"leech"`
	LeechRating                 *float64 `json:"leechRating"`
	LeechRatingBonus            *float64 `json:"leechRatingBonus"`
	Versatility                 *float64 `json:"versatility"`
	VersatilityDamageDoneBonus  *float64 `json:"versatilityDamageDoneBonus"`
	VersatilityHealingDoneBonus *float64 `json:"versatilityHealingDoneBonus"`
	VersatilityDamageTakenBonus *float64 `json:"versatilityDamageTakenBonus"`
	AvoidanceRating             *float64 `json:"avoidanceRating"`
	AvoidanceRatingBonus        *float64 `json:"avoidanceRatingBonus"`
	SpellPen                    *float64 `json:"spellPen"`
	SpellCrit                   *float64 `json:"spellCrit"`
	SpellCritRating             *float64 `json:"spellCritRating"`
	Mana5                       *float64 `json:"mana5"`
	Mana5Combat                 *float64 `json:"mana5Combat"`
	Armor                       *float64 `json:"armor"`
	Dodge                       *float64 `json:"dodge"`
	DodgeRating                 *float64 `json:"dodgeRating"`
	Parry                       *float64 `json:"parry"`
	ParryRating                 *float64 `json:"parryRating"`
	Block                       *float64 `json:"block"`
	BlockRating                 *float64 `json:"blockRating"`
	MainHandDmgMin              *float64 `json:"mainHandDmgMin"`
	MainHandDmgMax              *float64 `json:"mainHandDmgMax"`
	MainHandSpeed               *float64 `json:"mainHandSpeed"`
	MainHandDps                 *float64 `json:"mainHandDps"`
	OffHandDmgMin               *float64 `json:"offHandDmgMin"`
	OffHandDmgMax               *float64 `json:"offHandDmgMax"`
	OffHandSpeed                *float64 `json:"offHandSpeed"`
	OffHandDps                  *float64 `json:"offHandDps"`
	RangedDmgMin                *float64 `json:"rangedDmgMin"`
	RangedDmgMax                *float64 `json:"rangedDmgMax"`
	RangedSpeed                 *float64 `json:"rangedSpeed"`
	RangedDps                   *float64 `json:"rangedDps"`
}

type Item struct {
	ID                   *int                  `json:"id"`
	Name                 *string               `json:"name"`
	Icon                 *string               `json:"icon"`
	Quality              *int                  `json:"quality"`
	ItemLevel            *int                  `json:"itemLevel"`
	TooltipParams        *ItemTooltipParams    `json:"tooltipParams"`
	Stats                []*ItemStat           `json:"stats"`
	Armor                *int                  `json:"armor"`
	Context              *string               `json:"context"`
	BonusLists           []*int                `json:"bonusLists"`
	ArtifactID           *int                  `json:"artifactId"`
	DisplayInfoID        *int                  `json:"displayInfoId"`
	ArtifactAppearanceID *int                  `json:"artifactAppearanceId"`
	Appearance           *TransmogItem         `json:"appearance"`
	AzeriteItem          *AzeriteItem          `json:"azeriteItem"`
	AzeriteEmpoweredItem *AzeriteEmpoweredItem `json:"azeriteEmpoweredItem"`
}

type ItemStat struct {
	Stat   *int `json:"stat"`
	Amount *int `json:"amount"`
}

type ItemTooltipParams struct {
	TransmogItem      *int `json:"transmogItem"`
	TimewalkerLevel   *int `json:"timewalkerLevel"`
	AzeritePower0     *int `json:"azeritePower0"`
	AzeritePower1     *int `json:"azeritePower1"`
	AzeritePower2     *int `json:"azeritePower2"`
	AzeritePower3     *int `json:"azeritePower3"`
	AzeritePower4     *int `json:"azeritePower4"`
	AzeritePowerLevel *int `json:"azeritePowerLevel"`
}

type Mount struct {
	Name       *string `json:"name"`
	SpellID    *int    `json:"spellId"`
	CreatureID *int    `json:"creatureId"`
	ItemID     *int    `json:"itemId"`
	QualityID  *int    `json:"qualityId"`
	Icon       *string `json:"icon"`
	IsGround   *bool   `json:"isGround"`
	IsFlying   *bool   `json:"isFlying"`
	IsAquatic  *bool   `json:"isAquatic"`
	IsJumping  *bool   `json:"isJumping"`
}

type Pet struct {
	Name                        *string   `json:"name"`
	SpellID                     *int      `json:"spellId"`
	CreatureID                  *int      `json:"creatureId"`
	ItemID                      *int      `json:"itemId"`
	QualityID                   *int      `json:"qualityId"`
	Icon                        *string   `json:"icon"`
	Stats                       *PetStats `json:"stats"`
	BattlePetGUID               *string   `json:"battlePetGuid"`
	IsFavorite                  *bool     `json:"isFavorite"`
	IsFirstAbilitySlotSelected  *bool     `json:"isFirstAbilitySlotSelected"`
	IsSecondAbilitySlotSelected *bool     `json:"isSecondAbilitySlotSelected"`
	IsThirdAbilitySlotSelected  *bool     `json:"isThirdAbilitySlotSelected"`
	CreatureName                *string   `json:"creatureName"`
	CanBattle                   *bool     `json:"canBattle"`
}

type PetStats struct {
	SpeciesID    *int `json:"speciesId"`
	BreedID      *int `json:"breedId"`
	PetQualityID *int `json:"petQualityId"`
	Level        *int `json:"level"`
	Health       *int `json:"health"`
	Power        *int `json:"power"`
	Speed        *int `json:"speed"`
}

type Profession struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Icon    *string `json:"icon"`
	Rank    *int    `json:"rank"`
	Max     *int    `json:"max"`
	Recipes []*int  `json:"recipes"`
}

type Reputation struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Standing *int    `json:"standing"`
	Value    *int    `json:"value"`
	Max      *int    `json:"max"`
}

type TransmogItem struct {
	ItemID                      *int `json:"itemId"`
	ItemAppearanceModID         *int `json:"itemAppearanceModId"`
	TransmogItemAppearanceModID *int `json:"transmogItemAppearanceModId"`
}

type CharacterFeedItemType string

const (
	CharacterFeedItemTypeLoot        CharacterFeedItemType = "LOOT"
	CharacterFeedItemTypeBosskill    CharacterFeedItemType = "BOSSKILL"
	CharacterFeedItemTypeCriteria    CharacterFeedItemType = "CRITERIA"
	CharacterFeedItemTypeAchievement CharacterFeedItemType = "ACHIEVEMENT"
)

func (e CharacterFeedItemType) IsValid() bool {
	switch e {
	case CharacterFeedItemTypeLoot, CharacterFeedItemTypeBosskill, CharacterFeedItemTypeCriteria, CharacterFeedItemTypeAchievement:
		return true
	}
	return false
}

func (e CharacterFeedItemType) String() string {
	return string(e)
}

func (e *CharacterFeedItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CharacterFeedItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CharacterFeedItemType", str)
	}
	return nil
}

func (e CharacterFeedItemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}