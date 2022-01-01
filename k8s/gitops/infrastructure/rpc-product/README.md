# posapi

`posapi` is the API for the point of sale as well as the bridge webapp and support webapp.  It is written in Golang and provides a secure WebSocket API interface to clients.  It depends on the `authz` service and communicates via gRPC to `authz`.  It uses `PostgreSQL` as its persistent store.

* [Homepage](https://github.com/nabancard/pos-api)
* Releases are currently done via CI/CD anytime we push a successful build, then deployed by flux.

**NOTE** This image should have a flux `ImageRepository` and `ImagePolicy`.
