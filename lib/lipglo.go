/*
Fri Jul 28 02:39:03 UTC 2023
This for the colors
*/

package lib

import (
	f "fmt"
	t "time"

	lg "github.com/charmbracelet/lipgloss"
)

// Constant definitions - The text to be displayed

const (
	bannerText = `
https://github.com/m0ham3dx/mx-wdfs-g

╔╦╗ ╔═╗ ╦ ╦ ╔═╗ ╔╦╗ ╔═╗ ╔╦╗  ╦ ╦     
║║║ ║═║ ╠═╣ ╠═╣ ║║║  ╠║  ║║ ╔╩╦╝     
╩ ╩ ╚═╝ ╩ ╩ ╩ ╩ ╩ ╩ ╚═╝ ═╩╝ ╩ ╩      

╦ ╦ ╔═╗ ╔╗  ╔╦╗ ╔═╗ ╦  ╦   ╔═╗ ╔═╗ ╦═╗ ╦ ╔═╗ ╔╦╗   ╔╦╗ ╔═╗ ╦ ╦ ╔╗╔ ╦   ╔═╗ ╔═╗ ╔╦╗ ╔═╗ ╦═╗
║║║ ║╣  ╠╩╗  ║║ ║╣  ╚╗╔╝   ╚═╗ ║   ╠╦╝ ║ ╠═╝  ║     ║║ ║ ║ ║║║ ║║║ ║   ║ ║ ╠═╣  ║║ ║╣  ╠╦╝
╚╩╝ ╚═╝ ╚═╝ ═╩╝ ╚═╝  ╚╝    ╚═╝ ╚═╝ ╩╚═ ╩ ╩    ╩    ═╩╝ ╚═╝ ╚╩╝ ╝╚╝ ╩═╝ ╚═╝ ╩ ╩ ═╩╝ ╚═╝ ╩╚═
`

	paraText = `
This script will download - https://snips.sh/f/uEf7aLoioN -  and rename it as 🍌.fish file

VirusTotal Scan = https://www.virustotal.com/gui/url/a92d97ab321c6d91b0754d2c8e0c33393279bdd0836ed644abb079283ef172f0
`
)

// Definiing the Lipgloss styles

var (

	//Banner Stylea
	banner = lg.NewStyle().
		Foreground(lg.Color("#FF5F9E")).
		Background(lg.Color("#2E0249")).
		Width(100).Align(lg.Center).Bold(true).
		Border(lg.DoubleBorder(), true, true, true, true).
		PaddingTop(1).
		PaddingBottom(1).
		BorderForeground(lg.Color("#FC2947"))

	//ParaText Style
	paratext = lg.NewStyle().
			Foreground(lg.Color("#00DFA2")).
			Background(lg.Color("#04293A")).
			Width(100).Align(lg.Center).Bold(true)

	//Reachable Style
	chstyleGood = lg.NewStyle().
			Background(lg.Color("#02383C")).
			Foreground(lg.Color("#00DFA2")).
			Align(lg.Center).
			Width(100)

	//Unreachable Style
	chstyleBad = lg.NewStyle().
			Background(lg.Color("#570530")).
			Foreground(lg.Color("#FF008E")).
			Align(lg.Center).
			Width(100)

	errorBar = lg.NewStyle().
			Foreground(lg.Color("#FF008E")).
			Background(lg.Color("#570530")).
			Align(lg.Center).
			Width(100)
)

/*
+++++++++++++++++++++++++ Functions +++++++++++++++++++++++++
*/

// BannerFunction

func BannerText() {
	t := t.Now().UTC().Format("Mon Jan 02 15:04:05 UTC")
	f.Println(banner.Render(t, "\n", bannerText))
}

// ParaTextFunction

func ParaText() {
	f.Println(paratext.Render(paraText))
}

func ParaText2(s string) {
	f.Println(paratext.Render(s))
}

// Link Checker Functions
func ChGud(s string, t string, u string) {
	f.Println(chstyleGood.Render(s, t, u))
}

func ChBad(s string, t string, u string) {
	f.Println(chstyleBad.Render(s, t, u))
}

// Error Bar Function
func Err(s string) {
	f.Println(errorBar.Render(s))
}
