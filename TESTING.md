# TESTING

## `dungeon`

* `dungeon.go` - `dungeon-test.go`
  Run dungeon generation tests with `go test -v ./dungeon`. The idiomatic way
  to disable test result caching - and getting a new unique dungeon each run -
  add `-count=1` to the command line.
* `tilemap.go` - `tilemap_test.go`
  Run tilemap parsing tests with `go test -v ./tilemap`
