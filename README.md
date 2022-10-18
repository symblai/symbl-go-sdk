# Symbl.ai Go SDK

The Symbl.ai Go SDK provides a convenient method to work with our APIs, from applications written in Go. A prescribed or opinionated set of interfaces, classes, and functions help you quickly bootstrap to using the Symbl.ai platform and unlock conversation intelligence. 

We are working diligently to support every aspect of Symbl.ai APIs. Currently, we support the following:
- [Streaming API][streaming_api-docs]:
  - WebSocket enabled
  - Easy to implement real-time language understanding
  - For local voice capture, provides an easy-to-understand library to enable microphone audio as an input source
- [Asynchronous APIs][async_api-docs]:
  - Transcription/Messages
  - Dynamic Topics
  - Questions
  - Follow-Ups
  - Entities
  - Action Items
  - Summary
  - Analytics
  - Trackers
- [Management APIs][management_api-docs]:
  - Entities: Get, Create, Delete
  - Bookmarks: Get, Create, Delete
  - Trackers: Get, Create, Delete

> **_IMPORTANT:_** This project is in pre-release status. Since this is the first release, we reserve the right to make breaking API changes at this time. 

## Documentation

See the [API docs][symbl-docs].

### Requirements

The minimal Go version supported is 1.18. Realistically, much older versions will work, but let’s start with that first as we launch this project.

### Installation

First, make sure that Go is installed on your system with the following command in Mac/Linux console or Windows command prompt:

```sh
$ go version
go version go1.18.4 darwin/arm64
```

To install Go, visit the [installation guide](https://go.dev/doc/install) which covers Linux, Mac and Windows.

### Configuration

The SDK needs to be initialized with your account's credentials `APP_ID` and `APP_SECRET`, which are available in your [Symbl.ai Platform][api-keys]. If you don't have a Symbl.ai Platform account, you can [sign up here][symbl_signup] for free.

You must add your `APP_ID` and `APP_SECRET` to your list of environment variables. We use environment variables because they are easy to configure, support PaaS-style deployments, and work well in containerized environments like Docker and Kubernetes.

```sh
export APP_ID=YOUR-APP-ID-HERE
export APP_SECRET=YOUR-APP-SECRET-HERE
```

## Examples

You can find a list of very simple main-style examples to consume this SDK in the [examples folder][examples-folder].

## Community

If you have any questions, feel free to contact us at devrelations@symbl.ai or through our [Community Slack][slack].

This SDK is actively developed, and we love to hear from you! Please feel free to [create an issue][issues] or [open a pull request][pulls] with your questions, comments, suggestions, and feedback. If you liked our integration guide, please star our repo!

This library is released under the [MIT License][license]


[api-keys]: https://platform.symbl.ai/#/login
[symbl-docs]: https://docs.symbl.ai/docs
[streaming_api-docs]: https://docs.symbl.ai/docs/streaming-api
[async_api-docs]: https://docs.symbl.ai/docs/async-api
[management_api-docs]: https://docs.symbl.ai/docs/management-api
[symbl_signup]: https://platform.symbl.ai/signup?utm_source=symbl&utm_medium=blog&utm_campaign=devrel&_ga=2.226597914.683175584.1662998385-1953371422.1659457591&_gl=1*mm3foy*_ga*MTk1MzM3MTQyMi4xNjU5NDU3NTkx*_ga_FN4MP7CES4*MTY2MzEwNDQyNi44Mi4xLjE2NjMxMDQ0MzcuMC4wLjA.
[examples-folder]: examples/
[issues]: https://github.com/dvonthenen/symbl-go-sdk/issues
[pulls]: https://github.com/dvonthenen/symbl-go-sdk/pulls
[license]: LICENSE
[slack]: https://join.slack.com/t/symbldotai/shared_invite/zt-4sic2s11-D3x496pll8UHSJ89cm78CA
