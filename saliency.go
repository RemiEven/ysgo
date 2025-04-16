package ysgo

import (
	"slices"

	"github.com/remieven/ysgo/internal/rng"
)

// ContentSaliencyOption represents a piece of content that may be selected by an saliency strategy.
type ContentSaliencyOption struct {
	ContentID  string
	Complexity int
	ViewCount  int
}

// ContentSaliencyStrategy chooses a piece of content from a collection of options.
type ContentSaliencyStrategy interface {

	// QueryBestContent chooses an item between the given options that is the most
	// appropriate (or salient) for the user's current context.
	// options can be empty.
	// Implementations should be read-only.
	QueryBestContent(options []*ContentSaliencyOption) (*ContentSaliencyOption, bool)

	// ContentWasSelected is will be called by the dialogue runner when a piece
	// of salient content has been selected, and the strategy should update any
	// state related to how it selects content.
	// If a content saliency strategy does not need to keep track of any state,
	// then this method can be empty.
	ContentWasSelected(option *ContentSaliencyOption)
}

// FirstSaliencyStrategyName is the name used to select FirstSaliencyStrategy
// in a dialogue runner.
const FirstSaliencyStrategyName = "FirstSaliencyStrategy"

// FirstSaliencyStrategy always returns the first non-failing
// item in the list of available options.
type FirstSaliencyStrategy struct{}

// QueryBestContent returns the first option, if any.
func (fss *FirstSaliencyStrategy) QueryBestContent(options []*ContentSaliencyOption) (*ContentSaliencyOption, bool) {
	if len(options) == 0 {
		return nil, false
	}
	return options[0], true
}

// ContentWasSelected doesn't do anything.
func (fss *FirstSaliencyStrategy) ContentWasSelected(option *ContentSaliencyOption) {}

// BestSaliencyStrategyName is the name used to select BestSaliencyStrategy
// in a dialogue runner.
const BestSaliencyStrategyName = "BestSaliencyStrategy"

// BestSaliencyStrategy returns the first option (if any) between the best of the
// provided options.
type BestSaliencyStrategy struct{}

// QueryBestContent returns the first option (if any) between the best of the
// provided options.
func (bss *BestSaliencyStrategy) QueryBestContent(options []*ContentSaliencyOption) (*ContentSaliencyOption, bool) {
	if len(options) == 0 {
		return nil, false
	}
	return slices.MaxFunc(options, func(a, b *ContentSaliencyOption) int {
		return a.Complexity - b.Complexity
	}), true
}

// ContentWasSelected doesn't do anything.
func (bss *BestSaliencyStrategy) ContentWasSelected(option *ContentSaliencyOption) {}

// BestLeastRecentlyViewedSaliencyStrategyName is the name used to select BestLeastRecentlyViewedSaliencyStrategy
// in a dialogue runner.
const BestLeastRecentlyViewedSaliencyStrategyName = "BestLeastRecentlyViewedSaliencyStrategy"

// BestLeastRecentlyViewedSaliencyStrategy returns the first (if any) of the best,
// least-recently seen choices from the provided options.
type BestLeastRecentlyViewedSaliencyStrategy struct{}

// QueryBestContent returns the first (if any) of the best,
// least-recently seen choices from the provided options.
func (blrvss *BestLeastRecentlyViewedSaliencyStrategy) QueryBestContent(options []*ContentSaliencyOption) (*ContentSaliencyOption, bool) {
	if len(options) == 0 {
		return nil, false
	}
	return slices.MaxFunc(options, func(a, b *ContentSaliencyOption) int {
		if a.Complexity != b.Complexity {
			return a.Complexity - b.Complexity
		}
		return b.ViewCount - a.ViewCount
	}), true
}

// ContentWasSelected doesn't do anything.
func (blrvss *BestLeastRecentlyViewedSaliencyStrategy) ContentWasSelected(option *ContentSaliencyOption) {
}

// RandomBestLeastRecentlyViewedSaliencyStrategyName is the name used to select RandomBestLeastRecentlyViewedSaliencyStrategy
// in a dialogue runner.
const RandomBestLeastRecentlyViewedSaliencyStrategyName = "RandomBestLeastRecentlyViewedSaliencyStrategy"

// RandomBestLeastRecentlyViewedSaliencyStrategy returns a random choice of the best,
// least-recently seen choices from the provided options (if any).
type RandomBestLeastRecentlyViewedSaliencyStrategy struct {
	rng *rng.RNG
}

// QueryBestContent returns a random choice of the best,
// least-recently seen choices from the provided options (if any).
func (rblrvss *RandomBestLeastRecentlyViewedSaliencyStrategy) QueryBestContent(options []*ContentSaliencyOption) (*ContentSaliencyOption, bool) {
	if len(options) == 0 {
		return nil, false
	}
	candidateOptions := make([]*ContentSaliencyOption, 0, len(options))
	candidateOptions = append(candidateOptions, options[0])
	for _, option := range options[1:] {
		candidateOption := candidateOptions[0]
		if option.Complexity > candidateOption.Complexity {
			candidateOptions = candidateOptions[:1]
			candidateOptions[0] = option
		} else if option.Complexity == candidateOption.Complexity && option.ViewCount < candidateOption.ViewCount {
			candidateOptions = candidateOptions[:1]
			candidateOptions[0] = option
		} else if option.Complexity == candidateOption.Complexity && option.ViewCount == candidateOption.ViewCount {
			candidateOptions = append(candidateOptions, option)
		}
	}
	return candidateOptions[rblrvss.rng.IntBetween(0, len(candidateOptions)-1)], true
}

// ContentWasSelected doesn't do anything.
func (rblrvss *RandomBestLeastRecentlyViewedSaliencyStrategy) ContentWasSelected(option *ContentSaliencyOption) {
}

func viewCountVariableNameForContent(contentID string) string {
	return "$ysgo.internal.content.viewCount." + contentID
}
