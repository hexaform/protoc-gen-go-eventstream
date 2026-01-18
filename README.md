# protoc-gen-go-eventstream

`protoc-gen-go-eventstream` is a `protoc` plugin that generates Go helpers for **event-stream style Protobuf schemas**.

It is designed for systems where a single Kafka topic represents an **event stream for an aggregate**, and the topic schema is defined as a single Protobuf message containing a `oneof` of event messages.

The plugin removes the boilerplate and cognitive overhead introduced by Protobuf `oneof` code generation by eliminating the need to manually construct and unwrap the nested wrapper types generated for `oneof` fields.

## Problem

A common pattern for event streams is to define a topic schema like this:

```protobuf
message UserEvent {
  oneof event {
    UserRegistered user_registered = 1;
    UserBanned  user_banned  = 2;
    UserDeleted user_deleted = 3;
  }
}
```

This approach has clear benefits:

- one schema per topic
- explicit event types
- strong compatibility guarantees

However, the generated Go code is awkward to work with:

- events are wrapped in deeply nested `oneof` structs
- publishing requires manual construction of envelope messages
- consuming requires type switches over generated internal structures

This friction adds up quickly in event-driven systems.

## Solution

`protoc-gen-go-eventstream` treats the Protobuf message as an **event stream envelope** and generates Go helpers that:

- construct the envelope message from a concrete Protobuf event message (`oneof` member)
- unwrap the envelope back into the contained Protobuf event message
- hide all Protobuf `oneof` implementation details from application code
- allow producers and consumers to work directly with event message types, not envelope internals

The result is an API that feels natural in Go and aligns with event-driven and DDD-style architectures.

## What the Plugin Generates

Given an event stream message with a `oneof`, the plugin generates helper code that includes:

- constructors for wrapping event messages into the stream message envelope
- type-safe unwrapping helpers for consumers
- clear compile-time guarantees over supported event types

The exact API is generated from the schema and stays in sync with it.

## Example Usage

### Publishing

Instead of manually constructing the `oneof` envelope:

```go
msg := &pb.UserEvent{
  Event: &pb.UserEvent_UserBanned{
    UserBanned: &pb.UserBanned{
      UserId: 123,
      Reason: "violation",
    },
  },
}
```

You work with the concrete Protobuf event message (DTO) directly:

```go
dto := &pb.UserBanned{
  UserId: 123,
  Reason: "violation",
}

msg := pb.WrapEvent(dto)
```

### Consuming

Instead of switching over Protobuf-generated internals:

```go
switch wrapper := msg.GetEvent().(type) {
  case *pb.UserEvent_UserRegistered:
    event := wrapper.UserRegistered
    // ...
  case *pb.UserEvent_UserBanned:
    event := wrapper.UserBanned
    // ...
}
```

You unwrap the event cleanly:

```go
switch event := msg.UnwrapEvent().(type) {
  case *pb.UserRegistered:
  // ...
  case *pb.UserBanned:
  // ...
}
```

### Running code generation

The `go-eventstream` plugin is designed to work **alongside** the standard Go Protobuf generator. Both plugins must write their output into the **same Go package** so that the generated helper code can reference the generated Protobuf types directly.

If the outputs are split across different directories or packages, the generated code will not compile.

#### With protoc

When invoking `protoc` directly, make sure both plugins use the same output path and package options:

```bash
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-eventstream_out=. \
  --go-eventstream_opt=paths=source_relative \
  user-events.proto
```

Both `--go_out` and `--go-eventstream_out` point to the same destination, ensuring all generated files end up in the same package.

#### With Buf

When using Buf, the same rule applies: both plugins must emit code into the same output directory and therefore the same Go package.

```yaml
version: v2
plugins:
  - remote: buf.build/protocolbuffers/go
    out: generated/go
    opt: paths=source_relative
  - local: protoc-gen-go-eventstream
    out: generated/go
    opt: paths=source_relative
```

Using a shared output directory guarantees that the eventstream helpers and the Protobuf-generated types live in the same package and can work together without additional wiring or imports.

## When to Use This Plugin

Use `protoc-gen-go-eventstream` if:

- your Kafka topics represent aggregate-level event streams
- you follow a one-schema-per-topic policy
- you want to keep your code free of Protobuf internals
- you value explicit schemas but ergonomic Go APIs

## When Not to Use It

This plugin is probably not a good fit if:

- your topics carry unrelated message types
- you want to expose raw `oneof` structures intentionally