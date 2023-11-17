# Channels

Run by running

```
make run
```

and

```
# go run 06_channels/main.go
```

or simply

```
# docker exec -it learning_go go run 06_channels/main.go
```

We want to build a program that check the status of a set of websites.

![building](./images/building.png)

In a for loop we will just be waiting for each request to finish before starting another one.

![first_approach](./images/first_approach.png)

Instead, we can use go routines to make each request in parallel.

![second_approach](./images/second_approach.png)

## Concurrency is not parallelism

![theory](./images/theory.png)

**Concurrency**: run multiple go routines at the same time.

![concurrency](./images/concurrency.png)

**Parallelism**: requires multiple CPUs.

![parallelism](./images/parallelism.png)

When the main routine finishes before the child routines, the program quits automatically.

![bug](./images/bug.png)

To avoid this we will use channels, to communicate between go routines. They need to be typed specifying the type of information that needs to be sent through them.

![channels](./images/channels.png)

There are several ways to communicate with a channel explained below.

![communicate_channels](./images/communicate_with_channel.png)

A third approach will be done to continuously request the status of the urls.

![third_approach](./images/third_approach.png)
