package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type IntStringCaster struct{}

func (c IntStringCaster) IntSlice() []int64 {
	return []int64{1, 2, 3}
}

type NotIntStringCaster struct{}

func TestIntSliceCast(t *testing.T) {
	Convey("Cast []int to []int64", t, func() {
		val := []int{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int8 to []int64", t, func() {
		val := []int8{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int16 to []int64", t, func() {
		val := []int16{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int32 to []int64", t, func() {
		val := []int32{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int64 to []int64", t, func() {
		val := []int64{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint to []int64", t, func() {
		val := []uint{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint8 to []int64", t, func() {
		val := []uint8{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint16 to []int64", t, func() {
		val := []uint16{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint32 to []int64", t, func() {
		val := []uint32{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint64 to []int64", t, func() {
		val := []uint64{1, 2, 3}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []float32 to []int64", t, func() {
		val := []float32{1.1, 2.1, 3.1}
		castVal, _ := IntSlice(val)
		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []float64 to []int64", t, func() {
		val := []float64{1.1, 2.1, 3.1}
		castVal, _ := IntSlice(val)
		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []bool to []int64", t, func() {
		val := []bool{true, false, true}
		castVal, _ := IntSlice(val)
		So(reflect.DeepEqual(castVal, []int64{1, 0, 1}), ShouldBeTrue)
	})

	Convey("Cast []string to []int64", t, func() {

		val := []string{"1", "2", "3"}
		castVal, _ := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast invalid []string to []int64", t, func() {

		val := []string{"1", "a", "3"}
		castVal, err := IntSlice(val)

		So(reflect.DeepEqual(castVal, []int64{1, 0, 3}), ShouldBeTrue)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast []NotIntStringCaster to []int64 and return Error", t, func() {
		val := NotIntStringCaster{}
		castVal, err := IntSlice(val)
		So(len(castVal), ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast []IntStringCaster to []int64", t, func() {
		val := IntStringCaster{}
		castVal, err := IntSlice(val)
		So(reflect.DeepEqual(castVal, []int64{1, 2, 3}), ShouldBeTrue)
		So(err, ShouldEqual, nil)
	})

	Convey("Cast []interface{} to []int64", t, func() {
		val := []interface{}{"1", 2, true}
		castVal, err := IntSlice(val)
		So(reflect.DeepEqual(castVal, []int64{1, 2, 1}), ShouldBeTrue)
		So(err, ShouldEqual, nil)
	})

	Convey("Cast []interface{} to []int64 with error", t, func() {
		val := []interface{}{"1", 1, true, NotIntStringCaster{}, uint8(2)}
		castVal, err := IntSlice(val)
		So(reflect.DeepEqual(castVal, []int64{1, 1, 1, 0, 2}), ShouldBeTrue)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := IntStringCaster{}
			MustIntSlice(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotIntStringCaster{}
			MustIntSlice(val)
		}, ShouldPanic)
	})
}
