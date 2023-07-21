package namecheap

import (
	"encoding/xml"
	"fmt"
)

type DomainsCreateResponse struct {
	XMLName *xml.Name `xml:"ApiResponse"`
	Errors  *[]struct {
		Message *string `xml:",chardata"`
		Number  *string `xml:"Number,attr"`
	} `xml:"Errors>Error"`
	CommandResponse *DomainsCreateCommandResponse `xml:"CommandResponse"`
}

type DomainsCreateCommandResponse struct {
	Result *DomainsCreateResult `xml:"DomainCreateResult"`
}

type DomainsCreateResult struct {
	Domain        *string `xml:"Domain,attr"`
	Registered    *bool   `xml:"Registered,attr"`
	ChargedAmount *string `xml:"ChargedAmount,attr"`
}

func (ds *DomainsService) DomainsCreate(domainName string) (*DomainsCreateCommandResponse, error) {

	var response DomainsCreateResponse

	params := map[string]string{
		"Command":                 "namecheap.domains.create",
		"DomainName":              domainName,
		"Years":                   "2",
		"RegistrantFirstName":     "2",
		"RegistrantLastName":      "2",
		"RegistrantAddress1":      "2",
		"RegistrantCity":          "2",
		"RegistrantStateProvince": "2",
		"RegistrantPostalCode":    "2",
		"RegistrantCountry":       "2",
		"RegistrantPhone":         "2",
		"RegistrantEmailAddress":  "2",
		"TechFirstName":           "2",
		"TechLastName":            "2",
		"TechAddress1":            "2",
		"TechCity":                "2",
		"TechStateProvince":       "2",
		"TechPostalCode":          "2",
		"TechCountry":             "2",
		"TechPhone":               "2",
		"TechEmailAddress":        "2",
		"AdminFirstName":          "2",
		"AdminLastName":           "2",
		"AdminAddress1":           "2",
		"AdminCity":               "2",
		"AdminStateProvince":      "2",
		"AdminPostalCode":         "2",
		"AdminCountry":            "2",
		"AdminPhone":              "2",
		"AdminEmailAddress":       "2",
		"AuxBillingFirstName":     "2",
		"AuxBillingAddress1":      "2",
		"AuxBillingCity":          "2",
		"AuxBillingStateProvince": "2",
		"AuxBillingPostalCode":    "2",
		"AuxBillingCountry":       "2",
		"AuxBillingPhone":         "2",
		"AuxBillingEmailAddress":  "2",
		"Extended attributes":     "2",
		"Nameservers":             "2",
	}
	_, err := ds.client.DoXML(params, &response)

	if err != nil {
		return nil, err
	}

	if response.Errors != nil && len(*response.Errors) > 0 {
		apiErr := (*response.Errors)[0]
		return nil, fmt.Errorf("%s (%s)", *apiErr.Message, *apiErr.Number)
	}

	return response.CommandResponse, nil

}
