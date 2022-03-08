package engageapidocs

import (
	"testing"
	"io/ioutil"

	"github.com/grokify/spectrum/openapi3"
	"gopkg.in/yaml.v3"
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

type Tag struct {
	Name string
}

type TagGroup struct {
	Name string
	Tags []string
}

type TagData struct {
	Tags       []Tag
	XTagGroups []TagGroup `yaml:"x-tag-groups"`
}

func TestTags(t *testing.T) {
	for _, tt := range specTests {
		buf, err := ioutil.ReadFile(tt.filepath)
		if err != nil {
			panic(err)
		}

		tagData := &TagData{}
		err = yaml.Unmarshal(buf, tagData)
		if err != nil {
			panic(err)
		}

		var anyMissing = false;
		for _, tag := range tagData.Tags {
			var missing = true
			for _, xTagGroup := range tagData.XTagGroups {
				for _, xTagGroupTags := range xTagGroup.Tags {
					if xTagGroupTags == tag.Name {
						missing = false
					}
				}
			}
			if missing {
				anyMissing = true
				t.Errorf("[%s] [%s]: [%s] Tag is missing in x-tag-groups", tt.filepath, tt.title, tag)
			}
		}

		if !anyMissing {
			t.Logf("SPEC_IS_VALID [%s] TITLE [%s]\n", tt.filepath, tt.title)
		}
	}
}
