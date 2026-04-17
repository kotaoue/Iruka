import AppKit

public final class MascotLabel: NSTextField {
    public override func rightMouseDown(with event: NSEvent) {
        let menu = NSMenu()
        menu.addItem(
            NSMenuItem(
                title: "Quit",
                action: #selector(NSApplication.terminate(_:)),
                keyEquivalent: ""
            )
        )
        NSMenu.popUpContextMenu(menu, with: event, for: self)
    }
}
