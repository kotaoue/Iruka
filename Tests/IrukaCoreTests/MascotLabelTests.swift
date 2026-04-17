import XCTest
import AppKit
@testable import IrukaCore

final class MascotLabelTests: XCTestCase {
    func testMascotLabelInheritsFromNSTextField() {
        let label = MascotLabel(labelWithString: "🐓")
        XCTAssertTrue(label is NSTextField)
    }

    func testMascotLabelCanSetStringValue() {
        let label = MascotLabel(labelWithString: "🐓")
        label.stringValue = "🐬"
        XCTAssertEqual(label.stringValue, "🐬")
    }
}
