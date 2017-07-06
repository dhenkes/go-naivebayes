# Multinomial Naive Bayes

This package provides a simple multinomial naive bayes implementation.

## Example

```go
import (
  nb "github.com/dhenkes/go-naivebayes"
)

classifier := nb.NewClassifier(1)

classifier.Train("class1", "turn the lights off")
classifier.Train("class1", "lights off")

classifier.Train("class2", "what is my progress")
classifier.Train("class2", "how far am i")

probabilities := classifier.Classify("can you switch the lights off")

// probabilities -> map[class1:0.16666666666666666 class2:0.0500000000005]
```
