package terraform

import (
	"fmt"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/dag"
)

// EvalInterpolate is an EvalNode implementation that takes a raw
// configuration and interpolates it.
type EvalInterpolate struct {
	Config   *config.RawConfig
	Resource *Resource
	Output   **ResourceConfig
	Comment  string
}

func (n *EvalInterpolate) Eval(ctx EvalContext) (interface{}, error) {
	rc, err := ctx.Interpolate(n.Config, n.Resource)
	if err != nil {
		return nil, fmt.Errorf("V: %s, COMMENT: %s, ERR: %s", dag.VertexName(ctx.CurrentVertex()), n.Comment, err)
	}

	if n.Output != nil {
		*n.Output = rc
	}

	return nil, nil
}
