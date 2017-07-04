# Multinomial Naive Bayes

## Training

- lights:off - turn off the lights
- lights:off - Can you turn the lights off
- lights:off - Lights off please
- lights:off - Lights off
- lights:on - turn on the lights
- lights:on - Can you turn the lights on
- lights:on - Lights on please
- lights:on - Lights on
- progress:list - what is my progress
- progress:list - how far am i
- progress:list - can you show me my progress
- progress:list - what is my current lesson

## Test

- ? - what's the current progress

## Priors

P(class) = occurrences(class) / documents

P(lights:off) = 4/12

P(lights:on) = 4/12

P(progress:list) = 4/12

## Conditional Probabilities

P(word|class) = (occurrence class + smoothing) / (words in class / unique words total)

P(lights|lights:off) = (4 + 1) / (15 + 22) = 5/37

P(progress|lights:off) = (0 + 1) / (15 + 22) = 1/37

P(on|lights:on) = (4 + 1) / (15 + 22) = 5/37

P(lesson|progress:list) = (1 + 1) / (19 + 22) = 2/41
