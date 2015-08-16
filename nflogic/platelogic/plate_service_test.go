package platelogic

import (
	"testing"
)

func TestCalSimilarity(t *testing.T) {
	plate_new := "AB675"
	plate_old := "AB676"
	ret := calSimilarity(plate_new, plate_old)
	t.Log(ret)
}
