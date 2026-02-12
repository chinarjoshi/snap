// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "WorkoutParser",
    platforms: [
        .iOS(.v15),
        .macOS(.v12)
    ],
    products: [
        .library(
            name: "WorkoutParser",
            targets: ["WorkoutParser"]),
    ],
    dependencies: [
        .package(url: "https://github.com/antlr/antlr4", from: "4.13.0"),
    ],
    targets: [
        .target(
            name: "WorkoutParser",
            dependencies: [
                .product(name: "Antlr4", package: "antlr4"),
            ],
            exclude: [
                "Parser/Paragraph.interp",
                "Parser/Paragraph.tokens",
                "Parser/ParagraphLexer.interp",
                "Parser/ParagraphLexer.tokens",
            ]),
        .testTarget(
            name: "WorkoutParserTests",
            dependencies: ["WorkoutParser"]),
    ]
)
