```
template = "slides"
title = "Worse is better"
date = "13th of April 2018"
```

# Worse is Better

* * *

## Context

- It's 1991.

- You are Richard P. Gabriel, computer scientist.

- You love <span class=sc>lisp</span>.

- Everyone else loves C.

* * *

# What gives?

* * *

- C is worse than <span class=sc>lisp</span>.

- But C won the race.

- Therefore worse is better.

- You write [an essay][] describing the worse is better approach.

* * *

# The Right Thing

* * *

## The Right Thing

### Simplicity

- The design must be simple, both in implementation and interface.

- It is more important for the interface to be simple than the implementation.

* * *

## The Right Thing

### Correctness

- The design must be correct in all observable aspects.

- Incorrectness is simply not allowed.

* * *

## The Right Thing

### Consistency

- The design must not be inconsistent.

- A design is allowed to be slightly less simple and less complete to avoid
  inconsistency.

- Consistency is as important as correctness.

* * *

## The Right Thing

### Completeness

- The design must cover as many important situations as is practical.

- All reasonably expected cases must be covered.

- Simplicity is not allowed to overly reduce completeness.

* * *

# Worse is Better

* * *

## Worse is Better

### Simplicity

- The design must be simple, both in implementation and interface.

- It is more important for the implementation to be simple than the interface.

- Simplicity is the most important consideration in a design.

* * *

## Worse is Better

### Correctness

- The design must be correct in all observable aspects.

- It is slightly better to be simple than correct.

* * *

## Worse is Better

### Consistency

- The design must not be overly inconsistent.

- Consistency can be sacrificed for simplicity in some cases, but it is better
  to drop those parts of the design that deal with less common circumstances
  than to introduce either implementational complexity or inconsistency.

* * *

## Worse is Better

### Completeness

- The design must cover as many important situations as is practical.

- All reasonably expected cases should be covered.

- Completeness can be sacrificed in favor of any other quality.

- In fact, completeness must be sacrificed whenever implementation simplicity
  is jeopardized.

- Consistency can be sacrificed to achieve completeness if simplicity is
  retained; especially worthless is consistency of interface.

* * *

The lesson to be learned from this is that it is often undesirable to go for
the right thing first. It is better to get half of the right thing available so
that it spreads like a virus. Once people are hooked on it, take the time to
improve it to 90% of the right thing.

* * *

<span class=sc>unix</span> and C are the ultimate computer viruses.

[an essay]: https://dreamsongs.com/RiseOfWorseIsBetter.html
