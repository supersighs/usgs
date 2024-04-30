package water

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
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

	Member struct {
		Id           string        `xml:"identifier"`
		Name         string        `xml:"name"`
		Date         string        `xml:"metadata>DocumentMetadata>generationDate"`
		Observations []Observation `xml:"observationMember>OM_Observation"`
	}

	Observation struct {
		ResultTime       string           `xml:"resultTime>TimeInstant>timePosition"`
		Unit             UnitOfMeasure    `xml:"result>MeasurementTimeseries>defaultPointMetadata>DefaultTVPMeasurementMetadata>uom"`
		Value            string           `xml:"result>MeasurementTimeseries>point>MeasurementTVP>value"`
		ObservedProperty ObservedProperty `xml:"observedProperty"`
		Position         string           `xml:"featureOfInterest>SF_SpatialSamplingFeature>shape>Point>pos"`
	}

	UnitOfMeasure struct {
		UnitName string `xml:"title,attr"`
	}

	ObservedProperty struct {
		Title string `xml:"title,attr"`
	}
)

func (m Member) String() string {
	return fmt.Sprintf("Member id=%v, name=%v, observations=%v", m.Id, m.Name, len(m.Observations))
}

func (f Feed) String() string {
	return fmt.Sprintf("Feed id=%v memberCount=%v", f.Id, len(f.Members))
}

func (o Observation) String() string {
	return fmt.Sprintf("Observation title=%v time=%v unit=%v value=%v", o.ObservedProperty.Title, o.ResultTime, o.Unit.UnitName, o.Value)
}

func GetFeed(url string) Feed {
	fmt.Println("Getting feed")

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	body, err := io.ReadAll(res.Body)
	result := Feed{}
	xmlErr := xml.Unmarshal(body, &result)

	if xmlErr != nil {
		log.Fatal(xmlErr)
	}

	return result
}
