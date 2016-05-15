Go Wrappers for Sandstorm's API

Right now just the output of go-capnpc for both the stuff that ships
with capnproto and all of the sandstorm schemas.

* `capnp/capnp` contains the base capnproto schema
* `capnp/sandstorm` contains the sandstorm APIs.

Since the schema names are not all legal package names, the following
changes have been made:

* Schema with dashes in their names have had the dashes removed.
* The schema `package` has been mapped to the package `spk` (since
  `package` is a go reserved word).
* The schema `c++` has been mapped to the package `cxx`. Direct use of
  this is unlikely, but some of the other schema include it, so it needs
  to be present.

# Licensing

Everything is licensed under the MIT license (Both my work and that of
upstream). The `*.capnp` files are the output of `gen.py`, which just
copies them from upstream and appends the lines beginning with `$Go`.
The copyright notices have been retained. The `*.capnp.go` files are
generated by `go-capnpc`, and should be thought of as object code; they
are committed to the repository to make the library usable with `go
get`.