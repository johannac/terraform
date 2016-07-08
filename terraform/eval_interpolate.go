package terraform

import (
	"log"

	"github.com/hashicorp/terraform/config"
)

// EvalInterpolate is an EvalNode implementation that takes a raw
// configuration and interpolates it.
type EvalInterpolate struct {
	Config   *config.RawConfig
	Resource *Resource
	Output   **ResourceConfig
	Destroy  bool
}

func (n *EvalInterpolate) Eval(ctx EvalContext) (interface{}, error) {
	log.Printf("[XXXX] (I am in %T) EvalInterpolate:\n", ctx.CurrentVertex(), ctx.currentOp())
	rc, err := ctx.Interpolate(n.Config, n.Resource)
	if err != nil {
		return nil, err
	}

	if n.Output != nil {
		*n.Output = rc
	}

	return nil, nil
}
