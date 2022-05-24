## iter

[![Go Report Card](https://goreportcard.com/badge/github.com/PartyLich/go/iter)](https://goreportcard.com/report/github.com/PartyLich/go/iter)
[![Go Reference](https://pkg.go.dev/badge/github.com/PartyLich/go/iter.svg)](https://pkg.go.dev/github.com/PartyLich/go/iter)

Golang generic Iterators, with an intentionally Rust-y flavor.

I have zero expectation that any other party will use this. There are very likely better and more
complete alternatives already available - I haven't looked, but if you're reading this I encourage
you to run a quick search.

This is a reimplementation the Rust iterator traits I use frequently.
These wouldn't be possible in quite the same form prior to the release of generics in Go, which is
quite new at the time of writing. Some useful generic slice related functionality can be found in
the 'experimental' [slices package](https://pkg.go.dev/golang.org/x/exp/slices), though not quite
the iterator tools I was missing.

The API as it's currently shaping up isn't quite what I hoped for. The [absence of parameterized
methods in the language](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods) is a bit of a stumbling block. If I come across a better way to deal with
it, I'll rewrite accordingly. Nonetheless, there is a fluent api for nearly all operations that do
not change the type of the underlying collection (eg Map is not fluent). Manual iteration is
comparable to the style found in [container/list](https://pkg.go.dev/container/list). An alternative
versio following the style of Google's [api/iterator](google.golang.org/api/iterator) is certainly
possible, but not currently an objective.

## License

SPDX-License-Identifier: MIT or Apache 2.0 license
