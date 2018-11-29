package japanese

import (
	"testing"

	"github.com/Bios-Marcel/golang-api-facade-example/greeter"
)

func TestIfInterfaceIsImplemented(t *testing.T) {
	var _ greeter.Greeter = (*GreeterJapanese)(nil)
}
