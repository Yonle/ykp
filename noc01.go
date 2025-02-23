package ykp

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
)

// NOC-01 Spec Compliance Policies
// See the spec [here](https://github.com/Yonle/Nocturn/blob/af0970a912ddb9a5a3573da68d35626319a650b1/01.md)

func NOC01_FilterReject(ctx context.Context, filter nostr.Filter) (reject bool, msg string) {
	tags := filter.Tags

	hasIDs := len(filter.IDs) > 0
	hasAuthors := len(filter.Authors) > 0

	e, e_ok := tags["e"]
	p, p_ok := tags["p"]
	t, t_ok := tags["t"]

	hasETag := e_ok && len(e) > 0
	hasPTag := p_ok && len(p) > 0
	hasTTag := t_ok && len(t) > 0

	// if it doesn't have one of these tags, reject.
	if !hasIDs && !hasAuthors && !hasETag && !hasPTag && !hasTTag {
		return true, "your filter doesn't has one of the required field as defined on NOC-01."
	} else {
		return false, ""
  }
}
