# Streaming API (Real-Time) Example

This example uses the Microphone as input in order to detect conversation insights in what is being said. This example required additional components (for the microphone) to be installed in order for this example to function correctly. 

## Configuration

The SDK needs to be initialized with your account's credentials `APP_ID` and `APP_SECRET`, which are available in your [Symbl.ai Platform][api-keys]. If you don't have a Symbl.ai Platform account, you can [sign up here][symbl_signup] for free.

You must add your `APP_ID` and `APP_SECRET` to your list of environment variables. We use environment variables because they are easy to configure, support PaaS-style deployments, and work well in containerized environments like Docker and Kubernetes.

```sh
export APP_ID=YOUR-APP-ID-HERE
export APP_SECRET=YOUR-APP-SECRET-HERE
```

## Installation

The Streaming API (Real-Time) example makes use of a [microphone package](https://github.com/dvonthenen/symbl-go-sdk/tree/main/pkg/audio/microphone) contained within the repository. That package makes use of the [PortAudio library](http://www.portaudio.com/) which is a cross-platform open source audio library. If you are on Linux, you can install this library using whatever package manager is available (yum, apt, etc.) on your operating system. If you are on macOS, you can install this library using [brew](https://brew.sh/).
