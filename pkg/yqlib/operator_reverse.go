package yqlib

import (
	"container/list"
	"fmt"
)

func reverseOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	results := list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		candidateNode := candidate.unwrapDocument()

		if candidateNode.Kind != SequenceNode {
			return context, fmt.Errorf("node at path [%v] is not an array (it's a %v)", candidate.GetNicePath(), candidate.GetNiceTag())
		}

		reverseList := candidate.CreateReplacementWithComments(SequenceNode, "!!seq", candidateNode.Style)
		reverseContent := make([]*CandidateNode, len(candidateNode.Content))

		for i, originalNode := range candidateNode.Content {
			reverseContent[len(candidateNode.Content)-i-1] = originalNode
		}
		reverseList.AddChildren(reverseContent)
		results.PushBack(reverseList)

	}

	return context.ChildContext(results), nil

}
