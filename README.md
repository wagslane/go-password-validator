# go-password-validator
Validate the Strength of a Password in Go

[![](https://godoc.org/github.com/lane-c-wagner/go-password-validator?status.svg)](https://godoc.org/github.com/lane-c-wagner/go-password-validator)

This project can be used to front a password strength meter, or simplay validate password strength on the server. Benefits:

* No stupid rules (use uppercase, use special characters, etc)
* Everything is based on entropy (raw cryptographic strength of the password)
* Inspired by this [XKCD](https://xkcd.com/936/)

![XKCD Passwords](https://imgs.xkcd.com/comics/password_strength.png)

## ⚙️ Installation

Outside of a Go module:

```bash
go get github.com/lane-c-wagner/go-password-validator
```

## 🚀 Quick Start

```go
package main

import (
    passwordvalidator "github.com/lane-c-wagner/go-password-validator"
)

func main(){
    entropy := passwordvalidator.GetEntropy("a longer password")
    // entropy is a float64, representing the strength in base 2 (bits)
    
    err := passwordvalidator.Validate("some password", 60)
    // if the password has enough entropy, err is nil
    // otherwise, a formatted error message is provided explaining
    // how to increase the strength of the password
    // (safe to show to the client)
}
```

## What Entropy Value Should I Use?

It's up to you. That said, here is a pretty good graph that shows some timings for differnt values:

![entropy](https://external-preview.redd.it/rhdADIZYXJM2FxqNf6UOFqU5ar0VX3fayLFpKspN8uI.png?auto=webp&s=9c142ebb37ed4c39fb6268c1e4f6dc529dcb4282)

Somewhere in the 50-70 range seems "average"

## How It Works

First we determines the "base" number. The base is a sum of the different "character sets" found in the password.

The current character sets include:

* 26 lowercase letters
* 26 uppercase
* 10 digits
* 32 special characters - ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~`

Using at least one character from each set your base number will be 94: `26+26+10+32 = 94`

Every unique character that doesn't match one of those sets will add `1` to the base.

If you only use, for example, lowercase letters and numbers, your base will be 36: `26+10 = 36`.

After we have calculated a base, the total number of brute-force-guesses is found using the following formulae: `base^length`

A password using base 26 with 7 characters would require `26^7`, or `8031810176` guesses.

### Additional Safety

To add further safety to dumb passwords like aaaaaaaaaaaaa, or 123123123, We modify the length of the password to count any more than two of the same character as 0.

* `aaaa` has length 2
* `12121234` has length 6

## 💬 Contact

[![Twitter Follow](https://img.shields.io/twitter/follow/wagslane.svg?label=Follow%20Wagslane&style=social)](https://twitter.com/intent/follow?screen_name=wagslane)

Submit an issue (above in the issues tab)

## Transient Dependencies

None! And it will stay that way, except of course for the standard library.

## 👏 Contributing

I love help! Contribute by forking the repo and opening pull requests. Please ensure that your code passes the existing tests and linting, and write tests to test your changes if applicable.

All pull requests should be submitted to the "master" branch.

```bash
go test
```

```bash
go fmt
```
