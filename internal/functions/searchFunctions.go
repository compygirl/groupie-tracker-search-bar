package groupieTrecker

import (
	"strconv"
	"strings"

	data "groupieTrecker/internal/data"
)

// we should format locations(capitalize them and replace "_" by " ")
func SearchSuggestions(searchTerm string) ([]string, map[int]string) {
	set := make(map[int]string, 0)

	suggestions := make([]string, 0) // change to map for unique searchTerms [maybe?]
	// here we should create another map for unique bands [done!]
	for _, artist := range data.ArtistsList {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchTerm)) {
			suggestions = append(suggestions, artist.Name+"-band/artist")
			set[artist.Id] = artist.Name
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(searchTerm)) {
				suggestions = append(suggestions, member+"-member "+artist.Name)
				set[artist.Id] = artist.Name
			}
		}

		creationDate := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDate, searchTerm) {
			suggestions = append(suggestions, creationDate+"-creationDate")
			set[artist.Id] = artist.Name
		}

		if strings.Contains(artist.FirstAlbum, searchTerm) { // first album date e.g. "01-01-2012"
			suggestions = append(suggestions, artist.FirstAlbum+"-first album date")
			set[artist.Id] = artist.Name
		}

		for country, valMap := range artist.DatesLocation {
			if strings.Contains(strings.ToLower(country), strings.ToLower(searchTerm)) {
				suggestions = append(suggestions, country+"-country "+artist.Name)
				set[artist.Id] = artist.Name
			}
			for city, dates := range valMap {
				if strings.Contains(strings.ToLower(city), strings.ToLower(searchTerm)) {
					suggestions = append(suggestions, city+"-city "+artist.Name)
					set[artist.Id] = artist.Name
				}
				for _, date := range dates {
					if strings.Contains(date, strings.ToLower(searchTerm)) {
						suggestions = append(suggestions, date+"-date "+artist.Name)
						set[artist.Id] = artist.Name
					}
				}
			}
		}
	}
	return suggestions, set
}
