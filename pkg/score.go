package pkg

import (
	"strings"
	"sync"
)

var doneScore sync.WaitGroup

type Score struct {
	Information *ExtractInformation
	Rules       struct {
		MainTitle       float32
		MainParagraph   float32
		MetaDescription float32
	}
}

func (s *Score) Init(information *ExtractInformation) {
	s.Information = information
}

func (s *Score) Call() {
	doneScore.Add(3)
	go s.scoreMainTitle()
	go s.scoreMainParagraph()
	go s.scoreMetaDescription()
	doneScore.Wait()
}

func (s *Score) scoreMainTitle() {

	if s.Information.MainTitle != "" {
		s.Rules.MainTitle = scoreNumberWords(s.Information.MainTitle, s.Information.MainTitleMin, 10)
	}
	doneScore.Done()
}

func (s *Score) scoreMainParagraph() {
	if s.Information.MainParagraph != "" {
		s.Rules.MainParagraph = scoreNumberWords(s.Information.MainParagraph, s.Information.MainParagraphMin, 10)
	}
	doneScore.Done()
}

func (s *Score) scoreMetaDescription() {
	if s.Information.MetaDescription != "" {
		s.Rules.MetaDescription = scoreNumberWords(s.Information.MetaDescription, s.Information.MetaDescriptionMin, 10)
	}
	doneScore.Done()
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
