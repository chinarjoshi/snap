package transform

import (
	"strings"

	"github.com/chijoshi/workoutliner/paragraph"
	"github.com/chijoshi/workoutliner/table"
)

func Transform(input string) string {
	paragraphs := splitParagraphs(input)
	result := make([]string, 0, len(paragraphs))

	for _, para := range paragraphs {
		transformed := transformParagraph(para)
		result = append(result, transformed)
	}

	return strings.Join(result, "\n\n")
}

func splitParagraphs(input string) []string {
	return strings.Split(input, "\n\n")
}

func containsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func transformParagraph(para string) string {
	if !containsDigit(para) {
		return para
	}

	parseResult, err := paragraph.Parse(para)
	if err != nil {
		return para
	}

	if !parseResult.HasExercises() {
		return para
	}

	return formatResult(parseResult)
}

func formatResult(result *paragraph.ParseResult) string {
	var parts []string

	if len(result.ProseLines) > 0 {
		parts = append(parts, strings.Join(result.ProseLines, "\n"))
	}

	tableStr := table.FormatExercises(result.Exercises)
	parts = append(parts, tableStr)

	return strings.Join(parts, "\n\n")
}
