sms
===

### What is sms?

A simple Golang sms library. Use email to send a sms. inspired from [sms-fu](https://github.com/brendanlim/sms-fu).

### Install

`go get github.com/acmacalister/sms`

### Usage

sms is really simple. Check out the example below or have a look at the test. *Note* for the test to run you need to update the `sms.yml` with your own email account's username/password and  a phone number to successfully send a sms.

```go
package main

import (
  "github.com/acmacalister/sms"
  "log"
)

func main() {
  client, err := createClient("smtp.gmail.com", 587, "tester@gmail.com", "test")
  if err != nil {
    log.Fatal(err)
  }
  if err := client.Deliver("5555555555", "AT&T", "sms golang library!"); err != nil {
    log.Fatal(err)
  }
}
```

### License

sms is licensed under the Apache License.

### Authors

* [acmacalister](http://twitter.com/acmacalister)
* you

### ToDo

- [ ] Test other carriers. (Currently only AT&T has been tested.)
- [ ] Bring test coverage from 75% to 100%.
