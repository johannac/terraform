package schema

import (
	//"reflect"
	"testing"

	"github.com/hashicorp/terraform/terraform"
)

func TestResourceDiff_HasChange(t *testing.T) {

	schema := map[string]*Schema{
		"string": &Schema{
			Type:     TypeString,
			Optional: true,
		},
		"map": &Schema{
			Type:     TypeMap,
			Optional: true,
		},
	}

	type TestCase struct {
		State           *terraform.InstanceState
		Diff            *terraform.InstanceDiff
		StringChanged   bool
		MapChanged      bool
		MapCountChanged bool
		MapValueChanged bool
	}

	testCases := []TestCase{
		{
			State: &terraform.InstanceState{
				ID: "no change",
				Attributes: map[string]string{
					"string": "foo",
					"map.%":  "0",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{},
			},
			StringChanged:   false,
			MapChanged:      false,
			MapCountChanged: false,
			MapValueChanged: false,
		},
		{
			State: &terraform.InstanceState{
				ID: "string change",
				Attributes: map[string]string{
					"string": "foo",
					"map.%":  "0",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"string": &terraform.ResourceAttrDiff{
						Old:  "foo",
						New:  "bar",
						Type: terraform.DiffAttrInput,
					},
				},
			},
			StringChanged:   true,
			MapChanged:      false,
			MapCountChanged: false,
			MapValueChanged: false,
		},
		{
			State: &terraform.InstanceState{
				ID: "string computed",
				Attributes: map[string]string{
					"string": "foo",
					"map.%":  "0",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					// This situation would arise either because
					// we're replacing the resource or because the
					// caller already called d.SetNewComputed("string"),
					// but in practice it should only happen on
					// an attr with Computed: true.
					"string": &terraform.ResourceAttrDiff{
						Old:         "foo",
						NewComputed: true,
						Type:        terraform.DiffAttrOutput,
					},
				},
			},
			StringChanged:   true,
			MapChanged:      false,
			MapCountChanged: false,
			MapValueChanged: false,
		},
		{
			State: &terraform.InstanceState{
				ID: "map change bigger",
				Attributes: map[string]string{
					"string": "foo",
					"map.%":  "0",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"map.%": &terraform.ResourceAttrDiff{
						Old:  "0",
						New:  "1",
						Type: terraform.DiffAttrInput,
					},
					"map.foo": &terraform.ResourceAttrDiff{
						Old:  "",
						New:  "baz",
						Type: terraform.DiffAttrInput,
					},
				},
			},
			StringChanged:   false,
			MapChanged:      true,
			MapCountChanged: true,
			MapValueChanged: true,
		},
		{
			State: &terraform.InstanceState{
				ID: "map change smaller",
				Attributes: map[string]string{
					"string":  "foo",
					"map.%":   "1",
					"map.foo": "baz",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"map.%": &terraform.ResourceAttrDiff{
						Old:  "1",
						New:  "0",
						Type: terraform.DiffAttrInput,
					},
					"map.foo": &terraform.ResourceAttrDiff{
						Old:  "baz",
						New:  "",
						Type: terraform.DiffAttrInput,
					},
				},
			},
			StringChanged:   false,
			MapChanged:      true,
			MapCountChanged: true,
			MapValueChanged: true,
		},
		{
			State: &terraform.InstanceState{
				ID: "map change value",
				Attributes: map[string]string{
					"string":  "foo",
					"map.%":   "1",
					"map.foo": "baz",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"map.foo": &terraform.ResourceAttrDiff{
						Old:  "baz",
						New:  "bar",
						Type: terraform.DiffAttrInput,
					},
				},
			},
			StringChanged:   false,
			MapChanged:      true,
			MapCountChanged: false,
			MapValueChanged: true,
		},
		{
			State: &terraform.InstanceState{
				ID: "map change even bigger",
				Attributes: map[string]string{
					"string":  "foo",
					"map.%":   "1",
					"map.foo": "baz",
				},
			},
			Diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"map.%": &terraform.ResourceAttrDiff{
						Old:  "1",
						New:  "2",
						Type: terraform.DiffAttrInput,
					},
					"map.pizza": &terraform.ResourceAttrDiff{
						Old:  "",
						New:  "cheese",
						Type: terraform.DiffAttrInput,
					},
				},
			},
			StringChanged:   false,
			MapChanged:      true,
			MapCountChanged: true,
			MapValueChanged: false,
		},
	}

	for _, tc := range testCases {
		testName := tc.State.ID
		d, err := schemaMap(schema).DiffWrapper(tc.State, tc.Diff)
		if err != nil {
			t.Errorf("test %s: error creating ResourceDiff object: %s", testName, err)
			continue
		}

		if want, got := tc.StringChanged, d.HasChange("string"); want != got {
			t.Errorf(
				"test %s: HasChange(\"string\") returned %v; want %v",
				testName,
				got, want,
			)
		}

		if want, got := tc.MapChanged, d.HasChange("map"); want != got {
			t.Errorf(
				"test %s: HasChange(\"map\") returned %v; want %v",
				testName,
				got, want,
			)
		}

		if want, got := tc.MapCountChanged, d.HasChange("map.%"); want != got {
			t.Errorf(
				"test %s: HasChange(\"map.%%\") returned %v; want %v",
				testName,
				got, want,
			)
		}

		if want, got := tc.MapValueChanged, d.HasChange("map.foo"); want != got {
			t.Errorf(
				"test %s: HasChange(\"map.foo\") returned %v; want %v",
				testName,
				got, want,
			)
		}
	}
}
