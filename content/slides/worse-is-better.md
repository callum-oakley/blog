```
template = "slides"
title = "Worse is better"
date = "13th of April 2018"
```

# worse is better

* * *

## context

- it's 1991

- you are Richard P. Gabriel, computer scientist

- you love LISP

- everyone else loves C

* * *

# what gives?

* * *

- C is worse than LISP

- but C won the race

- therefore worse is better

- you write [an essay][] describing the worse is better approach

* * *

# the right thing

* * *

## the right thing

### simplicity

- the design must be simple, both in implementation and interface

- it is more important for the interface to be simple than the implementation

* * *

## the right thing

### correctness

- the design must be correct in all observable aspects

- incorrectness is simply not allowed

* * *

## the right thing

### consistency

- the design must not be inconsistent

- a design is allowed to be slightly less simple and less complete to avoid
  inconsistency

- consistency is as important as correctness

* * *

## the right thing

### completeness

- the design must cover as many important situations as is practical

- all reasonably expected cases must be covered

- simplicity is not allowed to overly reduce completeness

* * *

# worse is better

* * *

## worse is better

### simplicity

- the design must be simple, both in implementation and interface

- it is more important for the implementation to be simple than the interface

- simplicity is the most important consideration in a design

* * *

## worse is better

### correctness

- the design must be correct in all observable aspects

- it is slightly better to be simple than correct

* * *

## worse is better

### consistency

- the design must not be overly inconsistent

- consistency can be sacrificed for simplicity in some cases, but it is better
  to drop those parts of the design that deal with less common circumstances
  than to introduce either implementational complexity or inconsistency

* * *

## worse is better

### completeness

- the design must cover as many important situations as is practical

- all reasonably expected cases should be covered

- completeness can be sacrificed in favor of any other quality

- in fact, completeness must be sacrificed whenever implementation simplicity
  is jeopardized

- consistency can be sacrificed to achieve completeness if simplicity is
  retained; especially worthless is consistency of interface

* * *

The lesson to be learned from this is that it is often undesirable to go for
the right thing first. It is better to get half of the right thing available so
that it spreads like a virus. Once people are hooked on it, take the time to
improve it to 90% of the right thing.

* * *

Unix and C are the ultimate computer viruses.

[an essay]: https://dreamsongs.com/RiseOfWorseIsBetter.html
