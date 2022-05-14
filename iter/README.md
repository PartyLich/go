Golang generic Iterators, with an intentionally Rust-y flavor.

I have zero expectation that any other party will use this. There are very likely better and more
complete alternatives already available - I haven't looked, but if you're reading this I encourage
you to run a quick search.

While practicing a bit of Go I found myself rather surprised at the some of the things that are
(intentionally, afaik) absent from the standard libraries. Like a Math.Max function for the default
inferred number type, int. So, like (presumably) countless others before me, I found myself
reinventing these otherwise common wheels.

Shortly thereafter I found myself reimplementing some of the Rust iterator traits I use frequently.
These wouldn't be possible in quite the same form prior to the release of generics in Go, which is
quite new at the time of writing. Some useful generic slice related functionality can be found in
the 'experimental' [slices package](https://pkg.go.dev/golang.org/x/exp/slices), though not quite
the iterator tools I was missing.

The API as it's currently shaping up isn't quite what I hoped for. The [absence of parameterized
methods in the language](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods) is a bit of a stumbling block. If I come across a better way to deal with
it, I'll rewrite accordingly.


MIT or Apache 2.0 license
