package water

import (
	"fmt"
	"strings"
)

type (
	Stations []Station

	Station struct {
		Id       string
		Readings []Reading
	}

	Reading struct {
		Id    string
		Title string
		Value float64
		Unit  string
	}
)

// GetStation returns a Station from a list of Stations by id
func (stations Stations) GetStation(id string) Station {
	for _, station := range stations {
		if strings.Contains(station.Id, id) {
			return station
		}
	}
	return Station{}
}

func GetStations(stationIds []string) Stations {
	// convert the station ids to a comma separated string
	stations := strings.Join(stationIds, ",")
	// create the url
	url := fmt.Sprintf(sitesUrl, stations)
	// get the feed
	feed := getFeed(url)
	// convert the feed to stations
	return feed.ToStations()
}

func (station Station) GetReading(id string) Reading {
	for _, reading := range station.Readings {
		if strings.Contains(reading.Id, id) {
			return reading
		}
	}
	return Reading{}
}

func (observation Observation) AsReading() Reading {
	return Reading{
		Id:    observation.Id,
		Title: observation.ObservedProperty.Title,
		Value: observation.Value,
		Unit:  observation.Unit.UnitName,
	}
}

func (feed Feed) ToStations() (stations []Station) {
	for _, member := range feed.Members {
		stations = append(stations, member.AsStation())
	}
	return
}

func (member Member) AsStation() Station {

	readings := make([]Reading, len(member.Observations))

	for i, obs := range member.Observations {
		readings[i] = obs.AsReading()
	}

	return Station{
		Id:       member.Id,
		Readings: readings,
	}
}
