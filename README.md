# yt-transcription-translation

---

**Current state: Work in progress, not yet usable.**

---

## Goals/Planned Features

A stateless microservice to download, transcribe and translate YouTube videos
that can be deployed to a serverless platform.

It will provide a few simple endpoints that can be used to:

- start a download, transcription and translation job.
- get the status/results of a job.

- The server/backend will be built using stdlib and no other routers.
  Some libraries that will be used by the backend include:
  - [youtube/v2](github.com/kkdai/youtube/v2) for downloading the video.
  - [zap](go.uber.org/zap) for logging.

This project could easily be extended to be production ready by adding a
database to store jobs and their results, etc.

## Planned Usage

- Can be deployed as a microservice.
- Simple CLI to use locally.
