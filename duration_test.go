package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

type NotDuration struct{}

func TestDurationCast(t *testing.T) {
	Convey("Cast string to duration", t, func() {
		val := "1h2m3s"
		castVal, _ := Duration(val)
		dur, _ := time.ParseDuration(val)
		So(castVal, ShouldEqual, dur)
	})

	Convey("Cast string int to duration", t, func() {
		val := "123"
		castVal, _ := Duration(val)
		dur := time.Duration(123)
		So(castVal, ShouldEqual, dur)
	})

	Convey("Cast duration int to duration", t, func() {
		val := time.Duration(123)
		castVal, _ := Duration(val)
		dur := time.Duration(123)
		So(castVal, ShouldEqual, dur)
	})

	Convey("Cast invalid string to duration", t, func() {
		val := "a"
		_, err := Duration(val)
		So(err, ShouldNotBeNil)
	})

	Convey("Cast uint8 to duration", t, func() {
		val := uint8(123)
		castVal, _ := Duration(val)
		dur := time.Duration(123)
		So(castVal, ShouldEqual, dur)
	})

	Convey("Cast NotDuration to duration and return error", t, func() {
		val := NotDuration{}
		_, err := Duration(val)
		So(err, ShouldNotBeNil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := 123
			MustDuration(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotDuration{}
			MustDuration(val)
		}, ShouldPanic)
	})

}
