package factory

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)



func Test_factory(t * testing.T) {
	Convey("TestFactory", t, func() {
		fmt.Println("yes.....")
		opr := operationFactory(1, 2, "+")
		So(opr, ShouldNotBeNil)
		fmt.Println(opr.getResult())

		opr = operationFactory(121, 21, "-")
		So(opr, ShouldNotBeNil)
		fmt.Println(opr.getResult())

		//opr = operationFactory(121, 21, "*")
		//So(opr, ShouldNotBeNil)
	})
}
