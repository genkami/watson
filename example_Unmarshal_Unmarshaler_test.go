package watson_test

import (
	"errors"
	"fmt"
	"strings"

	"github.com/genkami/watson"
	"github.com/genkami/watson/pkg/types"
)

type Email struct {
	Local  string
	Domain string
}

func (e *Email) UnmarshalWatson(v *types.Value) error {
	if v.Kind != types.String {
		return errors.New("expected string")
	}
	parts := strings.Split(string(v.String), "@")
	if len(parts) != 2 {
		return errors.New("value must be like 'local@domain.example.com'")
	}
	e.Local = parts[0]
	e.Domain = parts[1]
	return nil
}

const marshaledEmail = `
?SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-
`

func ExampleUnmarshal_unmarshaler() {
	var email Email
	err := watson.Unmarshal([]byte(marshaledEmail), &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Local: %s, Domain: %s\n", email.Local, email.Domain)
}
