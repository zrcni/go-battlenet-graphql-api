package api

import (
	"encoding/json"
	"log"
)

func mapDataToInterface(m map[string]interface{}, i interface{}) {
	item, err := json.Marshal(m)
	if err != nil {
		log.Printf("mapDataToInterface marshal: %v", err)
	}

	if err := json.Unmarshal(item, i); err != nil {
		log.Printf("mapDataToInterface unmarshal: %v", err)
	}
}

func mapItemsToValidType(items []CharacterFeedItem) []*CharacterFeedItem {
	var validFeedItems []*CharacterFeedItem

	for _, itm := range items {
		validFeedItems = append(validFeedItems, &itm)
	}

	return validFeedItems
}
