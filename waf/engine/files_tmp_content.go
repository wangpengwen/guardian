package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadFilesTmpContent() *TransactionMap {
	t.variableMap["FILES_TMP_CONTENT"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
