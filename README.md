# pubsub

Publish/Subscribe implementation written in go.

## Example

```go
func main() {
    server := NewPubSub()
    sub := server.Subscribe("news")
    server.Publish("what is up?", "news")
    fmt.Println(<-sub) // what is up
    server.AddSubscription(sub, "tech")
    server.Publish("tired of side projects? call me!", "tech")
    fmt.Println(<-sub) // tired of side projects? call me!
}
```