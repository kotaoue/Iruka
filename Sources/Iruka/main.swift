import AppKit

private let mascotChar = "🐓"
private let windowWidth: CGFloat = 100
private let windowHeight: CGFloat = 100
private let fontSize: CGFloat = 48

class MascotLabel: NSTextField {
    override func rightMouseDown(with event: NSEvent) {
        let menu = NSMenu()
        menu.addItem(
            NSMenuItem(
                title: "Quit", action: #selector(NSApplication.terminate(_:)), keyEquivalent: ""))
        NSMenu.popUpContextMenu(menu, with: event, for: self)
    }
}

class AppDelegate: NSObject, NSApplicationDelegate {
    var window: NSWindow!

    func applicationDidFinishLaunching(_ notification: Notification) {
        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: windowWidth, height: windowHeight),
            styleMask: [.borderless],
            backing: .buffered,
            defer: false
        )
        window.isOpaque = false
        window.backgroundColor = .clear
        window.level = .floating
        window.isMovableByWindowBackground = true
        if let visibleFrame = NSScreen.main?.visibleFrame {
            window.setFrameTopLeftPoint(NSPoint(x: visibleFrame.minX, y: visibleFrame.maxY))
        }

        let label = MascotLabel(labelWithString: mascotChar)
        label.font = NSFont.systemFont(ofSize: fontSize)
        label.alignment = .center
        label.frame = NSRect(x: 0, y: 0, width: windowWidth, height: windowHeight)
        label.backgroundColor = .clear

        window.contentView?.addSubview(label)
        window.makeKeyAndOrderFront(nil)

        NSEvent.addLocalMonitorForEvents(matching: .keyDown) { event in
            if event.keyCode == 53 {  // Escape
                NSApplication.shared.terminate(nil)
            }
            return event
        }
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        return true
    }
}

let app = NSApplication.shared
app.setActivationPolicy(.accessory)
let delegate = AppDelegate()
app.delegate = delegate
app.run()
