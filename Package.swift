// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "Snap",
    platforms: [.iOS(.v15), .macOS(.v12)],
    products: [
        .library(name: "Snap", targets: ["Snap"]),
    ],
    targets: [
        .target(
            name: "CSnap",
            path: ".",
            exclude: [
                "emacs",
                "swift",
                "Makefile",
                "README.org",
                "grammars/workout/grammar.js",
                "grammars/workout/src/grammar.json",
                "grammars/workout/src/node-types.json",
                "grammars/workout/src/tree_sitter",
            ],
            sources: [
                "vendor/tree-sitter/lib/src/lib.c",
                "grammars/workout/src/parser.c",
                "lib/workout.c",
            ],
            publicHeadersPath: "swift/Sources/CSnap/include",
            cSettings: [
                .headerSearchPath("vendor/tree-sitter/lib/src"),
                .headerSearchPath("vendor/tree-sitter/lib/include"),
                .headerSearchPath("grammars/workout/src"),
                .headerSearchPath("lib"),
            ]
        ),
        .target(
            name: "Snap",
            dependencies: ["CSnap"],
            path: "swift/Sources/Snap"
        ),
        .testTarget(
            name: "SnapTests",
            dependencies: ["Snap"],
            path: "swift/Tests/SnapTests"
        ),
    ]
)
