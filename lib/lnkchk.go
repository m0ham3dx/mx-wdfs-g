/*
Link status checker
*/

package lib

import (
	f "fmt"
	n "net/http"
	s "strconv"
)

func LnkChk() {

	links := []string{
		"https://snips.sh/f/uEf7aLoioN",
	}

	for _, link := range links {
		LnkChkStatus(link)
	}

}

func LnkChkStatus(link string) {
	response, err := n.Get(link)

	if err != nil {
		Err("===ERROR===")
		f.Println(err)
		ChBad("❌", link, "")
		return
	}
	rez := s.Itoa(int(response.StatusCode))
	ChGud("✅", link, rez)
}
