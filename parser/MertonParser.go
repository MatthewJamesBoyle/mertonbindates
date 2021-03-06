package parser

import (
	"fmt"
	"gopkg.in/headzoo/surf.v1"
	"os"
	"regexp"
)

func Parse(c Config) map[string]string {

	bow := surf.NewBrowser()
	err := bow.Open(os.Getenv("URL"))
	handleError(err)

	fm, err := bow.Form("form.form-vertical")
	handleError(err)

	fm.Input("InputPremiseHouseNumber", os.Getenv("HOUSE_NUMBER"))
	fm.Input("InputPremisePostcode", os.Getenv("POSTCODE"))
	if fm.Submit() != nil {
		handleError(err)
	}

	//refresh form
	fm2, err := bow.Form("form.form-vertical")
	handleError(err)

	err2 := fm2.SelectByOptionValue("CRM_gdit_property_gdit_uprn", "48104419")
	handleError(err2)

	fm2.Submit()

	result := bow.Find("form.form-vertical")

	dates := getDates(result.Text())

	return map[string]string{
		"rubbish":   dates[0][0],
		"recycling": dates[1][0],
		"food":      dates[2][0],
	}
}

func getDates(s string) [][]string {
	re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
	return re.FindAllStringSubmatch(s, -1)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
