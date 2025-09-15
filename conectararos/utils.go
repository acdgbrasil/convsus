package conectararos

import (
	"fmt"
	"time"

	"github.com/acdgbrasil/convsus"
)

func UpdateTimeVariables(variablesToUpdate []string, generic map[string]interface{}) (map[string]interface{}, error) {
	// Updating value for variables time.Time before running default unmarshal json
	for _, v := range variablesToUpdate {
		// Skipping not populated field
		if generic[v] == nil {
			continue
		}
		vstring, ok := generic[v].(string)
		if !ok {
			return nil, convsus.ErrUnmarshalJson(convsus.ErrParsingData(fmt.Errorf("")))
		}
		var err error
		generic[v], err = time.Parse(time.DateOnly, vstring)
		if err != nil {
			return nil, convsus.ErrUnmarshalJson(err)
		}
	}
	return generic, nil
}
