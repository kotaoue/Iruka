import AppKit

final class AppDelegate: NSObject, NSApplicationDelegate {
    private enum WindowStyle {
        static let mascotChar = "🐓"
        static let width: CGFloat = 100
        static let height: CGFloat = 100
        static let fontSize: CGFloat = 48
    }

    private var window: NSWindow!

    func applicationDidFinishLaunching(_ notification: Notification) {
        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: WindowStyle.width, height: WindowStyle.height),
            styleMask: [.borderless],
            backing: .buffered,
            defer: false
        )
        window.isOpaque = false
        window.backgroundColor = .clear
        window.level = .floating
        window.isMovableByWindowBackground = true

        if let screenFrame = NSScreen.main?.frame {
            let windowFrame = window.frame
            window.setFrameOrigin(
                NSPoint(
                    x: screenFrame.maxX - windowFrame.width,
                    y: screenFrame.minY
                )
            )
        }

        let label = MascotLabel(labelWithString: WindowStyle.mascotChar)
        label.font = NSFont.systemFont(ofSize: WindowStyle.fontSize)
        label.alignment = .center
        label.frame = NSRect(x: 0, y: 0, width: WindowStyle.width, height: WindowStyle.height)
        label.backgroundColor = .clear

        window.contentView?.addSubview(label)
        window.makeKeyAndOrderFront(nil)

        NSEvent.addLocalMonitorForEvents(matching: .keyDown) { event in
            if event.keyCode == 53 {
                NSApplication.shared.terminate(nil)
            }
            return event
        }
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        true
    }
}
