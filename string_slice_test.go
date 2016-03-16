package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type SliceStringCaster struct{}

func (c SliceStringCaster) StringSlice() []string {
	return []string{"a", "b", "c"}
}

type NotSliceStringCaster struct{}

func TestStringSliceCast(t *testing.T) {
	Convey("Cast []int to []string", t, func() {
		val := []int{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []int8 to []string", t, func() {
		val := []int8{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []int16 to []string", t, func() {
		val := []int16{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []int32 to []string", t, func() {
		val := []int32{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []int64 to []string", t, func() {
		val := []int64{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []uint to []string", t, func() {
		val := []uint{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []uint8 to []string", t, func() {
		val := []uint8{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []uint16 to []string", t, func() {
		val := []uint16{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []uint32 to []string", t, func() {
		val := []uint32{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []uint64 to []string", t, func() {
		val := []uint64{1, 2, 3}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"1", "2", "3"}), ShouldBeTrue)
	})

	Convey("Cast []float32 to []string", t, func() {
		val := []float32{1.1, 2.1, 3.1}
		castVal, _ := StringSlice(val)
		So(reflect.DeepEqual(castVal, []string{"1.100000023841858", "2.0999999046325684", "3.0999999046325684"}), ShouldBeTrue)
	})

	Convey("Cast []float64 to []string", t, func() {
		val := []float64{1.1, 2.1, 3.1}
		castVal, _ := StringSlice(val)
		So(reflect.DeepEqual(castVal, []string{"1.1", "2.1", "3.1"}), ShouldBeTrue)
	})

	Convey("Cast []bool to []string", t, func() {
		val := []bool{true, false, true}
		castVal, _ := StringSlice(val)
		So(reflect.DeepEqual(castVal, []string{"true", "false", "true"}), ShouldBeTrue)
	})

	Convey("Cast []string to []string", t, func() {

		val := []string{"a", "b", "c"}
		castVal, _ := StringSlice(val)

		So(reflect.DeepEqual(castVal, []string{"a", "b", "c"}), ShouldBeTrue)
	})

	Convey("Cast []NotSliceStringCaster to []string and return Error", t, func() {
		val := NotSliceStringCaster{}
		castVal, err := StringSlice(val)
		So(len(castVal), ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast []SliceStringCaster to []string", t, func() {
		val := SliceStringCaster{}
		castVal, err := StringSlice(val)
		So(reflect.DeepEqual(castVal, []string{"a", "b", "c"}), ShouldBeTrue)
		So(err, ShouldEqual, nil)
	})

	Convey("Cast []interface{} to []string", t, func() {
		val := []interface{}{"a", 1, true}
		castVal, err := StringSlice(val)
		So(reflect.DeepEqual(castVal, []string{"a", "1", "true"}), ShouldBeTrue)
		So(err, ShouldEqual, nil)
	})

	Convey("Cast []interface{} to []string with error", t, func() {
		val := []interface{}{"a", 1, true, NotSliceStringCaster{}, uint8(2)}
		castVal, err := StringSlice(val)
		So(reflect.DeepEqual(castVal, []string{"a", "1", "true", "", "2"}), ShouldBeTrue)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := SliceStringCaster{}
			MustStringSlice(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotSliceStringCaster{}
			MustStringSlice(val)
		}, ShouldPanic)
	})
}
