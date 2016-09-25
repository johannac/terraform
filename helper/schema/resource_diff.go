package schema

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform/terraform"
)

// ResourceDiff is used to query and customize a resource diff.
// It is the analog of ResourceData that is used during the plan phase
// to tell Terraform about any side-effects of a change that it
// cannot infer automatically just by comparing directly the
// values in the state and the configuration.
//
// Two types of diff customization are permitted:
// - Create a new attribute diff for any attribute that is declared as
//   Computed, e.g. if the implementation knows that changing one
//   attribute will cause another computed one to change as a side-effect.
// - Annotate an already-detected diff as forcing a new resource, for
//   situations where some changes can be applied directly while other
//   changes require the resource to be rebuilt.
//
// No other diff customizations are allowed because we cannot permit the
// diff to describe a change whose new state is not consistent with
// what is represented in the configuration.
type ResourceDiff struct {
	schema map[string]*Schema
	state  *terraform.InstanceState
	diff   *terraform.InstanceDiff

	// Wrap the state and diff so we can easily extract typed values from them.
	stateReader FieldReader
	diffReader  FieldReader
}

// The public interface to this is schemaMap.DiffWrapper
func newResourceDiff(
	schema map[string]*Schema,
	state *terraform.InstanceState,
	diff *terraform.InstanceDiff,
) *ResourceDiff {

	stateReader := &MapFieldReader{
		Map:    BasicMapReader(state.Attributes),
		Schema: schema,
	}

	diffReader := &DiffFieldReader{
		Diff:   diff,
		Source: stateReader,
		Schema: schema,
	}

	return &ResourceDiff{
		schema: schema,
		state:  state,
		diff:   diff,

		stateReader: stateReader,
		diffReader:  diffReader,
	}
}

// HasAnyChanges returns true if there are any attribute changes described
// in the diff at all.
func (d *ResourceDiff) HasAnyChanges() bool {
	return !d.diff.Empty()
}

// WillUpdate returns true if applying this diff will result in a call
// to the resource's "Update" function.
func (d *ResourceDiff) WillUpdate() bool {
	return d.diff.ChangeType() == terraform.DiffUpdate
}

// WillCreate returns true if applying this diff will result in a call
// to the resource's "Create" function, either because this is a new
// resource or because a change is forcing it to be replaced.
func (d *ResourceDiff) WillCreate() bool {
	changeType := d.diff.ChangeType()
	return changeType == terraform.DiffCreate || changeType == terraform.DiffDestroyCreate
}

// HasChange returns true if the named attribute has a change represented
// in the diff.
func (d *ResourceDiff) HasChange(key string) bool {
	addr := d.splitKey(key)

	oldResult, err := d.stateReader.ReadField(addr)
	if err != nil {
		// should never happen
		panic(err)
	}
	newResult, err := d.diffReader.ReadField(addr)
	if err != nil {
		// should never happen
		panic(err)
	}

	if oldResult.Exists != newResult.Exists {
		return true
	}
	if newResult.Computed {
		return true
	}

	oldValue := oldResult.Value
	newValue := newResult.Value

	// Since both values are read using the same schema we can
	// assume that they have the same type. If that type implements
	// "Equal" then we'll use that, or else we'll just do a standard
	// deep equality.
	if eqOld, ok := oldValue.(Equal); ok {
		return !eqOld.Equal(newValue)
	} else {
		return !reflect.DeepEqual(oldValue, newValue)
	}
}

// GetChange returns the old and new values for the named attribute.
//
// Use HasChange first to robustly check for differences, rather than
// trying to compare the results of this method directly.
func (d *ResourceDiff) GetChange(key string) (interface{}, interface{}) {
	addr := d.splitKey(key)

	oldResult, err := d.stateReader.ReadField(addr)
	if err != nil {
		// should never happen
		panic(err)
	}
	newResult, err := d.diffReader.ReadField(addr)
	if err != nil {
		// should never happen
		panic(err)
	}

	return oldResult.Value, newResult.Value
}

// SetNewComputed creates a new attribute diff for the given key that
// marks the new value as "computed". The attribute addressed by "key"
// must have Computed: true set.
func (d *ResourceDiff) SetNewComputed(key string) error {
	panic("SetNewComputed not yet implemented")
}

// SetNewComputed creates a new attribute diff for the given key that
// assigns a new concrete value. The attribute addressed by "key" must
// have Computed: true set.
func (d *ResourceDiff) SetNewValue(key string, value interface{}) error {
	panic("SetNewValue not yet implemented")
}

// SetForcesNew marks an existing attribute diff as "forces new
// resource". The attribute addressed by "key" must either have "Optional:true"
// or "Required:true" set, or in other words it must be a change that comes
// from the configuration.
func (d *ResourceDiff) SetForcesNew(key string) error {
	panic("SetForcesNew not yet implemented")
}

func (d *ResourceDiff) splitKey(key string) []string {
	return strings.Split(key, ".")
}
