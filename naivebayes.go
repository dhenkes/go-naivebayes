// This package provides a simple multinomial naive bayes implementation.
//
// Copyright (C) 2017 by Dominique Henkes <hello@henkes.io>
package naivebayes

import "strings"

const defaultProb = 0.00000000001

// Class contains the document count, an array of all words used in those documents
// (the array contains duplicates) and a map with the word frequency which can
// be used to obtain the unique word count.
type Class struct {
	Documents float64
	Words     []string
	WordFreq  map[string]float64
}

// Classifier gives us total document count, the length of the ngrams being used,
// a map with all classes and a map with the word frequency which again can be
// used to obtain the unique word count.
type Classifier struct {
	Ngram       int
	Documents   float64
	Classes     map[string]*Class
	UniqueWords map[string]float64
}

// GetNgrams returns an array of sequences of n items. The length n is defined
// by the parameter size.
//
// The input (1, "this outputs ngrams") would be ["this", "outputs", "ngrams"].
//
// The input (2, "this outputs ngrams") would be ["this outputs", "outputs ngrams"].
func GetNgrams(size int, sentence string) []string {
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

// NewClassifier returns a new classifier which initiates two empty maps. This
// could later be improved so that everything is saved more efficiently.
func NewClassifier(n int) *Classifier {
	return &Classifier{
		Ngram:       n,
		Classes:     make(map[string]*Class),
		UniqueWords: make(map[string]float64),
	}
}

// Train adds the ngrams of a sentence to an existing or new class.
func (c *Classifier) Train(class string, sentence string) {
	c.Documents++
	_, exists := c.Classes[class]
	if exists == false {
		c.Classes[class] = &Class{
			WordFreq: make(map[string]float64),
		}
	}

	c.Classes[class].Documents++
	words := GetNgrams(c.Ngram, sentence)
	for _, word := range words {
		c.UniqueWords[word]++
		c.Classes[class].Words = append(c.Classes[class].Words, word)
		c.Classes[class].WordFreq[word]++
	}
}

// GetPrior returns the prior probabilities of a document being in a specific
// class. It is calculated by dividing the class frequency by the total amount
// of documents.
func (c *Classifier) GetPrior(class string) float64 {
	return c.Classes[class].Documents / c.Documents
}

// Classify returns the probabilities for a sentence belonging to a
// certain class. These probabilities are calculated by taking the class prior
// P(class) and multiplying it by the conditional probabilities P(word|class).
func (c *Classifier) Classify(sentence string) map[string]float64 {
	probabilities := make(map[string]float64)
	uniqueWordCount := float64(len(c.UniqueWords))
	words := GetNgrams(c.Ngram, sentence)

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
