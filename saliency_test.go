package ysgo

import (
	"testing"

	"github.com/remieven/ysgo/internal/rng"
	"github.com/remieven/ysgo/internal/testutils"
)

// ensure that *FirstSaliencyStrategy satisfies ContentSaliencyStrategy
var _ ContentSaliencyStrategy = (*FirstSaliencyStrategy)(nil)

func TestFirstSaliencyStrategy(t *testing.T) {
	createStrategy := func() ContentSaliencyStrategy {
		return &FirstSaliencyStrategy{}
	}

	tests := map[string]saliencyStrategyTest{
		"no options given": {
			strategySupplier:       createStrategy,
			options:                []*ContentSaliencyOption{},
			expectedSelectedStatus: false,
		},
		"several options given": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    0,
		},
	}

	for name, test := range tests {
		t.Run(name, test.Run)
	}
}

// ensure that *BestSaliencyStrategy satisfies ContentSaliencyStrategy
var _ ContentSaliencyStrategy = (*BestSaliencyStrategy)(nil)

func TestBestSaliencyStrategy(t *testing.T) {
	createStrategy := func() ContentSaliencyStrategy {
		return &BestSaliencyStrategy{}
	}

	tests := map[string]saliencyStrategyTest{
		"no options given": {
			strategySupplier:       createStrategy,
			options:                []*ContentSaliencyOption{},
			expectedSelectedStatus: false,
		},
		"several options given with different complexities": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    1,
		},
		"several options given with same complexities": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  3,
				},
				{
					ContentID:  "40eaf7",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    1,
		},
	}

	for name, test := range tests {
		t.Run(name, test.Run)
	}
}

// ensure that *BestLeastRecentlyViewedSaliencyStrategy satisfies ContentSaliencyStrategy
var _ ContentSaliencyStrategy = (*BestLeastRecentlyViewedSaliencyStrategy)(nil)

func TestBestLeastRecentlyViewedSaliencyStrategy(t *testing.T) {
	createStrategy := func() ContentSaliencyStrategy {
		return &BestLeastRecentlyViewedSaliencyStrategy{}
	}

	tests := map[string]saliencyStrategyTest{
		"no options given": {
			strategySupplier:       createStrategy,
			options:                []*ContentSaliencyOption{},
			expectedSelectedStatus: false,
		},
		"several options given with different complexities": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    1,
		},
		"several options given with same complexities but different view counts": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  4,
				},
				{
					ContentID:  "40eaf7",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    2,
		},
		"several options given with same complexities and view counts": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 1,
					ViewCount:  2,
				},
				{
					ContentID:  "40eaf7",
					Complexity: 2,
					ViewCount:  3,
				},
				{
					ContentID:  "3a6c94",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    2,
		},
	}

	for name, test := range tests {
		t.Run(name, test.Run)
	}
}

func TestRandomBestLeastRecentlyViewedSaliencyStrategy(t *testing.T) {
	createStrategy := func() ContentSaliencyStrategy {
		rng, err := rng.NewRNG("1234")
		if err != nil {
			panic("failed to create rng: " + err.Error())
		}

		return &RandomBestLeastRecentlyViewedSaliencyStrategy{rng}
	}

	tests := map[string]saliencyStrategyTest{
		"no options given": {
			strategySupplier:       createStrategy,
			options:                []*ContentSaliencyOption{},
			expectedSelectedStatus: false,
		},
		"several options given with different complexities": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    1,
		},
		"several options given with same complexities but different view counts": {
			strategySupplier: createStrategy,
			options: []*ContentSaliencyOption{
				{
					ContentID:  "f65d07",
					Complexity: 0,
					ViewCount:  3,
				},
				{
					ContentID:  "1b159b",
					Complexity: 2,
					ViewCount:  4,
				},
				{
					ContentID:  "40eaf7",
					Complexity: 2,
					ViewCount:  3,
				},
			},
			expectedSelectedStatus: true,
			expectedOptionIndex:    2,
		},
	}

	for name, test := range tests {
		t.Run(name, test.Run)
	}

	t.Run("several options given with same complexities and view counts", func(t *testing.T) {
		strategy := createStrategy()
		options := []*ContentSaliencyOption{
			{
				ContentID:  "f65d07",
				Complexity: 0,
				ViewCount:  3,
			},
			{
				ContentID:  "1b159b",
				Complexity: 1,
				ViewCount:  2,
			},
			{
				ContentID:  "40eaf7",
				Complexity: 2,
				ViewCount:  3,
			},
			{
				ContentID:  "3a6c94",
				Complexity: 2,
				ViewCount:  3,
			},
		}

		selectedContentIds := map[string]int{}

		for range 100 {
			selectedContent, ok := strategy.QueryBestContent(options)
			if !ok {
				t.Error("unexpected ok = false")
				return
			}
			if _, ok := selectedContentIds[selectedContent.ContentID]; !ok {
				selectedContentIds[selectedContent.ContentID] = 1 
			} else {
				selectedContentIds[selectedContent.ContentID]++
			}
		}

		if selectedContentIds["40eaf7"] == 0 {
			t.Errorf("expected content with id [40eaf7] to have been selected at least once but it wasn't")
			return
		}
		if selectedContentIds["3a6c94"] == 0 {
			t.Errorf("expected content with id [3a6c94] to have been selected at least once but it wasn't")
			return
		}

		if len(selectedContentIds) != 2 {
			t.Errorf("more content was selected than expected: got [%v], wanted [%v]", len(selectedContentIds), 2)
		}
	})
}

type saliencyStrategyTest struct {
	strategySupplier       func() ContentSaliencyStrategy
	options                []*ContentSaliencyOption
	expectedSelectedStatus bool
	expectedOptionIndex    int
}

func (sst *saliencyStrategyTest) Run(t *testing.T) {
	strategy := sst.strategySupplier()
	selectedOption, ok := strategy.QueryBestContent(sst.options)
	if ok != sst.expectedSelectedStatus {
		t.Error("either expected an option to be selected but none were, or expected no option to be selected but one was")
		return
	}
	if ok != (selectedOption != nil) {
		t.Error("ok return value wasn't consistent with selectedOption return value")
		return
	}
	if !ok {
		return
	}
	if diff := testutils.DeepEqual(selectedOption, sst.options[sst.expectedOptionIndex]); diff != "" {
		t.Error("unexpected selected option: " + diff)
		return
	}
}
