package main

import (
	"fmt"
	"strings"
)

// Multinomial Naive Bayes
//
// Training
//
// - lights:off - turn off the lights
// - lights:off - Can you turn the lights off
// - lights:off - Lights off please
// - lights:off - Lights off
// - lights:on - turn on the lights
// - lights:on - Can you turn the lights on
// - lights:on - Lights on please
// - lights:on - Lights on
// - progress:list - what is my progress
// - progress:list - how far am i
// - progress:list - can you show me my progress
// - progress:list - what is my current lesson
//
// Test
//
// - ? - what's the current progress
//
// Priors
//
// P(class) = Occurences(class) / documents
//
// P(lights:off) = 4/12
// P(lights:on) = 4/12
// P(progress:list) = 4/12
//
// Conditional Probabilities
//
// P(Word|class) = (Occurence class + Smoothing) / (Words in class / Unique words total)
//
// P(lights|lights:off) = (4 + 1) / (15 + 22) = 5/37
// P(progress|lights:off) = (0 + 1) / (15 + 22) = 1/37
// P(on|lights:on) = (4 + 1) / (15 + 22) = 5/37
// P(lesson|progress:list) = (1 + 1) / (19 + 22) = 2/41

type Document struct {
	Class    string
	Sentence string
}

type Class struct {
	Occurence float64
	Prior     float64
	WordTotal float64
	Words     map[string]float64
}

func main() {
	var documents []Document
	documents = append(documents, Document{"lights:off", "turn off the lights"})
	documents = append(documents, Document{"lights:off", "can you turn the lights off"})
	documents = append(documents, Document{"lights:off", "lights off please"})
	documents = append(documents, Document{"lights:off", "lights off"})

	documents = append(documents, Document{"lights:on", "turn on the lights"})
	documents = append(documents, Document{"lights:on", "can you turn the lights on"})
	documents = append(documents, Document{"lights:on", "lights on please"})
	documents = append(documents, Document{"lights:on", "lights on"})

	documents = append(documents, Document{"progress:list", "what is my progress"})
	documents = append(documents, Document{"progress:list", "how far am i"})
	documents = append(documents, Document{"progress:list", "can you show me my progress"})
	documents = append(documents, Document{"progress:list", "what is my current lesson"})

	words := make(map[string]float64)
	classes := make(map[string]*Class)
	for _, v := range documents {
		if _, ok := classes[v.Class]; ok != true {
			classes[v.Class] = &Class{0.0, 0.0, 0.0, make(map[string]float64)}
		}
		classes[v.Class].Occurence++

		for _, w := range strings.Split(v.Sentence, " ") {
			classes[v.Class].Words[w]++
			classes[v.Class].WordTotal++
			words[w]++
		}
	}

	for _, v := range classes {
		v.Prior = v.Occurence / float64(len(documents))
		for w, _ := range words {
			v.Words[w] = ((v.Words[w] + 1.0) / (v.WordTotal + float64(len(words))))
		}
	}

	training := Document{
		Sentence: "what's the current progress",
	}

	proportionalProbability := make(map[string]float64)
	for k, v := range classes {
		proportionalProbability[k] = v.Prior
		for _, w := range strings.Split(training.Sentence, " ") {
			if _, ok := v.Words[w]; ok == true {
				proportionalProbability[k] *= v.Words[w]
			}
		}
	}

	var class string
	var biggest float64
	for k, v := range proportionalProbability {
		if v > biggest {
			class = k
			biggest = v
		}
	}

	fmt.Println(class)
	fmt.Println(biggest)
}