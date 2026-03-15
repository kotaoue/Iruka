#!/usr/bin/env python3
"""
Iruka Desktop Mascot
Displays 🐓 on the desktop as a resident mascot that stays on top of other windows.
The 🐓 character is used as the mascot icon for the Iruka project.
"""

import tkinter as tk
import sys
import platform


MASCOT_CHAR = "🐓"
FONT_SIZE = 48
BG_COLOR = "#F0F0F0"
FG_COLOR = "#333333"
INITIAL_X = 100
INITIAL_Y = 100


def is_mac():
    return platform.system() == "Darwin"


class DesktopMascot:
    def __init__(self):
        self.root = tk.Tk()
        self._setup_window()
        self._setup_label()
        self._setup_drag()
        self._setup_exit()

    def _setup_window(self):
        self.root.title("Iruka")
        self.root.overrideredirect(True)
        self.root.attributes("-topmost", True)
        self.root.configure(bg=BG_COLOR)
        self.root.geometry(f"+{INITIAL_X}+{INITIAL_Y}")

        if is_mac():
            self.root.attributes("-alpha", 0.85)

    def _setup_label(self):
        self.label = tk.Label(
            self.root,
            text=MASCOT_CHAR,
            font=("", FONT_SIZE),
            bg=BG_COLOR,
            fg=FG_COLOR,
            cursor="fleur",
        )
        self.label.pack(padx=8, pady=8)

    def _setup_drag(self):
        self._drag_x = 0
        self._drag_y = 0

        self.label.bind("<ButtonPress-1>", self._on_drag_start)
        self.label.bind("<B1-Motion>", self._on_drag_motion)

    def _on_drag_start(self, event):
        self._drag_x = event.x
        self._drag_y = event.y

    def _on_drag_motion(self, event):
        x = self.root.winfo_x() + event.x - self._drag_x
        y = self.root.winfo_y() + event.y - self._drag_y
        self.root.geometry(f"+{x}+{y}")

    def _setup_exit(self):
        self.label.bind("<Button-3>", lambda e: self.root.destroy())
        self.root.bind("<Escape>", lambda e: self.root.destroy())

    def run(self):
        self.root.mainloop()


def main():
    if not is_mac():
        print("Warning: This mascot is designed for macOS.", file=sys.stderr)

    mascot = DesktopMascot()
    mascot.run()


if __name__ == "__main__":
    main()
