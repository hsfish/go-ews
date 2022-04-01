package ewsutil

import (
	"math"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

// FindPeople find persona slice by query string
func FindPeople(c ews.Client, q string) ([]ewsxml.Persona, error) {
	req := &ews.FindPeopleRequest{IndexedPageItemView: ews.IndexedPageItemView{
		MaxEntriesReturned: math.MaxInt32,
		Offset:             0,
		BasePoint:          ewsxml.BasePointBeginning,
	}, ParentFolderId: ews.ParentFolderId{
		DistinguishedFolderId: ewsxml.DistinguishedFolderId{Id: "directory"}},
		PersonaShape: &ews.PersonaShape{BaseShape: ewsxml.BaseShape_IdOnly,
			AdditionalProperties: ews.AdditionalProperties{
				FieldURI: []ews.FieldURI{
					{FieldURI: "persona:DisplayName"},
					{FieldURI: "persona:Title"},
					{FieldURI: "persona:EmailAddress"},
					{FieldURI: "persona:Departments"},
				},
			}},
		QueryString: q,
	}

	resp, err := ews.FindPeople(c, req)
	if err != nil {
		return nil, err
	}

	return resp.People.Persona, nil
}
