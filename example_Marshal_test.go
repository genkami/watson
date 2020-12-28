package watson_test

import (
	"fmt"

	"github.com/genkami/watson"
)

type Person struct {
	FullName string `watson:"fullName"`
	Nickname string `watson:"nickname,omitempty"`
}

func ExampleMarshal() {
	user := Person{
		FullName: "Motoaki Tanigo",
		Nickname: "YAGOO",
	}
	buf, err := watson.Marshal(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}
