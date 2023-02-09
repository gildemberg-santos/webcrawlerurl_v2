package pkg

import "strings"

type Score struct {
	Url         string
	Information *ExtractInformation
	Rules       struct {
		MainTitle       float32
		MainParagraph   float32
		MetaDescription float32
	}
}

func (s *Score) Init(url string, information *ExtractInformation) {
	s.Url = url
	s.Information = information
}

func (s *Score) Call() {
	s.scoreMainTitle()
	s.scoreMainParagraph()
	s.scoreMetaDescription()
}

func (s *Score) scoreMainTitle() {

	if s.Information.MainTitle != "" {
		s.Rules.MainTitle = scoreNumberWords(s.Information.MainTitle, s.Information.MainTitleMin, 10)
	}
}

func (s *Score) scoreMainParagraph() {
	if s.Information.MainParagraph != "" {
		s.Rules.MainParagraph = scoreNumberWords(s.Information.MainParagraph, s.Information.MainParagraphMin, 10)
	}
}

func (s *Score) scoreMetaDescription() {
	if s.Information.MetaDescription != "" {
		s.Rules.MetaDescription = scoreNumberWords(s.Information.MetaDescription, s.Information.MetaDescriptionMin, 10)
	}
}

func (s *Score) GetScore() float32 {
	return s.Rules.MainTitle + s.Rules.MainParagraph + s.Rules.MetaDescription
}

func scoreNumberWords(text string, limitWords int, note int) float32 {
	current := len(strings.Split(text, " "))

	if current == 1 && strings.Split(text, " ")[0] == "" {
		return 0
	}

	if current >= limitWords {
		return float32(note)
	}

	score := (float32(current) / float32(limitWords)) * float32(note)
	return score
}
