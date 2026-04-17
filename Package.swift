// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "Iruka",
    platforms: [.macOS(.v13)],
    targets: [
        .target(
            name: "IrukaCore",
            path: "Sources/IrukaCore"
        ),
        .executableTarget(
            name: "Iruka",
            dependencies: ["IrukaCore"],
            path: "Sources/Iruka"
        ),
        .testTarget(
            name: "IrukaCoreTests",
            dependencies: ["IrukaCore"],
            path: "Tests/IrukaCoreTests"
        ),
    ]
)
