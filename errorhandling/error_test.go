package errorhandling

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestHello(t *testing.T) {

	//assertCorrectMessage := func(t *testing.T, got, want string) {
	//	t.Helper()
	//	if got != want {
	//		t.Errorf("got %q want %q", got, want)
	//	}
	//}

	t.Run("Fehler Messages", func(t *testing.T) {
		err := ShowErrors()
		//fmt.Println(errors.Cause(err))
		if errors.Cause(err) == ErrNoOne {
			fmt.Printf("GCP down:-) , %v\n", err)
			fmt.Printf("%+v\n", err) // %v as format parameter, change the format parameter to %+v ，to get the complete call stack.
			return
		}
		if err != nil {
			// unknown error
		}
	})

}
