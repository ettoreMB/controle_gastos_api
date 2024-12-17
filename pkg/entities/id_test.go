package pkg_entities_test

import (
	"ettoreMB/controle_gastos/internal/entities"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	"fmt"
	"testing"
)

func Test_CreateId_From_String(t *testing.T) {
	uuid := "4bb12431-6dc7-422a-aff3-12faddd8e29e"

	result, err := pkg_entities.NewIdStruct[entities.CategoryId](uuid)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}
