# Todo app v2

* Normal todo app.
* Create task, and assign member, and set attribute(i.g. close date, story point), and move some status (todo, do, done, cancel).

## State diagram

* At first, this app has simple state only.

```mermaid
stateDiagram-v2
    [*] --> backlog: create task
    backlog --> done: check complete
    backlog --> cancel: check task cancel
    done --> [*] 
    cancel --> backlog: checkout task cancel
    cancel --> [*]
```

## Use case

* For example, when Bob use this app.
* Scenario.  
  * Bob joined Agile project of 6 sprint.
  * He use to his task controls.


```mermaid
sequenceDiagram
    PM ->> Bob: give some tasks.
    Bob -->> App: add his tasks with sprint.
    Bob -->> App: check complete his tasks.
    App -->> Bob: none complete tasks list.
    Bob ->> PM: report his tasks state.
```

## Architecture

* [layard + DDD](https://qiita.com/tono-maron/items/345c433b86f74d314c8d)

## Code Guideline

https://github.com/knsh14/uber-style-guide-ja