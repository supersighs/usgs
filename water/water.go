package water

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type (
	Feed struct {
		Name           xml.Name `xml:"FeatureCollection"`
		Id             string   `xml:"id,attr"`
		GmlSchema      string   `xml:"xmlns:gml,attr"`
		WmlSchema      string   `xml:"xmlns:wml2,attr"`
		SchemaLocation string   `xml:"xsi:schemaLocation,attr"`
		Members        []Member `xml:"featureMember>Collection"`
	}

	// Member represents a single monitoring station
	Member struct {
		Id           string        `xml:"identifier"`
		Name         string        `xml:"name"`
		Date         string        `xml:"metadata>DocumentMetadata>generationDate"`
		Observations []Observation `xml:"observationMember>OM_Observation"`
	}

	// Observation represents a single observation from a monitoring station
	Observation struct {
		Id                string            `xml:"id,attr"`
		ResultTime        string            `xml:"resultTime>TimeInstant>timePosition"`
		Unit              UnitOfMeasure     `xml:"result>MeasurementTimeseries>defaultPointMetadata>DefaultTVPMeasurementMetadata>uom"`
		Value             float64           `xml:"result>MeasurementTimeseries>point>MeasurementTVP>value"`
		ObservedProperty  ObservedProperty  `xml:"observedProperty"`
		FeatureOfInterest FeatureOfInterest `xml:"featureOfInterest"`
	}

	// FeatureOfInterest represents the location of the observation
	FeatureOfInterest struct {
		Title           string          `xml:"title,attr"`
		MonitoringPoint MonitoringPoint `xml:"MonitoringPoint"`
	}

	MonitoringPoint struct {
		Id string `xml:"id,attr"`
	}

	// Shape represents the location of the monitoring point
	Shape struct {
		Position string `xml:"shape>point>pos"`
	}

	UnitOfMeasure struct {
		UnitName string `xml:"title,attr"`
	}

	ObservedProperty struct {
		Title string `xml:"title,attr"`
	}
)

func (feed Feed) GetMember(id string) (station Member) {
	for _, station = range feed.Members {
		if strings.Contains(station.Id, id) {
			return
		}
	}
	return
}

func (member Member) GetObservation(id string) (observation Observation, err error) {
	for _, observation = range member.Observations {
		if strings.Contains(observation.Id, id) {
			return observation, nil
		}
	}
	return Observation{}, fmt.Errorf("Observation not found")
}

func getFeed(url string) (feed Feed, err error) {

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	// todo: defer these errors
	if err != nil {
		log.Fatal(err)
		return Feed{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return Feed{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return Feed{}, err
	}

	xmlErr := xml.Unmarshal(body, &feed)

	if xmlErr != nil {
		log.Fatal(xmlErr)
		return Feed{}, xmlErr
	}

	return
}
