package usgr

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type payload struct {
	Name         string       `json:"name"`
	DeclaredType string       `json:"declaredType"`
	Scope        string       `json:"scope"`
	Value        payloadValue `json:"value"`
}

type payloadValue struct {
	TimeSeries []timeSeries `json:"timeSeries"`
}

type timeSeries struct {
	Variable variable `json:"variable"`
	Values   []value  `json:"values"`
}

type variable struct {
	Name        string         `json:"variableName"`
	Code        []variableCode `json:"variableCode"`
	Description string         `json:"variableDescription"`
}

type variableCode struct {
	Value string `json:"value"`
}

type valueCollection struct {
	Values []value `json:"values"`
}

type value struct {
	Value string `json:"value"`
}

func main() {
	url := "https://waterservices.usgs.gov/nwis/iv/?format=json&sites=03377500&indent=on&siteStatus=all&siteType=ST"

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "supersighs-usgsclient")

	res, err := client.Do(req)

	body, err := io.ReadAll(res.Body)

	result := payload{}
	jsonErr := json.Unmarshal(body, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, v := range result.Value.TimeSeries {
		fmt.Println(v.Variable.Name)
		for _, c := range v.Variable.Code {
			fmt.Println(c.Value)
		}
	}

}
