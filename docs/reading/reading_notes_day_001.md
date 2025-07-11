# CLI Tools and Go Lifecycle Tools and Libraries

## [Cobra](https://github.com/spf13/cobra)
- CLI application building library in go
    - used by git, hugo, kubernetes

### Features
    - sub command based CLIs (app server, app fetch etc.)
    - POSIX and Go flag library complaint
    - Nested subcommand support
    - Cascading flags that move down to children if needed
    - Auto complete for
        - flag recoginiton of -h and --help
        - shell autocomplete for bash, zsh, fish, psh
        - man page generation
        - intelligent suggestions
    - command aliases

### Core Concepts
- Built on 3 atomics - Commands, Args and Flags
- iteractions flow like APPNAME COMMAND ARG --FLAG
    - APPNAME - name of the application
    - COMMAND - what funcitonality of the app is being used
    - ARG - a noun that is the subject of the command (url for example)
    - FLAG - modifies the command
    - example: git clone URL --bare

## [Context - Concurrency Patterns in Go](https://go.dev/blog/context)

### Reasoning
- Usual request handling pattern in go is a per-request-goroutine handling.
- Each handler then might start its own goroutines to fetch services from
  the backend like DB and RPC services.
- For each per-request-goroutine and its children, client specific data is needed
  to be kept like
    - request deadline
    - authentication token
    - identity of the user
- All nested goroutines handling a single request must stop executing on cancellation
  or timeout and exit quickly
- The context package in go helps manage and pass request scoped values, cancellation, signals
  and deadlines

### Context type

type Context interface {

    Done() <-chan struct{}
    Err() error
    Deadline() (deadline time.Time, ok bool)
    Value(key interface{}) interface{}
}

- Done() is a function that returns a channel
    - this channel is closed via close(chan_name) when the context is cancelled
    - functions implementing the Context interface should loop on this channel
      check if its closed and exit immediately

- Err() returns an error describing why the context was cancelled

- Context doesnt have a cancel method
    - the function that recieves a cancel singal doesnt usually send one
    - this is the same reason why the Done() returned chan is recieve only
    - Cancellations should come from the parent or root

- Contexts can be safely passed to any number of goroutines and a cancel signal will
  be passed to all of them

- Deadline gives us the deadline time which the function can check before starting work
    - to see if it has time to complete before the deadline
    - useful when setting timeouts for I/O operations

- Value is used to carry the request specific data (like user data) which we were
  talking about earlier
    - key passed must be concurrency safe to be used by multiple goroutines

### Derive contexts
- we can derive new Context values from exisiting ones
- all contexts derived from a context are cancelled on cancellation of the original context
- Background is the the root context of all contexts and is never cancelled

- CancelFunc, WithCancel, WithTimeout and WithValue
   - func WithCancel(parent Context)  (ctx Context, cancel CancelFunc)
   - func WithTimeout(parent Context, timeout time.Duration)  (ctx Context, cancel CancelFunc)
   - func WithValue(parent Context, key interface{}, val interface{}) Context
   - type CancelFunc func()

    - The CancelFunc is used to cancel a context

    - WithCancel creates a copy of the parent whose Done is closed when the parent.Done is closed
      or cancel is called

    - WithTimeout is same as WithCancel but also closes done on timeout
        - deadline of the WithTimeout is the lesser of now+timeout or parents deadline
        - cancel releases all resources if the timer is still going
    - WithValue copy of the parent whose Value returns val for key

### Closing Comments
- the context is the first argument for every func between incoming and outgoing

## [Graceful shutdown in Go](https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/)

### Reasoning
- Stateful long lasting processes need to gracefully shutdown for obvious reasons
    - DB would lose unflushed writes
    - Servers would crash without responding properly to closed connections
- Graceful Shutdown consists of
    - close all pending processes
    - dont start any new processes
    - dont accept any requests
    - close all open connections to external services
- Problems to solve
    - When do we shutdown?
    - How we communicate the shutdown to processes (especially async tasks)
