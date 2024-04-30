package water

import (
	//"encoding/xml"
	"fmt"
	"strings"
)

type (
	Station struct {
		Id    string
		Value string
	}
)

func GetStations(stationIds []string) []Station {
	fmt.Sprintf("Getting data for %v", strings.Join(stationIds, ","))
	url := "https://waterservices.usgs.gov/nwis/iv/?format=waterml,2.0&sites=03377500,03378500&indent=on&siteStatus=all&siteType=ST"
	feed := GetFeed(url)
	for _, member := range feed.Members {
		fmt.Println(member.Name)
		for _, observation := range member.Observations {
			fmt.Println(observation)
		}
	}

	return mapToStations(feed)
}

func mapToStations(feed Feed) (stations []Station) {
	for _, member := range feed.Members {
		for _, observation := range member.Observations {
			stations = append(stations, fromObservation(observation))
		}
	}
	return
}

// maybe use a pointer here?
func fromObservation(obs Observation) Station {
	return Station{
		Id:    obs.ObservedProperty.Title,
		Value: obs.Value,
	}
}
