<div align="center">
  <h1>BRHT</h1>
  <p>Brain Rot hEcker Terminal</p>
  <i>A fake, fully harmless Hollywood hacker terminal, somehow infected with TikTok brainrot.</i>
  <br/><br/>
</div>

![asciicast](docs/brht.gif)

BRHT looks like every "I'm in" hacker scene you've ever seen in a movie: green text,
scrolling logs, a dozen progress bars, a suspiciously confident AI. Except every
single thing on screen is procedurally generated nonsense, remixed with skibidi,
sigma, Ohio, and whatever else the internet is doing this week.

## Tech stack

- **Language:** Go
- **Terminal UI:** [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **Styling:** [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- **Widgets:** [Bubbles](https://github.com/charmbracelet/bubbles)

## ⚠️ Disclaimer

**BRHT performs zero real hacking.** No networking, no scanning, no exploitation, no
packet inspection, no brute forcing, no anything that touches a real system other
than your own terminal. Every IP address, log line, progress bar,
and stat is randomly generated text for entertainment purposes only. If it looks
scary, that's the joke.

## Installation

Download a prebuilt binary from the [Releases](../../releases) page for your platform
(Linux amd64/arm64, macOS Intel/Apple Silicon, or Windows x64), then run it:

```sh
chmod +x brht        # macOS / Linux only
./brht
```

Or install it with Go directly:

```sh
go install github.com/s4nj1th/brht/cmd/brht@latest
```

## Building

Requires Go 1.22+.

```sh
git clone https://github.com/s4nj1th/brht.git
cd brht
go build -o brht ./cmd/brht
./brht
```

Run the test suite:

```sh
go test ./...
```

Cross-compile for another platform:

```sh
GOOS=linux GOARCH=arm64 go build -o brht-linux-arm64 ./cmd/brht
```

A full release build for every supported platform is handled by
[GoReleaser](https://goreleaser.com) — see [`.goreleaser.yaml`](./.goreleaser.yaml) and
[`.github/workflows/release.yml`](./.github/workflows/release.yml).

## Keyboard shortcuts

- **Any key:** Generates more chaos, including logs, packets, progress, and brainrot.
- **Ctrl+C:** Exits immediately without confirmation.

There is no text input anywhere in BRHT. Every keystroke is just fuel for the
simulation. It all gets converted into fake logs and nonsense progress bars and
nothing else.

## How it works

BRHT is driven by a simple simulation engine:

- An **engine** holds all simulated state: logs, progress bars, stats, fake
  network packets, AI chatter, rare alerts, and glitches.
- A **tick** advances that state by one step. Idle ticks fire on a timer so the
  screen never goes still. Keypress ticks fire on every key and inject a bigger,
  showier burst of activity.
- Every panel (`SYSTEM`, `MAIN TERMINAL`, `NETWORK`, `AI`, `PROGRESS`, `STATS`,
  `LOGS`) is an isolated component that only reads from the shared engine state.
- All flavour text, including fake commands, brainrot phrases, log messages, and
  progress bar labels, lives in `internal/data`.

## Contributing

Contributions, new brainrot phrases, and increasingly unhinged progress bar labels
are all welcome.

1. Fork the repo and create a branch off `main`.
2. Keep it dependency-light (Bubble Tea / Lip Gloss / Bubbles only).
3. Run `gofmt -l .`, `go vet ./...`, and `go test ./...` before opening a PR.
4. Open a PR.

When adding new brainrot content, please add it to `internal/data/data.go`.

## License

[MIT](./LICENSE). Do whatever you want with it.
