package types_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/types"
)

func TestBindConvertsInt(t *testing.T) {
	var err error
	var got int
	var val = types.NewIntValue(123)
	var want int = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt8(t *testing.T) {
	var err error
	var got int8
	var val = types.NewIntValue(123)
	var want int8 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt16(t *testing.T) {
	var err error
	var got int16
	var val = types.NewIntValue(123)
	var want int16 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt32(t *testing.T) {
	var err error
	var got int32
	var val = types.NewIntValue(123)
	var want int32 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt64(t *testing.T) {
	var err error
	var got int64
	var val = types.NewIntValue(123)
	var want int64 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint(t *testing.T) {
	var err error
	var got uint
	var val = types.NewUintValue(123)
	var want uint = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint8(t *testing.T) {
	var err error
	var got uint8
	var val = types.NewUintValue(123)
	var want uint8 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint16(t *testing.T) {
	var err error
	var got uint16
	var val = types.NewUintValue(123)
	var want uint16 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint32(t *testing.T) {
	var err error
	var got uint32
	var val = types.NewUintValue(123)
	var want uint32 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint64(t *testing.T) {
	var err error
	var got uint64
	var val = types.NewUintValue(123)
	var want uint64 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsFloat32(t *testing.T) {
	var err error
	var got float32
	var val = types.NewFloatValue(1.23e4)
	var want float32 = 1.23e4
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	if !closeEnough(float64(want), float64(got)) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestBindConvertsFloat64(t *testing.T) {
	var err error
	var got float64
	var val = types.NewFloatValue(1.23e45)
	var want float64 = 1.23e45
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsString(t *testing.T) {
	var err error
	var got string
	var val = types.NewStringValue([]byte("hoge"))
	var want string = "hoge"
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsBool(t *testing.T) {
	var err error
	var got bool
	var val = types.NewBoolValue(true)
	var want bool = true
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilInterface(t *testing.T) {
	var err error
	var got interface{}
	var val = types.NewNilValue()
	var want interface{} = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilPtr(t *testing.T) {
	var err error
	var got *int
	var val = types.NewNilValue()
	var want *int = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilSlice(t *testing.T) {
	var err error
	var got []int
	var val = types.NewNilValue()
	var want []int = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewNilValue()
	var want map[string]int = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
	})
	var want map[string]int = map[string]int{
		"hoge": 123,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsHeteroMap(t *testing.T) {
	var err error
	var got map[string]interface{}
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
		"fuga": types.NewStringValue([]byte("foo")),
		"bar": types.NewObjectValue(map[string]*types.Value{
			"baz": types.NewStringValue([]byte("quux")),
		}),
	})
	var want map[string]interface{} = map[string]interface{}{
		"hoge": int64(123),
		"fuga": "foo",
		"bar": map[string]interface{}{
			"baz": "quux",
		},
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsSlice(t *testing.T) {
	var err error
	var got []int
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewIntValue(456),
	})
	var want []int = []int{123, 456}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsHeteroSlice(t *testing.T) {
	var err error
	var got []interface{}
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("456")),
	})
	var want []interface{} = []interface{}{int64(123), "456"}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsArray(t *testing.T) {
	var err error
	var got [2]int
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewIntValue(456),
	})
	var want [2]int = [2]int{123, 456}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsHeteroArray(t *testing.T) {
	var err error
	var got [2]interface{}
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("456")),
	})
	var want [2]interface{} = [2]interface{}{int64(123), "456"}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsPtr(t *testing.T) {
	var err error
	var got map[string]*int
	var v int = 123
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
	})
	var want map[string]*int = map[string]*int{
		"hoge": &v,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUntaggedStruct(t *testing.T) {
	var err error
	var got untagged
	var val = types.NewObjectValue(map[string]*types.Value{
		"name":     types.NewStringValue([]byte("hoge")),
		"longname": types.NewStringValue([]byte("longhoge")),
	})
	var want untagged = untagged{
		Name:     "hoge",
		LongName: "longhoge",
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNestedStruct(t *testing.T) {
	var err error
	var got nested
	var val = types.NewObjectValue(map[string]*types.Value{
		"value": types.NewObjectValue(map[string]*types.Value{
			"value": types.NewIntValue(123),
		}),
	})
	var want nested = nested{
		Value: &nestedInner{
			Value: 123,
		},
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsEmbeddedStruct(t *testing.T) {
	var err error
	var got embedded
	var val = types.NewObjectValue(map[string]*types.Value{
		"field": types.NewIntValue(123),
		"embeddedinner": types.NewObjectValue(map[string]*types.Value{
			"anotherfield": types.NewIntValue(456),
		}),
	})
	var want embedded = embedded{
		Field: 123,
		EmbeddedInner: EmbeddedInner{
			AnotherField: 456,
		},
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsTaggedStruct(t *testing.T) {
	var err error
	var got tagged
	var val = types.NewObjectValue(map[string]*types.Value{
		"customName": types.NewIntValue(123),
	})
	var want tagged = tagged{
		Field: 123,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindOmitsPrivateField(t *testing.T) {
	var err error
	var got private
	var val = types.NewObjectValue(map[string]*types.Value{
		"publicfield":  types.NewIntValue(123),
		"privatefield": types.NewIntValue(456),
	})
	var want private = private{
		PublicField: 123,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if want.PublicField != got.PublicField || want.privateField != got.privateField {
		t.Errorf("expected %#v but got %#v", &want, &got)
	}
}

func TestBindConvertsFieldTaggedWithHyphen(t *testing.T) {
	var err error
	var got alwaysomit
	var val = types.NewObjectValue(map[string]*types.Value{
		"shouldBeIncluded": types.NewIntValue(123),
		"shouldbeomitted":  types.NewIntValue(456),
	})
	var want alwaysomit = alwaysomit{
		ShouldBeIncluded: 123,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindEmbedsFieldTaggedWithInline(t *testing.T) {
	var err error
	var got inline
	var val = types.NewObjectValue(map[string]*types.Value{
		"field":       types.NewIntValue(123),
		"nestedfield": types.NewIntValue(456),
	})
	var want inline = inline{
		Field: 123,
		Inner: inlineInner{
			NestedField: 456,
		},
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUnmarshaler(t *testing.T) {
	var err error
	var got customUnmarshaler
	var val = types.NewObjectValue(map[string]*types.Value{
		"customKey": types.NewIntValue(123),
	})
	var want customUnmarshaler = customUnmarshaler{
		SomeField: 123,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUnmarshalerEvenIfNotStruct(t *testing.T) {
	var err error
	var got primitiveUnmarshaler
	var val = types.NewObjectValue(map[string]*types.Value{
		"customKey": types.NewIntValue(123),
	})
	var want primitiveUnmarshaler = primitiveUnmarshaler(123)
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNestedUnmarshaler(t *testing.T) {
	var err error
	var got customUnmarshalerOuter
	var val = types.NewObjectValue(map[string]*types.Value{
		"unmarshaler": types.NewObjectValue(map[string]*types.Value{
			"customKey": types.NewIntValue(123),
		}),
	})
	var want customUnmarshalerOuter = customUnmarshalerOuter{
		Unmarshaler: &customUnmarshaler{
			SomeField: 123,
		},
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindReturnsErrorWhenTypeMismatchInTopLevel(t *testing.T) {
	var err error
	var val = types.NewIntValue(123)
	var bindTo string
	err = val.Bind(&bindTo)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	pat := regexp.MustCompile(`\(at <root>\)`)
	if !pat.MatchString(err.Error()) {
		t.Errorf("expected \"%s\" to match /%s/, but it didn't", err.Error(), pat.String())
	}
}

func TestBindReturnsErrorWhenTypeMismatchInMapElem(t *testing.T) {
	var err error
	var val = types.NewObjectValue(map[string]*types.Value{
		"validValue": types.NewStringValue([]byte("123")),
		"value":      types.NewIntValue(456),
	})
	var bindTo map[string]string
	err = val.Bind(&bindTo)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	pat := regexp.MustCompile(`\(at <root>.value\)`)
	if !pat.MatchString(err.Error()) {
		t.Errorf("expected \"%s\" to match /%s/, but it didn't", err.Error(), pat.String())
	}
}

func TestBindReturnsErrorWhenTypeMismatchInArrayElem(t *testing.T) {
	var err error
	var val = types.NewArrayValue([]*types.Value{
		types.NewStringValue([]byte("123")),
		types.NewIntValue(456),
	})
	var bindTo []string
	err = val.Bind(&bindTo)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	pat := regexp.MustCompile(`\(at <root>\[1\]\)`)
	if !pat.MatchString(err.Error()) {
		t.Errorf("expected \"%s\" to match /%s/, but it didn't", err.Error(), pat.String())
	}
}

func TestBindByReflectionConvertsInt(t *testing.T) {
	var err error
	var got int
	var val = types.NewIntValue(123)
	var want int = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt8(t *testing.T) {
	var err error
	var got int8
	var val = types.NewIntValue(123)
	var want int8 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt16(t *testing.T) {
	var err error
	var got int16
	var val = types.NewIntValue(123)
	var want int16 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt32(t *testing.T) {
	var err error
	var got int32
	var val = types.NewIntValue(123)
	var want int32 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt64(t *testing.T) {
	var err error
	var got int64
	var val = types.NewIntValue(123)
	var want int64 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint(t *testing.T) {
	var err error
	var got uint
	var val = types.NewUintValue(123)
	var want uint = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint8(t *testing.T) {
	var err error
	var got uint8
	var val = types.NewUintValue(123)
	var want uint8 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint16(t *testing.T) {
	var err error
	var got uint16
	var val = types.NewUintValue(123)
	var want uint16 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint32(t *testing.T) {
	var err error
	var got uint32
	var val = types.NewUintValue(123)
	var want uint32 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint64(t *testing.T) {
	var err error
	var got uint64
	var val = types.NewUintValue(123)
	var want uint64 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsFloat32(t *testing.T) {
	var err error
	var got float32
	var val = types.NewFloatValue(1.23e4)
	var want float32 = 1.23e4
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	if !closeEnough(float64(want), float64(got)) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestBindByReflectionConvertsFloat64(t *testing.T) {
	var err error
	var got float64
	var val = types.NewFloatValue(1.23e45)
	var want float64 = 1.23e45
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsString(t *testing.T) {
	var err error
	var got string
	var val = types.NewStringValue([]byte("hoge"))
	var want string = "hoge"
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsBool(t *testing.T) {
	var err error
	var got bool
	var val = types.NewBoolValue(true)
	var want bool = true
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilInterface(t *testing.T) {
	var err error
	var got interface{}
	var val = types.NewNilValue()
	var want interface{} = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilPtr(t *testing.T) {
	var err error
	var got *int
	var val = types.NewNilValue()
	var want *int = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilSlice(t *testing.T) {
	var err error
	var got []int
	var val = types.NewNilValue()
	var want []int = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewNilValue()
	var want map[string]int = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
	})
	var want map[string]int = map[string]int{
		"hoge": 123,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsHeteroMap(t *testing.T) {
	var err error
	var got map[string]interface{}
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
		"fuga": types.NewStringValue([]byte("foo")),
		"bar": types.NewObjectValue(map[string]*types.Value{
			"baz": types.NewStringValue([]byte("quux")),
		}),
	})
	var want map[string]interface{} = map[string]interface{}{
		"hoge": int64(123),
		"fuga": "foo",
		"bar": map[string]interface{}{
			"baz": "quux",
		},
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsSlice(t *testing.T) {
	var err error
	var got []int
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewIntValue(456),
	})
	var want []int = []int{123, 456}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsHeteroSlice(t *testing.T) {
	var err error
	var got []interface{}
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("456")),
	})
	var want []interface{} = []interface{}{int64(123), "456"}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsArray(t *testing.T) {
	var err error
	var got [2]int
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewIntValue(456),
	})
	var want [2]int = [2]int{123, 456}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsHeteroArray(t *testing.T) {
	var err error
	var got [2]interface{}
	var val = types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("456")),
	})
	var want [2]interface{} = [2]interface{}{int64(123), "456"}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsPtr(t *testing.T) {
	var err error
	var got map[string]*int
	var v int = 123
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
	})
	var want map[string]*int = map[string]*int{
		"hoge": &v,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUntaggedStruct(t *testing.T) {
	var err error
	var got untagged
	var val = types.NewObjectValue(map[string]*types.Value{
		"name":     types.NewStringValue([]byte("hoge")),
		"longname": types.NewStringValue([]byte("longhoge")),
	})
	var want untagged = untagged{
		Name:     "hoge",
		LongName: "longhoge",
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNestedStruct(t *testing.T) {
	var err error
	var got nested
	var val = types.NewObjectValue(map[string]*types.Value{
		"value": types.NewObjectValue(map[string]*types.Value{
			"value": types.NewIntValue(123),
		}),
	})
	var want nested = nested{
		Value: &nestedInner{
			Value: 123,
		},
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsEmbeddedStruct(t *testing.T) {
	var err error
	var got embedded
	var val = types.NewObjectValue(map[string]*types.Value{
		"field": types.NewIntValue(123),
		"embeddedinner": types.NewObjectValue(map[string]*types.Value{
			"anotherfield": types.NewIntValue(456),
		}),
	})
	var want embedded = embedded{
		Field: 123,
		EmbeddedInner: EmbeddedInner{
			AnotherField: 456,
		},
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsTaggedStruct(t *testing.T) {
	var err error
	var got tagged
	var val = types.NewObjectValue(map[string]*types.Value{
		"customName": types.NewIntValue(123),
	})
	var want tagged = tagged{
		Field: 123,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionOmitsPrivateField(t *testing.T) {
	var err error
	var got private
	var val = types.NewObjectValue(map[string]*types.Value{
		"publicfield":  types.NewIntValue(123),
		"privatefield": types.NewIntValue(456),
	})
	var want private = private{
		PublicField: 123,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if want.PublicField != got.PublicField || want.privateField != got.privateField {
		t.Errorf("expected %#v but got %#v", &want, &got)
	}
}

func TestBindByReflectionConvertsFieldTaggedWithHyphen(t *testing.T) {
	var err error
	var got alwaysomit
	var val = types.NewObjectValue(map[string]*types.Value{
		"shouldBeIncluded": types.NewIntValue(123),
		"shouldbeomitted":  types.NewIntValue(456),
	})
	var want alwaysomit = alwaysomit{
		ShouldBeIncluded: 123,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionEmbedsFieldTaggedWithInline(t *testing.T) {
	var err error
	var got inline
	var val = types.NewObjectValue(map[string]*types.Value{
		"field":       types.NewIntValue(123),
		"nestedfield": types.NewIntValue(456),
	})
	var want inline = inline{
		Field: 123,
		Inner: inlineInner{
			NestedField: 456,
		},
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUnmarshaler(t *testing.T) {
	var err error
	var got customUnmarshaler
	var val = types.NewObjectValue(map[string]*types.Value{
		"customKey": types.NewIntValue(123),
	})
	var want customUnmarshaler = customUnmarshaler{
		SomeField: 123,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUnmarshalerEvenIfNotStruct(t *testing.T) {
	var err error
	var got primitiveUnmarshaler
	var val = types.NewObjectValue(map[string]*types.Value{
		"customKey": types.NewIntValue(123),
	})
	var want primitiveUnmarshaler = primitiveUnmarshaler(123)
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNestedUnmarshaler(t *testing.T) {
	var err error
	var got customUnmarshalerOuter
	var val = types.NewObjectValue(map[string]*types.Value{
		"unmarshaler": types.NewObjectValue(map[string]*types.Value{
			"customKey": types.NewIntValue(123),
		}),
	})
	var want customUnmarshalerOuter = customUnmarshalerOuter{
		Unmarshaler: &customUnmarshaler{
			SomeField: 123,
		},
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionReturnsErrorWhenTypeMismatchInTopLevel(t *testing.T) {
	var err error
	var val = types.NewIntValue(123)
	var bindTo string
	err = val.BindByReflection(reflect.ValueOf(&bindTo))
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	pat := regexp.MustCompile(`\(at <root>\)`)
	if !pat.MatchString(err.Error()) {
		t.Errorf("expected \"%s\" to match /%s/, but it didn't", err.Error(), pat.String())
	}
}

func TestBindByReflectionReturnsErrorWhenTypeMismatchInMapElem(t *testing.T) {
	var err error
	var val = types.NewObjectValue(map[string]*types.Value{
		"validValue": types.NewStringValue([]byte("123")),
		"value":      types.NewIntValue(456),
	})
	var bindTo map[string]string
	err = val.BindByReflection(reflect.ValueOf(&bindTo))
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	pat := regexp.MustCompile(`\(at <root>.value\)`)
	if !pat.MatchString(err.Error()) {
		t.Errorf("expected \"%s\" to match /%s/, but it didn't", err.Error(), pat.String())
	}
}

func TestBindByReflectionReturnsErrorWhenTypeMismatchInArrayElem(t *testing.T) {
	var err error
	var val = types.NewArrayValue([]*types.Value{
		types.NewStringValue([]byte("123")),
		types.NewIntValue(456),
	})
	var bindTo []string
	err = val.BindByReflection(reflect.ValueOf(&bindTo))
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	pat := regexp.MustCompile(`\(at <root>\[1\]\)`)
	if !pat.MatchString(err.Error()) {
		t.Errorf("expected \"%s\" to match /%s/, but it didn't", err.Error(), pat.String())
	}
}
