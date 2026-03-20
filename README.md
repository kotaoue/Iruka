# Iruka

Iruka makes you smile.

## About The Iruka

When you're coding, you're expressionless or frowning, right?  
Are you surprised when you see yourself on the screen of a web conference?  

Iruka stare at you from desktop and urge you to smile.  
It's a little bothering.

## Getting Started

### Requirements

- macOS
- Go 1.24+
- Xcode Command Line Tools (`xcode-select --install`)

### Run the Desktop Mascot

```bash
go run .
```

Or build a standalone binary:

```bash
go build -o Iruka .
./Iruka
```

🐓 will appear on your desktop.

- **Drag** the window by its title bar to move it around the screen.
- **Escape** to quit.
- **Close button** hides the mascot to the system tray (right-click tray icon → Quit to exit).

## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.
