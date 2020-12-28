package watson_test

import (
	"fmt"

	"github.com/genkami/watson"
)

type Music struct {
	Track int    `watson:"track"`
	Title string `watson:"title"`
}

const marshaledMusic = `
~?
SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkShaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaarrkzg$
BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SShaaarrkShaaaaaarrk-SShkShaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaaarrk-g
`

func ExampleUnmarshal() {
	var music Music
	err := watson.Unmarshal([]byte(marshaledMusic), &music)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Track: %d, Title: %s\n", music.Track, music.Title)
}