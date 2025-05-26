package services

import (
	"GOTASK/combo"
	"GOTASK/model"
	"time"
)

func TextAnalysis(Str string) (model.Results, time.Duration) {
	startTime := time.Now()
	resultchann := make(chan combo.ResultStruct)
	go combo.Combo(Str, resultchann)
	value := <-resultchann
	timeTaken := time.Since(startTime)
	result := model.Results{
		Words:      value.Words,
		Lines:      value.Lines,
		Spaces:     value.Spaces,
		Punc:       value.Punc,
		Characters: value.Characters,
		Sentences:  value.Sentences,
		Vowels:     value.Vowels,
		Consonants: value.Consonants,
		Digits:     value.Digits,
		Para:       value.Para,
	}

	return result, timeTaken
}
