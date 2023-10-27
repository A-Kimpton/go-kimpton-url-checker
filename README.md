# URL Checker

This is a dummy project that contains two buildable binaries:
- Job Dispatcher: publishes jobs to a message queue
- Job Processor: consumes jobs from a message queue

The primary focus of this project was to play with RabbitMQ and message queues in general.  Eventually I would like both binaries to be horizontally scalable, and the dispatcher can receive batch requests via an endpoint.

## Architecture

![Label](.github/images/architecture.drawio.svg)
Internally, the design is very basic
- There is an abstraction layer ontop of rabbitmq so I can easily swap out implementation
- There is a message queue worker, that listens to a single queue name to consume tasks
- Workers and tasks are written in a way so that it is easy to expand on.