- Goals of shutdown
    - No data loss during a shutdown
    - better service control (integration testing)

### Anti Patterns and Solutions
- Block artificially

    Problem:
    - block main goroutine busy waiting on nothing
        - var ch chan int
          <-ch

    Solution:
    - control concurrency and wait using

      - Channel method
        - make(chan struct{}, 1) // make empty channel
        - every child publishes to the channel when done
        - parent consumes from the channel as many times
          as the number of goroutines created
        - useful while waiting on a single go routine

      - WaitGroup
        - sync.WaitGroup()
        - add goroutines to WaitGroup
            wg.add(1)
                go func() {
                    defer wg.Done()
                    for {
                        select {
                            case <-ctx.Done():
                                // do some
                            case <-time.After(i * time.Second):
                                // do some
                        }
                     }
                 }
            wg.Wait()
        - slightly better for multiple goroutines

       - errgroup
        - sync/errgroup
        - Go and Wait methods

          g.gCtx := errgroup.WithContext(ctx)
            g.Go(func() error{});
          g.Wait()


        -



- os.Exit()

    Problem:
    - calling os.Exit() when other goroutines are still running
    - this prematurely terminates open connections and inflight requests

    Solution:
    - None blocking signal handling
        - c := make(chan os.Signal, 1)
        signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
        - use this channel to get signals from SIGINT and SIGTERM
        - select used to consume from multiple channels
          select {
              case <-c:
                  // handle signals

              case <-time.After(i * time.Second):
                  // handle timeout
          }

- Problem with signal handling
    - channels for signals dont work for broadcast
    - only one of the channels will return when a signal arrives
    - in the case of wg where we add multiple goroutines and select on a signal channel
      it wont work out for us since only one of the goroutines will do the signals exit case

- Solution to this
    - use context.WithCancel
    - creates a context and returns a CancelFunc method
    - we can use the cancel to send a signal to every ctx.Done
    - this functionality of signals and context was combined in

      func NotifyContext(parent context.Context, signal ...os.Signal)
        (ctx context.Context, stop context.CancelFunc)  {}
    this returns a context that sends a signal to the Done channel when the signal arrives

### Common libraries
- HTTP Server
    - Problems on ungraceful shutdown:
        - never get a response
        - waste of resources due to incomplete response
        - data inconsitencies

    - Cloud native enviorments have shutdowns often with horizontal scaling
      of pods and containers
    - graceful shutdown of http servers is very important
        - httpServer.Shutdown(contxt.Background()) after waiting on gCtx.Done()
        - add this to a goroutine which is part of an error group
        - closes all open listeners -> close all idle connections -> wait for
          connections to return to idle then shutdown
        - we pass it a context so if the context expires before the shutdown it will return
          the context error otherwise close any error returned in the listener closing process

    - this approach is fine for most cases but for long running connections like websocket
      we could be waiting a long time for it to shutdown gracefully
      - to avoid this we get the context from the http request using req.Context()
      - use the BaseContext paramter of http.Server
        - BaseContext: func(_ net.Listener) context.Context {

                return mainCtx;
        }
        - make it return the main functions context this makes the server
          use the main ctx for every response
- HTTP Client
    - we can pass a context with a request using NewRequestWithContext
    - if we use this we have to follow certain techniques in certain cases
        - Draining Woker Channels
            - once the process shutsdown we have to make sure no messages
              are left in the channels
            - writing to a closed channel will result in a panic
            - if multiple workers are writing close the channel only after all
              workers publish on the channel
            - if reading exit after the channel has no more data
            - writer stops the readers by stopping the writing process and closing

## Graceful Methods
- ways to expose methods that can be used to facilitate graceful shutdown
- Blocking with ctx
    - call a method
    - pass the context
    - method blocks
    - returns if error or if cancelled/timeout on context
- Setup/Shutdown
    - if you cant block we can an expose a Shutdown() method
    - the setup method creates a context with cancel
    - the setup method creates a errgroup with the context that is used to wait on
    - we store the cancel which is then called when wecall our Shutdown method



