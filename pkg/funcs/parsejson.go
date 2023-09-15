package funcs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// main struct
var All API

func ParseJson() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations := "https://groupietrackers.herokuapp.com/api/relation"
	ParseInfo(urlArtists, &All.Artists)
	ParseInfo(urlRelations, &All.Relation)
	for i, v := range All.Relation.Index {
		All.Artists[i].Rel = v.DatesLocations
	}
	preview := []string{
		"https://open.spotify.com/embed/track/3z8h0TU7ReDPLIbEnYhWZb?utm_source=generator",
		"https://open.spotify.com/embed/track/1HnriuDThLq7bEl1QKiaJL?utm_source=generator",
		"https://open.spotify.com/embed/track/7rPzEczIS574IgPaiPieS3?utm_source=generator",
		"https://open.spotify.com/embed/track/1PbL8RWQlpPRIVAk8by05R?utm_source=generator",
		"https://open.spotify.com/embed/track/7AFASza1mXqntmGtbxXprO?utm_source=generator",
		"https://open.spotify.com/embed/track/5bJ1DrEM4hNCafcDd1oxHx?utm_source=generator",
		"https://open.spotify.com/embed/track/4h7qcXBtaOJnmrapxoWxGf?utm_source=generator",
		"https://open.spotify.com/embed/track/7KXjTSCq5nL1LoYtL7XAwS?utm_source=generator",
		"https://open.spotify.com/embed/track/08mG3Y1vljYA6bvDt4Wqkj?utm_source=generator",
		"https://open.spotify.com/embed/track/6QewNVIDKdSl8Y3ycuHIei?utm_source=generator",
		"https://open.spotify.com/embed/track/1pXrR5Y9OgcIV2JEAl2lCB?utm_source=generator",
		"https://open.spotify.com/embed/track/6O20JhBJPePEkBdrB5sqRx?utm_source=generator",
		"https://open.spotify.com/embed/track/480dDrqG7LO6qDaphHeXlM?utm_source=generator",
		"https://open.spotify.com/embed/track/1NCuYqMc8hKMb4cpNTcJbD?utm_source=generator",
		"https://open.spotify.com/embed/track/78lgmZwycJ3nzsdgmPPGNx?utm_source=generator",
		"https://open.spotify.com/embed/track/2aoo2jlRnM3A0NyLQqMN2f?utm_source=generator",
		"https://open.spotify.com/embed/track/4UDmDIqJIbrW0hMBQMFOsM?utm_source=generator",
		"https://open.spotify.com/embed/track/0XHWClSz0v6RIaRSGqJH3X?utm_source=generator",
		"https://open.spotify.com/embed/track/1xsYj84j7hUDDnTTerGWlH?utm_source=generator",
		"https://open.spotify.com/embed/track/37Tmv4NnfQeb0ZgUC4fOJj?utm_source=generator",
		"https://open.spotify.com/embed/track/2iJpjciYl8vfQbb543b5Pb?utm_source=generator",
		"https://open.spotify.com/embed/track/1PYG9Akj0LAZZUDXzV9m1S?utm_source=generator",
		"https://open.spotify.com/embed/track/62yJjFtgkhUrXktIoSjgP2?utm_source=generator",
		"https://open.spotify.com/embed/track/4VXIryQMWpIdGgYR4TrjT1?utm_source=generator",
		"https://open.spotify.com/embed/track/3e7Y6sfFlIdBMJhX7wpqVO?utm_source=generator",
		"https://open.spotify.com/embed/track/2qxmye6gAegTMjLKEBoR3d?utm_source=generator",
		"https://open.spotify.com/embed/track/5YbgcwHjQhdT1BYQ4rxWlD?utm_source=generator",
		"https://open.spotify.com/embed/track/4HYUdQyUNZL1FMlPjRIGs5?utm_source=generator",
		"https://open.spotify.com/embed/track/0e7ipj03S05BNilyu5bRzt?utm_source=generator",
		"https://open.spotify.com/embed/track/6gBFPUFcJLzWGx4lenP6h2?utm_source=generator",
		"https://open.spotify.com/embed/track/4tqcoej1zPvwePZCzuAjJd?utm_source=generator",
		"https://open.spotify.com/embed/track/4QJLKU75Rg4558f4LbDBRi?utm_source=generator",
		"https://open.spotify.com/embed/track/7N1Vjtzr1lmmCW9iasQ8YO?utm_source=generator",
		"https://open.spotify.com/embed/track/4JiEyzf0Md7KEFFGWDDdCr?utm_source=generator",
		"https://open.spotify.com/embed/track/6KIKRz9eSTXdNsGUnomdtW?utm_source=generator",
		"https://open.spotify.com/embed/track/0gzqZ9d1jIKo9psEIthwXe?utm_source=generator",
		"https://open.spotify.com/embed/track/0BxE4FqsDD1Ot4YuBXwAPp?utm_source=generator",
		"https://open.spotify.com/embed/track/04aAxqtGp5pv12UXAg4pkq?utm_source=generator",
		"https://open.spotify.com/embed/track/0d28khcov6AiegSCpG5TuT?utm_source=generator",
		"https://open.spotify.com/embed/track/40riOy7x9W7GXjyGp4pjAv?utm_source=generator",
		"https://open.spotify.com/embed/track/60a0Rd6pjrkxjPbaKzXjfq?utm_source=generator",
		"https://open.spotify.com/embed/track/3ZOEytgrvLwQaqXreDs2Jx?utm_source=generator",
		"https://open.spotify.com/embed/track/7lQ8MOhq6IN2w8EYcFNSUk?utm_source=generator",
		"https://open.spotify.com/embed/track/64yrDBpcdwEdNY9loyEGbX?utm_source=generator",
		"https://open.spotify.com/embed/track/2MuWTIM3b0YEAskbeeFE1i?utm_source=generator",
		"https://open.spotify.com/embed/track/3RiPr603aXAoi4GHyXx0uy?utm_source=generator",
		"https://open.spotify.com/embed/track/6ECp64rv50XVz93WvxXMGF?utm_source=generator",
		"https://open.spotify.com/embed/track/3CRDbSIZ4r5MsZ0YwxuEkn?utm_source=generator",
		"https://open.spotify.com/embed/track/63T7DJ1AFDD6Bn8VzG6JE8?utm_source=generator",
		"https://open.spotify.com/embed/track/3lPr8ghNDBLc2uZovNyLs9?utm_source=generator",
		"https://open.spotify.com/embed/track/5UWwZ5lm5PKu6eKsHAGxOk?utm_source=generator",
		"https://open.spotify.com/embed/track/1i1fxkWeaMmKEB4T7zqbzK?utm_source=generator",
	}

	for i, v := range preview {
		All.Artists[i].Preview = v
	}
}

func ParseInfo(url string, temp interface{}) {
	res, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		return
	}
	text, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(text, &temp)
}
