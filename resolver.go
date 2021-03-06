// go:generate go run scripts/gqlgen.go -v

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/olivere/elastic"
	"github.com/zrcni/go-battlenet-graphql-api/battlenet"
	"github.com/zrcni/go-battlenet-graphql-api/elasticsearch"
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

func (r *queryResolver) Mounts(ctx context.Context, input SearchInput) ([]*Mount, error) {
	var query elastic.Query

	if input.Type != nil {
		switch *input.Type {
		case SearchTypeFuzzy:
			query = elastic.NewFuzzyQuery("name", *input.Term).Boost(1).Fuzziness(2)
		case SearchTypeRegexp:
			query = elastic.NewRegexpQuery("name", *input.Term)
		case SearchTypeWildcard:
			query = elastic.NewWildcardQuery("name", *input.Term)
		case SearchTypeNormal:
			fallthrough
		default:
			query = elastic.NewTermQuery("name", *input.Term)
		}
	} else {
		query = elastic.NewTermQuery("name", *input.Term)
	}

	searchResult, err := elasticsearch.Client.Search().Index("mounts").Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	var mounts []*Mount
	var mountType Mount
	for _, item := range searchResult.Each(reflect.TypeOf(mountType)) {
		m := item.(Mount)
		mounts = append(mounts, &m)
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
