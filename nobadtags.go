package ykp

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"strings"
)

var t_prefix = []string{"t"}

// content is able to have #UpperCasedTag, But the ["t"] must in lowercase.
// if the event has one with uppercased string on t tag, Reject it.
func NoUpperCaseValueOfTTag(_ context.Context, event *nostr.Event) (reject bool, msg string) {
	t_tags := event.Tags.GetAll(t_prefix)
	for _, tag := range t_tags {
		val := tag.Value()
		lowerCasedValue := strings.ToLower(val)
		if val != lowerCasedValue {
			return true, "t tag value must in lowercase."
		}
	}

	return false, ""
}
