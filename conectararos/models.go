package conectararos

import (
	"encoding/json"
	"time"

	"github.com/acdgbrasil/convsus"
)

type rg struct {
	IssueDate   string `json:"issueDate"`
	IssuingBody string `json:"issuingBody"`
	Number      string `json:"number"`
	UF          string `json:"uf"`
}

// Validate for RG
func (r rg) Validate() error {
	if r.IssueDate == "" {
		return ErrValidation("issueDate is required")
	}
	if r.IssuingBody == "" {
		return ErrValidation("issuingBody is required")
	}
	if r.Number == "" {
		return ErrValidation("number is required")
	}
	if r.UF == "" {
		return ErrValidation("uf is required")
	}
	return nil
}

type observation struct {
	Observation      string    `json:"observation"`
	WhoIsObservingID string    `json:"whoIsObservingId"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
}

// Validate for Observation
func (o observation) Validate() error {
	if o.Observation == "" {
		return ErrValidation("observation is required")
	}
	if o.WhoIsObservingID == "" {
		return ErrValidation("whoIsObservingId is required")
	}
	return nil
}

type referencePerson struct {
	FullName      string        `json:"fullName"`
	SocialName    string        `json:"socialName"`
	MotherName    string        `json:"motherName"`
	NIS           string        `json:"nis,omitempty"`
	CPF           string        `json:"cpf"`
	Diagnosis     string        `json:"diagnosis"`
	RG            rg            `json:"rg"`
	IsShelter     bool          `json:"isShelter"`
	LocalLocation string        `json:"localLocalization"`
	CEP           string        `json:"cep,omitempty"`
	Address       string        `json:"adress"`
	Neighborhood  string        `json:"neighborhood"`
	AddressComp   string        `json:"adressComplement"`
	AddressNumber string        `json:"adressNumber"`
	City          string        `json:"city"`
	Phone         string        `json:"phone"`
	State         string        `json:"state"`
	WhoIsOpening  string        `json:"whoIsOpeningId"`
	BirthDate     time.Time     `json:"birthDate"`
	Gender        string        `json:"biologicalGender"`
	Observations  []observation `json:"observations,omitempty"`
	CreatedAt     time.Time     `json:"createdAt,omitempty"`
	UpdatedAt     time.Time     `json:"updatedAt,omitempty"`
}

type referencePersonJsonTarget referencePerson // This alias is need for the json unmarshal method to avoid recursion
var VARIABLES_TO_UPDATE = []string{"birthDate", "createdAt", "updatedAt"}

func (rp *referencePerson) UnmarshalJSON(data []byte) error {
	// Unmarshalling into generic value
	var generic map[string]interface{}
	if err := json.Unmarshal(data, &generic); err != nil {
		return convsus.ErrUnmarshalJson(err)
	}
	// Update time variables parsing from string to time.Time
	var err error
	if generic, err = UpdateTimeVariables(VARIABLES_TO_UPDATE, generic); err != nil {
		return convsus.ErrUnmarshalJson(err)
	}
	// Returning to bytes
	data, err = json.Marshal(generic)
	if err != nil {
		return convsus.ErrUnmarshalJson(err)
	}
	// Running default json call after adapting data
	var alias referencePersonJsonTarget
	if err := json.Unmarshal(data, &alias); err != nil {
		return convsus.ErrUnmarshalJson(err)
	}
	// Populate original struct with the data
	rp.copyFrom(alias)

	return nil
}

func (rp *referencePerson) copyFrom(alias referencePersonJsonTarget) {
	rp.FullName = alias.FullName
	rp.SocialName = alias.SocialName
	rp.MotherName = alias.MotherName
	rp.NIS = alias.NIS
	rp.CPF = alias.CPF
	rp.Diagnosis = alias.Diagnosis
	rp.RG = alias.RG
	rp.IsShelter = alias.IsShelter
	rp.LocalLocation = alias.LocalLocation
	rp.CEP = alias.CEP
	rp.Address = alias.Address
	rp.Neighborhood = alias.Neighborhood
	rp.AddressComp = alias.AddressComp
	rp.AddressNumber = alias.AddressNumber
	rp.City = alias.City
	rp.Phone = alias.Phone
	rp.State = alias.State
	rp.WhoIsOpening = alias.WhoIsOpening
	rp.BirthDate = alias.BirthDate
	rp.Gender = alias.Gender
	rp.Observations = alias.Observations
	rp.CreatedAt = alias.CreatedAt
	rp.UpdatedAt = alias.UpdatedAt
}

// Validate for ReferencePerson
func (rp *referencePerson) Validate() error {
	if rp.FullName == "" {
		return ErrValidation("fullName is required")
	}
	if rp.SocialName == "" {
		return ErrValidation("socialName is required")
	}
	if rp.MotherName == "" {
		return ErrValidation("motherName is required")
	}
	if rp.CPF == "" {
		return ErrValidation("cpf is required")
	}
	if rp.Diagnosis == "" {
		return ErrValidation("diagnosis is required")
	}
	if rp.LocalLocation == "" {
		return ErrValidation("localLocalization is required")
	}
	if rp.CEP == "" {
		return ErrValidation("cep is required")
	}
	if rp.Address == "" {
		return ErrValidation("adress is required")
	}
	if rp.Neighborhood == "" {
		return ErrValidation("neighborhood is required")
	}
	if rp.AddressNumber == "" {
		return ErrValidation("adressNumber is required")
	}
	if rp.City == "" {
		return ErrValidation("city is required")
	}
	if rp.Phone == "" {
		return ErrValidation("phone is required")
	}
	if rp.State == "" {
		return ErrValidation("state is required")
	}
	if rp.WhoIsOpening == "" {
		return ErrValidation("whoIsOpeningId is required")
	}
	if rp.BirthDate.IsZero() {
		return ErrValidation("birthDate is required")
	}
	if rp.Gender == "" {
		return ErrValidation("biologicalGender is required")
	}
	// Validate fields in RG
	if err := rp.RG.Validate(); err != nil {
		return err
	}
	// Validate all observations
	for _, o := range rp.Observations {
		if err := o.Validate(); err != nil {
			return err
		}
	}

	return nil
}
