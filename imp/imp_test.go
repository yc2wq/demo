package imp
import (
	"testing"
	"strings"
)

func TestShow(t *testing.T) {
	result := Show("tomxu", 30)
	if !strings.HasPrefix(result, "My Name is") {
		t.Errorf("Greet() = %s; My Name is tomxu, My age is 30", result)
	}

}
