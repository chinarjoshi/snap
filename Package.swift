// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "SnapParsers",
    platforms: [.iOS(.v15), .macOS(.v12)],
    products: [
        .library(name: "SnapParsers", targets: ["SnapParsers"]),
    ],
    targets: [
        .target(
            name: "CSnap",
            path: ".",
            exclude: [
                "docs",
                "emacs",
                "swift",
                "Makefile",
                "README.org",
                "grammars/workout/grammar.js",
                "grammars/workout/package.json",
                "grammars/workout/src/grammar.json",
                "grammars/workout/src/node-types.json",
                "grammars/workout/src/tree_sitter",
                "vendor/tree-sitter/src/wasm",
                "vendor/tree-sitter/src/portable",
                "vendor/tree-sitter/src/unicode",
            ],
            sources: [
                "vendor/tree-sitter/src/lib.c",
                "grammars/workout/src/parser.c",
                "lib/workout.c",
            ],
            publicHeadersPath: "swift/Sources/CSnap/include",
            cSettings: [
                .headerSearchPath("vendor/tree-sitter/src"),
                .headerSearchPath("vendor/tree-sitter/include"),
                .headerSearchPath("grammars/workout/src"),
                .headerSearchPath("lib"),
            ]
        ),
        .target(
            name: "SnapParsers",
            dependencies: ["CSnap"],
            path: "swift/Sources/SnapParsers"
        ),
        .testTarget(
            name: "SnapTests",
            dependencies: ["SnapParsers"],
            path: "swift/Tests/SnapTests"
        ),
    ]
)
