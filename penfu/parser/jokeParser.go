package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strings"
)

const jokeRe = `<div class="content-txt pt10">([^<]+)<a`

func JokeParser(context []byte, title string) engine.ParserResult {
	re := regexp.MustCompile(jokeRe)
	match := re.FindSubmatch(context)

	var result engine.ParserResult
	var joke model.Joke
	joke.Title = title
	if len(match) >= 2 {
		joke.Context = formatMatch(string(match[1]))
		result.Items = append(result.Items, joke)
	}
	return result
}

func formatMatch(match string) string {
	m := strings.Trim(match, "\n")
	m = strings.Trim(m, "\t")
	return m
}
