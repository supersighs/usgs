package water

import (
	//"encoding/xml"
	"fmt"
	"strings"
)

type (
	Station struct {
		Id       string
		Readings []Reading
	}

	Reading struct {
		Id    string
		Title string
		Value string
		Unit  string
	}
)

func (reading Reading) String() string {
	return fmt.Sprintf("%v: %v %v", reading.Title, reading.Value, reading.Unit)
}

func (station Station) String() string {
	readings := make([]string, len(station.Readings))
	for i, reading := range station.Readings {
		readings[i] = reading.String()
	}
	return fmt.Sprintf("Station %v\n%v", station.Id, strings.Join(readings, "\n"))
}

func GetStations(stationIds []string) []Station {
	stations := strings.Join(stationIds, ",")
	url := fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=waterml,2.0&sites=%v&siteStatus=all&siteType=ST", stations)
	feed := GetFeed(url)

	return feed.ToStations()
}

func (feed Feed) ToStations() (stations []Station) {
	for _, member := range feed.Members {
		stations = append(stations, member.AsStation())
	}
	return
}

func (observation Observation) AsReading() Reading {
	return Reading{
		Id:    observation.Id,
		Title: observation.ObservedProperty.Title,
		Value: observation.Value,
		Unit:  observation.Unit.UnitName,
	}
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
