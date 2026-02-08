import XCTest
@testable import WorkoutLiner

final class WorkoutLinerTests: XCTestCase {

    func testTransformPureWorkout() {
        let input = "squat 8 8 8 135\nbench 3x8 95"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Squat"))
        XCTAssertTrue(result.contains("| Bench"))
        XCTAssertTrue(result.contains("8@135"))
        XCTAssertTrue(result.contains("8@95"))
    }

    func testTransformProseBeforeWorkout() {
        let input = "ate a banana beforehand\nsquat 8 8 8 135"
        let result = transform(input)

        XCTAssertTrue(result.hasPrefix("ate a banana beforehand"))
        XCTAssertTrue(result.contains("| Squat"))
    }

    func testTransformProseAfterWorkout() {
        let input = "squat 8 8 8 135\n\nRegular notes here."
        let result = transform(input)

        XCTAssertTrue(result.contains("| Squat"))
        XCTAssertTrue(result.hasSuffix("Regular notes here."))
    }

    func testTransformMultipleProseLines() {
        let input = "morning workout\nfelt strong\nsquat 8 8 8 135"
        let result = transform(input)

        XCTAssertTrue(result.contains("morning workout"))
        XCTAssertTrue(result.contains("felt strong"))
        XCTAssertTrue(result.contains("| Squat"))
    }

    func testTransformInlineNotes() {
        let input = "squat 8 8 8 135 light headed\nbench 3x8 95"
        let result = transform(input)

        XCTAssertTrue(result.contains("- Squat :: light headed"), "Expected note as description list, got: \(result)")
    }

    func testTransformPureProse() {
        let input = "Just some notes today.\nNothing special."
        let result = transform(input)

        XCTAssertEqual(result, input)
    }

    func testTransformMixedParagraphs() {
        let input = "squat 8 8 8 135\n\nRegular notes here.\n\nbench 3x8 95"
        let result = transform(input)

        let parts = result.components(separatedBy: "\n\n")
        XCTAssertEqual(parts.count, 3)
        XCTAssertTrue(parts[0].contains("| Squat"))
        XCTAssertEqual(parts[1], "Regular notes here.")
        XCTAssertTrue(parts[2].contains("| Bench"))
    }

    func testTransformMultiplierSyntax() {
        let testCases = [
            ("squat 3x8 135", "8@135"),
            ("squat 3x8x135", "8@135"),
            ("squat 8 by 135", "8@135"),
            ("squat 3 by 8 by 135", "8@135"),
        ]

        for (input, expected) in testCases {
            let result = transform(input)
            XCTAssertTrue(result.contains(expected), "Input: \(input), Expected: \(expected), Got: \(result)")
        }
    }

    func testTransformEmptyInput() {
        let result = transform("")
        XCTAssertEqual(result, "")
    }

    func testTransformNoDigitsParagraphSkipped() {
        let input = "Just words here\nno numbers at all"
        let result = transform(input)

        XCTAssertEqual(result, input)
    }

    func testTransformTitleCase() {
        let input = "BARBELL SQUAT 8 8 8 135"
        let result = transform(input)

        XCTAssertTrue(result.contains("Barbell Squat"))
    }

    func testTransformTrailingReps() {
        let input = "squat 8 8 8 135 8 8"
        let result = transform(input)

        let count = result.components(separatedBy: "8@135").count - 1
        XCTAssertEqual(count, 5)
    }

    func testTransformDefaultSetsForBySyntax() {
        let input = "squat 8 by 135"
        let result = transform(input)

        let count = result.components(separatedBy: "8@135").count - 1
        XCTAssertEqual(count, 3)
    }

    func testTransformDefaultRepsForWeight() {
        let input = "squat 135"
        let result = transform(input)

        XCTAssertTrue(result.contains("8@135"))
    }

    func testTransformMultipleExercises() {
        let input = "squat 8 8 8 135\nbench 3x8 95\npress 8 by 35"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Squat"))
        XCTAssertTrue(result.contains("| Bench"))
        XCTAssertTrue(result.contains("| Press"))
    }

    func testTransformPreservesBlankLines() {
        let input = "paragraph one\n\nparagraph two"
        let result = transform(input)

        XCTAssertTrue(result.contains("\n\n"))
    }
}
