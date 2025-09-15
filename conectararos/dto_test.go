package conectararos_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/acdgbrasil/convsus"
	"github.com/acdgbrasil/convsus/conectararos"
	"github.com/stretchr/testify/require"
)

const PATH_REFERENCE_PERSON_DATA_FILE = "../Schemas/reference_person_data.json"

var plainPersonData []string

func TestMain(m *testing.M) {
	// Unmarshalling data
	data, err := os.ReadFile(PATH_REFERENCE_PERSON_DATA_FILE)
	if err != nil {
		log.Fatal("failed to read data from reference person data file")
	}
	var generic []map[string]interface{}
	if err := json.Unmarshal(data, &generic); err != nil {
		log.Fatal("failed to unmarshal data from reference person data file into generic array of data")
	}

	for _, rpd := range generic {
		plainPerson, err := json.Marshal(rpd)
		if err != nil {
			log.Fatal("failed to put data back in plain text form")
		}
		plainPersonData = append(plainPersonData, string(plainPerson))
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestReferencePersonFromJson_JsonData(t *testing.T) {
	var generic map[string]interface{}
	err := json.Unmarshal([]byte(plainPersonData[0]), &generic)
	require.NoError(t, err)

	// Testing correct json array
	for _, plainjson := range plainPersonData {
		_, err := conectararos.ReferencePersonFromJson(plainjson)
		require.NoError(t, err)
	}
	// Testing incorrect json array
	testCases := []struct {
		purpose     string
		plainjson   string
		expectedErr error
	}{
		{
			purpose:     "Should work",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: nil,
		},
		{
			purpose:     "Invalid JSON",
			plainjson:   `{"fullName": 1,"socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: convsus.ErrUnmarshalJson(nil),
		},
		{
			purpose:     "Missing field - fullName",
			plainjson:   `{"fullName": "","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - socialName",
			plainjson:   `{"fullName": "João da Silva","socialName": "","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - motherName",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - cpf",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - diagnosis",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - rg",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - rg - issueDate",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - rg - issuingBody",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - rg - number",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - rg - uf",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": ""},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - localLocalization",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - cep",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - address",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - neighborhood",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - addressNumber",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - city",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - phone",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - state",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - whoIsOpeningId",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "","birthDate": "1980-01-15","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - birthDate",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "","biologicalGender": "Masculino"}`,
			expectedErr: conectararos.ErrValidation(""),
		},
		{
			purpose:     "Missing field - biologicalGender",
			plainjson:   `{"fullName": "João da Silva","socialName": "João","motherName": "Maria da Silva","nis": "12345678901","cpf": "111.111.111-11","diagnosis": "Nenhum","rg": {"issueDate": "2010-05-20","issuingBody": "SSP","number": "1234567","uf": "SP"},"isShelter": false,"localLocalization": "URBAN","cep": "01001-000","adress": "Avenida Paulista","neighborhood": "Bela Vista","adressComplement": "Apto 101","adressNumber": "1000","city": "São Paulo","phone": "11999999999","state": "SP","whoIsOpeningId": "user1","birthDate": "1980-01-15","biologicalGender": ""}`,
			expectedErr: conectararos.ErrValidation(""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.purpose, func(t *testing.T) {
			_, err := conectararos.ReferencePersonFromJson(tc.plainjson)
			// Make verifications in case of throwing error
			if err != nil {
				require.NotNil(t, tc.expectedErr, "Did expect error in this test, but throwed none", err)
				// In case a error was expected. Validate if the expected error
				cErr, ok := err.(*convsus.Error)
				require.True(t, ok, "Returned error should be a formatted error by the system")
				require.True(t, cErr.Equals(err))
			} else {
				require.Nil(t, tc.expectedErr, "Did not expect error in this test, but throwed one", err)
			}
		})
	}
}
