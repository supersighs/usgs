package water

import (
	"encoding/xml"
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
		Id               string           `xml:"id,attr"`
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

func GetFeed(url string) Feed {

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := Feed{}
	xmlErr := xml.Unmarshal(body, &result)

	if xmlErr != nil {
		log.Fatal(xmlErr)
	}

	return result
}
