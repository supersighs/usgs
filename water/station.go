package water

import (
	"fmt"
	"strings"
)

const (
	sitesUrl = "https://waterservices.usgs.gov/nwis/iv/?format=waterml,2.0&sites=%v&siteStatus=all&siteType=ST"
)

func GetStations(siteIds []string) Feed {
	// convert the station ids to a comma separated string
	stations := strings.Join(siteIds, ",")
	// create the url
	url := fmt.Sprintf(sitesUrl, stations)
	// get the feed
	feed := getFeed(url)
	// convert the feed to stations
	return feed
}
