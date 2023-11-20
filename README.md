# yt-transcription-translation

---

**Current state: Work in progress, not yet usable.**

---

## Goals/Planned Features

- Use a command line tool with a simple interface to download a YouTube video,
  get a transcription and a translation for the downloaded video.
- Two components in this project: cli/frontend and server/backend.
- The server/backend will be built using stdlib and no other routers.
  Some libraries that will be used by the backend include:
  - [end](https://github.com/ent/ent) for ORM.
  - [youtube/v2](github.com/kkdai/youtube/v2) for downloading the video.
  - [go-sqlite3](github.com/mattn/go-sqlite3) for a `sqllite3` database driver.
  - [zap](go.uber.org/zap) for logging.
- The cli/frontend will be written using:
  - [cobra](https://github.com/spf13/cobra)
  - [bubbletea](https://github.com/charmbracelet/bubbletea).

The backend decision are cemented, but the frontend/cli is still up for debate.
This project isn't necessarily meant to be a production-ready tool,
but rather a learning experience for me to learn more about Go
and how to build a full-stack application.
