package conectararos

import (
	"encoding/json"

	"github.com/acdgbrasil/convsus"
)

func ReferencePersonFromJson(plainjson string) (*referencePerson, error) {
	// Unmarshal the data
	var person referencePerson
	if err := json.Unmarshal([]byte(plainjson), &person); err != nil {
		return nil, convsus.ErrUnmarshalJson(err)
	}
	// Verify the unmarshalled data
	if err := person.Validate(); err != nil {
		return nil, err
	}
	return &person, nil
}
