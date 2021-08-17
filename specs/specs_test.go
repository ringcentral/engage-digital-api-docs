package engageapidocs

import (
	"testing"

	"github.com/grokify/spectrum/openapi3"
)

var specTests = []struct {
	filepath string
	title    string
}{
	{"engage-digital_openapi3.yaml", "RingCentral Engage Digital API"},
}

// TestSpecs test reading specs.
func TestSpecs(t *testing.T) {
	for _, tt := range specTests {
		spec, err := openapi3.ReadAndValidateFile(tt.filepath)
		if err != nil {
			t.Errorf("openapi3.ReadAndValidateFile('%s') Error [%s]", tt.filepath, err.Error())
		} else if tt.title != spec.Info.Title {
			t.Errorf("openapi3.ReadAndValidateFile('%s') Want [%s] Got [%s]", tt.filepath, tt.title, spec.Info.Title)
		} else {
			specmore := openapi3.SpecMore{Spec: spec}
			_, opParamsWo, opParamsAll := specmore.OperationParametersDescriptionStatusCounts()
			_, schPropsWo, schPropsAll := specmore.SchemaPropertiesDescriptionStatusCounts()
			t.Logf("SPEC_IS_VALID [%s] TITLE [%s] OP_COUNT [%d] OP_PARAM_WO_DESC [%d/%d] SCH_PROP_WO_DESC [%d/%d]\n",
				tt.filepath, spec.Info.Title, specmore.OperationsCount(),
				opParamsWo, opParamsAll,
				schPropsWo, schPropsAll,
			)
		}
	}
}
