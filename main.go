package main

import (
	"fmt"
	"strings"
)

const defaultProb = 0.00000000001

type Class struct {
	Documents float64
	Words     []string
	WordFreq  map[string]float64
}

type Classifier struct {
	Ngram       int
	Documents   float64
	Classes     map[string]*Class
	UniqueWords map[string]int
}

func GetNGrams(size int, sentence string) []string {
	ngrams := []string{}
	words := strings.Split(sentence, " ")

	if len(words) <= size {
		ngrams = append(ngrams, strings.Join(words, " "))
		return ngrams
	}

	for i := 0; i+size <= len(words); i++ {
		ngrams = append(ngrams, strings.Join(words[i:i+size], " "))
	}

	return ngrams
}

func NewClassifier(n int) *Classifier {
	return &Classifier{
		Ngram:       n,
		Classes:     make(map[string]*Class),
		UniqueWords: make(map[string]int),
	}
}

func (c *Classifier) AddWord(class string, word string) {
	c.UniqueWords[word]++
	c.Classes[class].Words = append(c.Classes[class].Words, word)
	c.Classes[class].WordFreq[word]++
}

func (c *Classifier) Learn(class string, sentence string) {
	c.Documents++
	_, exists := c.Classes[class]
	if exists == false {
		c.Classes[class] = &Class{
			WordFreq: make(map[string]float64),
		}
	}

	c.Classes[class].Documents++
	words := GetNGrams(c.Ngram, sentence)
	for _, word := range words {
		c.AddWord(class, word)
	}
}

// GetPrior returns the prior probabilities of a document being in a specific
// class. It is calculated by dividing the class frequency by the total amount
// of documents.
func (c *Classifier) GetPrior(class string) float64 {
	return c.Classes[class].Documents / c.Documents
}

func (c *Classifier) GetProbabilities(sentence string) map[string]float64 {
	probabilities := make(map[string]float64)
	uniqueWordCount := float64(len(c.UniqueWords))
	words := GetNGrams(c.Ngram, sentence)

	for class, data := range c.Classes {
		probabilities[class] = c.GetPrior(class)
		classWordCount := float64(len(data.Words))
		for _, word := range words {
			frequency, exists := data.WordFreq[word]
			if exists == false {
				frequency = defaultProb
			}

			probabilities[class] = (frequency + 1.0) / (classWordCount + uniqueWordCount)
		}
	}

	return probabilities
}

func main() {
	// Everything in main is just for testing

	unigram := NewClassifier(1)
	unigram.Learn("lights:toggle", "lights")
	unigram.Learn("lights:off", "turn off the lights")
	unigram.Learn("lights:off", "can you turn the lights off")
	unigram.Learn("lights:off", "lights off please")
	unigram.Learn("lights:off", "lights off")
	unigram.Learn("lights:on", "turn on the lights")
	unigram.Learn("lights:on", "can you turn the lights on")
	unigram.Learn("lights:on", "lights on please")
	unigram.Learn("lights:on", "lights on")
	unigram.Learn("progress:list", "what is my progress")
	unigram.Learn("progress:list", "how far am i")
	unigram.Learn("progress:list", "can you show me my progress")
	unigram.Learn("progress:list", "what is my current lesson")

	bigram := NewClassifier(2)
	bigram.Learn("lights:toggle", "lights")
	bigram.Learn("lights:off", "turn off the lights")
	bigram.Learn("lights:off", "can you turn the lights off")
	bigram.Learn("lights:off", "lights off please")
	bigram.Learn("lights:off", "lights off")
	bigram.Learn("lights:on", "turn on the lights")
	bigram.Learn("lights:on", "can you turn the lights on")
	bigram.Learn("lights:on", "lights on please")
	bigram.Learn("lights:on", "lights on")
	bigram.Learn("progress:list", "what is my progress")
	bigram.Learn("progress:list", "how far am i")
	bigram.Learn("progress:list", "can you show me my progress")
	bigram.Learn("progress:list", "what is my current lesson")

	uniProb1 := unigram.GetProbabilities("lights")
	biProb1 := bigram.GetProbabilities("lights")
	for k, v := range uniProb1 {
		fmt.Printf("Input:   lights\n")
		fmt.Printf("Class:   %s\n", k)
		fmt.Printf("Unigram: %f\n", v)
		fmt.Printf("Bigram:  %f\n\n", biProb1[k])
	}

	fmt.Printf("-----\n\n")

	uniProb2 := unigram.GetProbabilities("could you turn the lights off")
	biProb2 := bigram.GetProbabilities("could you turn the lights off")
	for k, v := range uniProb2 {
		fmt.Printf("Input:   could you turn the lights off\n")
		fmt.Printf("Class:   %s\n", k)
		fmt.Printf("Unigram: %f\n", v)
		fmt.Printf("Bigram:  %f\n\n", biProb2[k])
	}

	fmt.Printf("-----\n\n")

	uniProb3 := unigram.GetProbabilities("whats the current progress")
	biProb3 := bigram.GetProbabilities("whats the current progress")
	for k, v := range uniProb3 {
		fmt.Printf("Input:   what's the current progress\n")
		fmt.Printf("Class:   %s\n", k)
		fmt.Printf("Unigram: %f\n", v)
		fmt.Printf("Bigram:  %f\n\n", biProb3[k])
	}
}
