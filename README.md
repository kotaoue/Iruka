# Iruka

Iruka makes you smile.

## About The Iruka

When you're coding, you're expressionless or frowning, right?  
Are you surprised when you see yourself on the screen of a web conference?  

Iruka stare at you from desktop and urge you to smile.  
It's a little bothering.

## Getting Started

### Requirements

- macOS 13 (Ventura) or later
- Xcode 15 or later (includes Swift 5.9+ and SwiftPM)

### Run the Desktop Mascot (Swift)

```bash
swift run
```

Or build a standalone binary:

```bash
swift build -c release
.build/release/Iruka
```

🐓 will appear on your desktop.

- **Drag** the window anywhere to move it around the screen.
- **Right-click** the mascot to open a menu and quit.
- **Escape** to quit.

### Run with Go (legacy)

> Requires Go 1.24+ and Xcode Command Line Tools (`xcode-select --install`).

```bash
go run .
```

## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.
