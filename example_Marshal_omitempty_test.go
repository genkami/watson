package watson_test

import (
	"fmt"

	"github.com/genkami/watson"
)

type Profile struct {
	Name             string `watson:"name"`
	Age              int    `watson:"age,omitempty"`
	UnsecurePassword string `watson:"-"`
}

func ExampleMarshal_omitempty() {
	profile := Profile{
		Name:             "Calliope Mori",
		UnsecurePassword: "jaoijgeoaivj#&*RJI",
	}
	buf, err := watson.Marshal(&profile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}
