package sqlbuilder

import (
	. "github.com/doug-martin/goqu/v9"
)

func Count(sd *SelectDataset, col any) *SelectDataset {
	return sd.Select(COUNT(col))
}

func DistinctCount(sd *SelectDataset, col any) *SelectDataset {
	return sd.Select(COUNT(L("DISTINCT ?", col)))
}
