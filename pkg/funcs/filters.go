package funcs

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Filter(members []string, firstmin string, firstmax string, creatmin string, creatmax string, location string) {
	var found []Artist
	for i, v := range All.Artists {
		if CreationDateCheck(creatmax, creatmin, v.CreationDate) && FirstAlbumCheck(firstmax, firstmin, v.FirstAlbum) && LocationCheck(location, v.Rel) && MembersCheck(members, len(v.Members)) {
			found = append(found, All.Artists[i])
		}
	}
	FoundArtists = found
}

func CreationDateCheck(creatmin string, creatmax string, creatDate int) bool {
	return creatDate <= ToNum(creatmin) && creatDate >= ToNum(creatmax)
}

func FirstAlbumCheck(firstmin string, firstmax string, firstdate string) bool {
	year := firstdate[6:]
	return ToNum(year) <= ToNum(firstmin) && ToNum(year) >= ToNum(firstmax)
}

func LocationCheck(location string, rel map[string][]string) bool {
	_, ok := rel[strings.ToLower(location)]
	if ok {
		return true
	}
	if !ok {
		for j := range rel {
			j = strings.ReplaceAll(strings.ReplaceAll(j, "_", " "), "-", " ")
			if strings.Contains(strings.ToLower(j), strings.ToLower(location)) {
				return true
			}
		}
	}
	return false
}

func MembersCheck(nmembers []string, num int) bool {
	if nmembers == nil {
		return true
	}
	for _, v := range nmembers {
		if strconv.Itoa(num) == v {
			return true
		}
	}
	return false
}

func ToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return num
}
