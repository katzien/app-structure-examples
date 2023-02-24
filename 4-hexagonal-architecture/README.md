# Hexagonal architecture

We arrive at one of the well-known ways of grouping functionality together - Hexagonal Architecture, also known as Ports & Adapters or the onion architecture. 🙂

This is my favourite way of structuring Go apps, especially when coupled with a microservice architecture.
I find that it scales well to problems of different sizes and complexity, and you can add or remove as many layers as you need:

```
├── main.go
├── dao
├── domain
└── handler
```

You can also combine many smaller hex structures to form your entire program:

(imagine an app which compiles a restaurant menu by pulling information from three completely different sources)

```
├── main.go
├── beer
│ ├── dao
│ ├── domain
│ ├── handler
│ └── service.go
├── wine
│ ├── dao
│ ├── domain
│ ├── handler
│ └── service.go
└── food
  ├── dao
  ├── domain
  ├── handler
  └── service.go
```

This structure is great for:
- complex projects which require a well-defined structure to stay maintainable and understandable
- applications which have dependencies such as databases, user interfaces or third party APIs
