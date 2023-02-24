# All examples implement a solution to the Advent of Code 2022 Day 1 puzzle. 

Here's the puzzle description (copied from [Advent of Code 2022](https://adventofcode.com/2022/day/1)):

--

The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

```
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
```

This list represents the Calories of the food carried by five Elves:

- The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
- The second Elf is carrying one food item with 4000 Calories.
- The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
- The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
- The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

# FAQ
### Where would you put your Kubernetes manifests / Dockerfiles / database migrations / non-Go files?
I usually add them to the root directory of the project, or group them in a top-level directory. E.g.
```
├── main.go
├── README.md
├── .gitignore
├── config (e.g. database migration or schema files)
├── consumer (if event-driven)
├── dao
├── domain
├── handler
├── manifests
│ ...
```

### Do you use internal/external packages to indicate what's safe to import in other projects?
Not really. If anything I would only add an `external` directory to the root of my project for this purpose (everything else being internal by default). So not:
```
├── main.go
├── internal
│ ├── dao
│ ├── domain
│ ├── handler
│ └── service.go
├── external
│ └── types.go
```

But rather:
```
├── main.go
├── dao
├── domain
├── handler
└── service.go
├── external
│ └── types.go
```

I just like to keep things simple, and stick to at most one level of indentation. You lose that if your internal/external packages contain subdirectories. 

I also don't really like the name `external`, because Go packages are typically "things" (not adjectives).

At work we actually use a `proto` package for this purpose. We use protobuf for microservice to microservice communication, and have an explicit rule that services can import only proto packages of other services (enforced by static analysis). This achieves the intent behind internal/external, but conveys a bit more meaning - `proto` contains protobuf definitions and constants which other services can use to interact with that service.

### Do you use cmd / pkg directories? 
No, and this is something I used to advocate in my talks on structuring Go apps!

I don't use these by default anymore. I realised they make the structure more complicated without bringing much benefit, unless I needed to have multiple main.go files. But that isn't the case for me most of the time, so I prefer to have my main.go file directly at the root of the project.

Pkg adds another layer of indentation for the rest of the packages, and I prefer to have a flat structure as much as possible to keep things simple. It doesn't add much benefit unless you prefer to clearly separate your Go code from any other non-Go files or directories. I guess it might help if your project tree is huge, but I tend to use microservices so rarely run into this problem. ☺️  