// go:generate go run scripts/gqlgen.go -v

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/mapstructure"
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

func (r *queryResolver) Mounts(ctx context.Context, searchTerm string) ([]*Mount, error) {
	mountSearchURL := fmt.Sprintf("http://localhost:9200/mounts/_search?q=name:%s", searchTerm)

	body, err := utils.Fetch(mountSearchURL)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	var mountHits []ElasticSearchMountHit

	m, ok := data["hits"].(map[string]interface{})
	if !ok {
		return nil, nil
	}

	err = mapstructure.Decode(m["hits"], &mountHits)
	if err != nil {
		return nil, err
	}

	// []Mount into []*Mount
	var mounts []*Mount
	for _, hit := range mountHits {
		var mount Mount
		mapstructure.Decode(hit.Source, &mount)
		mounts = append(mounts, &mount)
	}

	return mounts, nil
}

func (m *Mount) Icons() *Icons {
	smallIcon := fmt.Sprintf("https://wow.zamimg.com/images/wow/icons/small/%s.jpg", *m.Icon)
	mediumIcon := fmt.Sprintf("https://wow.zamimg.com/images/wow/icons/medium/%s.jpg", *m.Icon)
	largeIcon := fmt.Sprintf("https://wow.zamimg.com/images/wow/icons/large/%s.jpg", *m.Icon)

	return &Icons{
		Small:  &smallIcon,
		Medium: &mediumIcon,
		Large:  &largeIcon,
	}
}

func (m *Mount) WowheadURL() *string {
	if *m.ItemID == 0 {
		spellURL := fmt.Sprintf("https://www.wowhead.com/spell=%v", *m.SpellID)
		return &spellURL
	}

	itemURL := fmt.Sprintf("https://www.wowhead.com/item=%v", *m.ItemID)
	return &itemURL
}

func (c *Character) Class() string {
	return battlenet.MapClassIDToName(*c.ClassID)
}

func (c *Character) Faction() string {
	return battlenet.MapFactionIDToName(*c.FactionID)
}

func (c *Character) Gender() string {
	return battlenet.MapGenderIDToName(*c.GenderID)
}

func (c *Character) Race() string {
	return battlenet.MapRaceIDToName(*c.RaceID)
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
