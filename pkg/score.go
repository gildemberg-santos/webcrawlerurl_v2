package pkg

import (
	"strings"
	"sync"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
)

var doneScore sync.WaitGroup

type Score struct {
	Information *extract.LeadsterCustom
	Rules       struct {
		TitleWebSite     float32
		MostRelevantText float32
		MetaDescription  float32
	}
}

func (s *Score) Init(information *extract.LeadsterCustom) {
	s.Information = information
}

func (s *Score) Call() {
	doneScore.Add(3)
	go s.scoreTitleWebSite()
	go s.scoreMostRelevantText()
	go s.scoreMetaDescription()
	doneScore.Wait()
}

func (s *Score) scoreTitleWebSite() {

	if s.Information.TitleWebSite != "" {
		s.Rules.TitleWebSite = scoreNumberWords(s.Information.TitleWebSite, s.Information.TitleWebSiteMin, 10)
	}
	doneScore.Done()
}

func (s *Score) scoreMostRelevantText() {
	if s.Information.MostRelevantText != "" {
		s.Rules.MostRelevantText = scoreNumberWords(s.Information.MostRelevantText, s.Information.MostRelevantTextMin, 10)
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
	return s.Rules.TitleWebSite + s.Rules.MostRelevantText + s.Rules.MetaDescription
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
