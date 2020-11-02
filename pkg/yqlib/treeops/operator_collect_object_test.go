package treeops

import (
	"testing"
)

var collectObjectOperatorScenarios = []expressionScenario{
	{
		document:   `{name: Mike, age: 32}`,
		expression: `{.name: .age}`,
		expected: []string{
			"D0, P[], (!!map)::Mike: 32\n",
		},
	},
	{
		document:   `{name: Mike, pets: [cat, dog]}`,
		expression: `{.name: .pets[]}`,
		expected: []string{
			"D0, P[], (!!map)::Mike: cat\n",
			"D0, P[], (!!map)::Mike: dog\n",
		},
	},
	{
		document:   `{name: Mike, pets: [cat, dog], food: [hotdog, burger]}`,
		expression: `{.name: .pets[], "f":.food[]}`,
		expected: []string{
			"D0, P[], (!!map)::Mike: cat\nf: hotdog\n",
			"D0, P[], (!!map)::Mike: cat\nf: burger\n",
			"D0, P[], (!!map)::Mike: dog\nf: hotdog\n",
			"D0, P[], (!!map)::Mike: dog\nf: burger\n",
		},
	},
	{
		document:   `{name: Mike, pets: {cows: [apl, bba]}}`,
		expression: `{"a":.name, "b":.pets}`,
		expected: []string{
			`D0, P[], (!!map)::a: Mike
b: {cows: [apl, bba]}
`,
		},
	},
	{
		document:   ``,
		expression: `{"wrap": "frog"}`,
		expected: []string{
			"D0, P[], (!!map)::wrap: frog\n",
		},
	},
	{
		document:   `{name: Mike}`,
		expression: `{"wrap": .}`,
		expected: []string{
			"D0, P[], (!!map)::wrap: {name: Mike}\n",
		},
	},
	{
		document:   `{name: Mike}`,
		expression: `{"wrap": {"further": .}}`,
		expected: []string{
			`D0, P[], (!!map)::wrap:
    further: {name: Mike}
`,
		},
	},
}

func TestCollectObjectOperatorScenarios(t *testing.T) {
	for _, tt := range collectObjectOperatorScenarios {
		testScenario(t, &tt)
	}
}