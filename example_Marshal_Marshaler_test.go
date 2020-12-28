package watson_test

import (
	"fmt"
	"time"

	"github.com/genkami/watson"
	"github.com/genkami/watson/pkg/types"
)

type Time time.Time

func (t Time) MarshalWatson() (*types.Value, error) {
	return types.NewObjectValue(map[string]*types.Value{
		"unix": types.NewIntValue(time.Time(t).Unix()),
	}), nil
}

func ExampleMarshal_marshaler() {
	now := Time(time.Now())
	buf, err := watson.Marshal(now)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}
