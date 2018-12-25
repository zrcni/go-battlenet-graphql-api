// go:generate go run scripts/gqlgen.go -v

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/zrcni/go-bnet-graphql-api/bnet"
	"github.com/zrcni/go-bnet-graphql-api/utils"
)

type contextKey struct {
	name string
}

func makeURL(input CharacterQueryInput, ctx context.Context) string {
	bnetFields := utils.MapFieldsToBnetFields(ctx)
	fields := strings.Join(bnetFields, "%2C")

	return fmt.Sprintf("https://eu.api.blizzard.com/wow/character/%s/%s?fields=%s", "Nordrassil", "Nien", fields)
}

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, input CharacterQueryInput) (*Character, error) {
	url := makeURL(input, ctx)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bnetAuth := bnet.GetAuthFromContext(ctx)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bnetAuth.GetToken()))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	utils.WriteResponseBodyToJSONFile(body, "tmpcharacter")

	character := &Character{}

	err = json.Unmarshal(body, character)
	if err != nil {
		log.Println("Character response unmarshal:", err)
		return nil, err
	}

	return character, nil
}

// AverageItemLevel maps to JSON field items.averageItemLevelEquipped
func (c *Character) AverageItemLevel() *int {
	return c.Items.AverageItemLevelEquipped
}

// AverageItemLevelInBags maps to JSON field items.averageItemLevel
func (c *Character) AverageItemLevelInBags() *int {
	return c.Items.AverageItemLevel
}
