# hello-go
Testing various new features in versions of [Go](https://go.dev/).

## Running
 - `docker-compose up` to start containers for versions

## Versions
 - [go1.24](https://go.dev/doc/go1.24)
   - [os.Root isolates filesystem operations to a single directory parent](https://go.dev/blog/go1.24#standard-library-additions) :bricks:
 - [go1.23](https://go.dev/doc/go1.23)
   - [slices package now has Sort and other functions that work with iterators](https://go.dev/doc/go1.23#iterators) :twisted_rightwards_arrows:
 - [go1.22](https://go.dev/doc/go1.22)
   - [for loop iterations no longer share variables](https://go.dev/blog/loopvar-preview) :tada:
 - [go1.16](https://go.dev/doc/go1.16)
    - static files can now be embedded into the executable
    - `io/fs` allows for traversing a tree of files
