// go:generate go run scripts/gqlgen.go -v

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/zrcni/go-battlenet-graphql-api/battlenet"
	"github.com/zrcni/go-battlenet-graphql-api/utils"
)

type contextKey struct {
	name string
}

func makeBaseURL(region string) string {
	if region != "" {
		if region == "cn" {
			return "gateway.battlenet.com.cn"
		}
		return fmt.Sprintf("%s.api.blizzard.com", region)
	}
	return `eu.api.blizzard.com`
}

func makeURL(ctx context.Context, input CharacterQueryInput) string {
	battlenetFields := utils.MapFieldsToBattleNetFields(ctx)
	fields := strings.Join(battlenetFields, "%2C")
	baseURL := makeBaseURL(input.Region)

	return fmt.Sprintf("https://%s/wow/character/%s/%s?fields=%s", baseURL, input.Realm, input.Name, fields)
}

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, input CharacterQueryInput) (*Character, error) {
	url := makeURL(ctx, input)

	battlenetAuth := battlenet.GetAuthFromContext(ctx)
	authToken := battlenetAuth.GetToken()

	body, err := battlenet.Fetch(url, authToken)
	if err != nil {
		log.Printf("Get.queryBattleNet: %v", err)
		return nil, err
	}
	// utils.WriteResponseBodyToJSONFile(body, "tmpcharacter")

	character := &Character{}

	if err := json.Unmarshal(body, character); err != nil {
		log.Printf("Character response unmarshal: %v", err)
		return nil, err
	}

	return character, nil
}

func (c *Character) Class() (string, error) {
	return battlenet.MapClassIDToClassName(*c.ClassID), nil
}

func (c *Character) Faction() (string, error) {
	return battlenet.MapFactionIDToFactionName(*c.FactionID), nil
}

func (c *Character) Gender() (string, error) {
	return battlenet.MapGenderIDToGenderName(*c.GenderID), nil
}

func (c *Character) Race() (string, error) {
	return battlenet.MapRaceIDToRaceName(*c.RaceID), nil
}

// Feed of character activity
func (c *Character) Feed() ([]*CharacterFeedItem, error) {
	var feedItems []CharacterFeedItem

	for _, feedItem := range c.TempFeed {
		switch feedItem["type"] {
		case "LOOT":
			i := &CharacterFeedLoot{}
			mapDataToInterface(feedItem, i)
			feedItems = append(feedItems, i)
		case "BOSSKILL":
			i := &CharacterFeedBossKill{}
			mapDataToInterface(feedItem, i)
			feedItems = append(feedItems, i)
		case "CRITERIA":
			i := &CharacterFeedCriteria{}
			mapDataToInterface(feedItem, i)
			feedItems = append(feedItems, i)
		case "ACHIEVEMENT":
			i := &CharacterFeedAchievement{}
			mapDataToInterface(feedItem, i)
			feedItems = append(feedItems, i)
		}
	}

	validFeedItems := mapItemsToValidType(feedItems)

	return validFeedItems, nil
}

// AverageItemLevel maps to JSON field items.averageItemLevelEquipped
func (c *Character) AverageItemLevel() *int {
	return c.Items.AverageItemLevelEquipped
}

// AverageItemLevelInBags maps to JSON field items.averageItemLevel
func (c *Character) AverageItemLevelInBags() *int {
	return c.Items.AverageItemLevel
}

func mapDataToInterface(m map[string]interface{}, i interface{}) {
	item, err := json.Marshal(m)
	if err != nil {
		log.Printf("mapDataToInterface marshal: %v", err)
	}

	if err := json.Unmarshal(item, i); err != nil {
		log.Printf("mapDataToInterface unmarshal:", err)
	}
}

func mapItemsToValidType(items []CharacterFeedItem) []*CharacterFeedItem {
	var validFeedItems []*CharacterFeedItem

	for _, itm := range items {
		validFeedItems = append(validFeedItems, &itm)
	}

	return validFeedItems
}
