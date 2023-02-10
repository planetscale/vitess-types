# Vitess types

These types are extracted out of [Vitess source](https://github.com/vitessio/vitess) and only modified to be more friendly for others.

## Why?

The main interface, besides MySQL, for Vitess is via gRPC. gRPC uses protobuf. If a third party module wanted to interact with Vitess's gRPC APIs, or to just work with their types such as a `QueryResult`, the logical choices is to use the generated interfaces from `vitess.io/vitess/go/vt/proto`, but this dependency is quite large since this is the entirety of Vitess.

So this extraction is an attempt at being a much lighter-weight version that is fully capable of interoping with Vitess. This makes it more friendly for consumers to wrap and depend on Vitess types, without a dependency on Vitess.
