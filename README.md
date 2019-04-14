# ab-go excercise

**ab-go excercise** is a Coderschool golang course prework

Submitted by: **hiepndd**

Time spent: **2** hours spent in total

## Usage

```bash
-timeout 3 -timelimit 10  -n 9 -c 3   http://coderschool.vn/
```

![](http://g.recordit.co/A3b23KcuB8.gif)

## Set of User Stories

### Required

- [x] Command-line argument parsing
- [x] Input params
  - [x] Requests - Number of requests to perform
  - [x] Concurrency - Number of multiple requests to make at a time
  - [x] URL - The URL for testing.
- [x] Prints use information if wrong arguments provided
- [x] Implements HTTP load and summarize it
- [x] Concurrency must be implemented with goroutine.

### Bonus

- [ ] Extend input params with:
  - [x] Timeout - Seconds to max. wait for each response
  - [x] Timelimit - Maximum number of seconds to spend for benchmarking
- [ ] Prints key metrics of summary, such:
  - [ ] Server Hostname
  - [ ] Server Port
  - [ ] Document Path
  - [ ] Document Length
  - [ ] Concurrency Level
  - [x] Time taken for tests
  - [x] Complete requests
  - [x] Failed requests (count request based on time out)
  - [ ] Total transferred
  - [ ] Requests per second
  - [x] Time per request
  - [ ] Transfer rate
