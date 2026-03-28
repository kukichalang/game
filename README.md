# kukichalang/game

A beginner-friendly 2D game library for [Kukicha](https://github.com/kukichalang/kukicha), wrapping [Ebitengine](https://ebitengine.org/) with a simple, pipe-friendly API. Build browser games with WebAssembly — no game engine experience required.

## Install

Requires the [Kukicha compiler](https://github.com/kukichalang/kukicha). The game package is part of Kukicha's standard library — just import it:

```kukicha
import "stdlib/game"
```

## Quick Start

```kukicha
import "stdlib/game"

function main()
    _ = game.Window("My Game", 640, 480)
        |> game.OnDraw(draw)
        |> game.Run() onerr panic "{error}"

function draw(screen game.Screen)
    game.Clear(screen, game.MakeColor(30, 30, 60, 255))
    game.DrawText(screen, "Hello, Kukicha!", 240, 220, game.White)
```

Build for the browser:

```bash
kukicha build --wasm my-game.kuki
# Produces: my-game.wasm, wasm_exec.js, index.html
```

## API

### Window & Lifecycle

| Function | Description |
|----------|-------------|
| `game.Window(title, width, height)` | Create a game window |
| `game.OnSetup(app, fn)` | Register a one-time setup callback |
| `game.OnUpdate(app, fn)` | Register a per-frame update callback |
| `game.OnDraw(app, fn)` | Register a per-frame draw callback |
| `game.Run(app)` | Start the game loop |

All builder functions return `App` and are designed for pipe chains:

```kukicha
_ = game.Window("Breakout", 800, 600)
    |> game.OnSetup(setup)
    |> game.OnUpdate(update)
    |> game.OnDraw(draw)
    |> game.Run() onerr panic "{error}"
```

### Drawing

| Function | Description |
|----------|-------------|
| `game.Clear(screen, color)` | Fill the screen with a solid color |
| `game.DrawRect(screen, x, y, w, h, color)` | Draw a filled rectangle |
| `game.DrawCircle(screen, x, y, radius, color)` | Draw a filled circle |
| `game.DrawLine(screen, x1, y1, x2, y2, color)` | Draw a line between two points |
| `game.DrawText(screen, text, x, y, color)` | Draw debug text |

### Input

| Function | Description |
|----------|-------------|
| `game.IsKeyDown(key)` | `true` if the key is currently held |
| `game.IsKeyPressed(key)` | `true` if the key was just pressed this frame |
| `game.MousePosition()` | Current cursor position `(float64, float64)` |
| `game.MouseClicked()` | `true` if left mouse button was just clicked |

Key constants: `KeyLeft`, `KeyRight`, `KeyUp`, `KeyDown`, `KeySpace`, `KeyEnter`, `KeyEscape`

### Collision Detection

| Function | Description |
|----------|-------------|
| `game.Overlaps(a, b)` | Two rectangles overlap |
| `game.OverlapsCircle(a, b)` | Two circles overlap |
| `game.CircleOverlapsRect(c, r)` | Circle and rectangle overlap |

### Utilities

| Function | Description |
|----------|-------------|
| `game.MakeColor(r, g, b, a)` | Create a color from RGBA (0-255) |
| `game.Random(min, max)` | Random int in [min, max) |
| `game.RandomFloat(min, max)` | Random float64 in [min, max) |
| `game.FrameCount()` | Frames elapsed since game start |

### Types

- **`Color`** — RGBA color (R, G, B, A int)
- **`Position`** — X, Y coordinate (float64)
- **`Size`** — Width, Height (float64)
- **`Rect`** — X, Y, Width, Height (float64) — for drawing and collision
- **`Circle`** — X, Y, Radius (float64) — for drawing and collision
- **`Screen`** — Draw target (passed to your draw callback)
- **`App`** — Game configuration (built via `Window` + builder functions)

### Color Constants

`Red`, `Green`, `Blue`, `White`, `Black`, `Yellow`, `Orange`, `Purple`, `Gray`

## Tutorials

The Kukicha repo includes an 8-lesson game tutorial series that builds up to a full Breakout clone:

1. [Hello World](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/01-hello-world.md) — Window, text, colors
2. [Drawing Shapes](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/02-drawing-shapes.md) — Rectangles, circles, lines
3. [Keyboard Input](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/03-keyboard-input.md) — Moving objects with keys
4. [Animation](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/04-animation.md) — Frame-based movement
5. [Collision](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/05-collision.md) — Detecting overlaps
6. [Score & State](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/06-score-and-state.md) — Game state management
7. [Sound & Setup](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/07-sound-and-setup.md) — Audio and initialization
8. [Breakout](https://github.com/kukichalang/kukicha/blob/main/docs/tutorials/game/08-breakout.md) — Full game project

## Architecture

This package is a separate Go module (`github.com/kukichalang/game`) that wraps Ebitengine. The Kukicha compiler maps `import "stdlib/game"` to this module automatically. A registry stub in the main Kukicha repo (`stdlib/game/game.kuki`) provides type information for compile-time checks.

### WASM-only build constraint

The generated `game.go` has a `//go:build js` constraint. Ebitengine's native backends require platform-specific headers (X11 on Linux) which aren't needed for WASM games. The Kukicha compiler automatically adds this constraint to both the game package and any user code that imports `stdlib/game`, so `go build ./...` and `go test ./...` skip game code on native platforms. Always build with `kukicha build --wasm`.

## License

MIT
