package model
import(
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)
func TestSaveExpense(t *testing.T) {
	Convey("Save Expense IsLunch true ",t, func() {
		exp := Expense{
			IsLunch: true,
		}
		err := SaveExpense(exp)
		Convey("Error should be nil",func(){
			So(err, ShouldBeNil)
		})
	})
}
