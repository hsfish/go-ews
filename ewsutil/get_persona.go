package ewsutil

import (
	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

// FindPeople find persona slice by query string
func GetPersona(c ews.Client, personaID string) (*ewsxml.Persona, error) {

	resp, err := ews.GetPersona(c, &ews.GetPersonaRequest{
		PersonaId: ewsxml.PersonaId{Id: personaID},
	})

	if err != nil {
		return nil, err
	}

	return &resp.Persona, nil
}
