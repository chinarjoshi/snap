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

    func testTransformTableMerge() {
        let input = "|      | 0     | 1     | 2     |\n|-------|-------|-------|-------|\n| Squat| 8@225 | 8@225 | 8@225 |\nbench 3x8 155"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Squat"), "Expected Squat in merged table")
        XCTAssertTrue(result.contains("| Bench"), "Expected Bench in merged table")
        XCTAssertTrue(result.contains("8@225"), "Expected 8@225 in merged table")
        XCTAssertTrue(result.contains("8@155"), "Expected 8@155 in merged table")
    }

    func testTransformBodyweight() {
        let input = "dips 8 8 8"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Dips"), "Expected Dips in table")
        XCTAssertFalse(result.contains("@"), "Bodyweight should not have @ symbol, got: \(result)")
    }

    func testTransformSupersetBasic() {
        let input = "dips* press* 8 135 8 135 8 135"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Dips*"), "Expected Dips* in table, got: \(result)")
        XCTAssertTrue(result.contains("| Press*"), "Expected Press* in table, got: \(result)")
    }

    func testTransformSupersetMixedBodyweightWeighted() {
        let input = "dips* shoulder press* 8 8x135 8 8x125 8 8x100"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Dips*"), "Expected Dips* in table, got: \(result)")
        XCTAssertTrue(result.contains("| Shoulder Press*"), "Expected Shoulder Press* in table, got: \(result)")
        XCTAssertTrue(result.contains("8@135"), "Expected 8@135 for press, got: \(result)")
        XCTAssertTrue(result.contains("8@125"), "Expected 8@125 for press, got: \(result)")
        XCTAssertTrue(result.contains("8@100"), "Expected 8@100 for press, got: \(result)")
    }

    func testTransformSupersetOddSets() {
        let input = "dips* press* 8 135 8 135 8"
        let result = transform(input)

        // With 5 sets: indices 0,2,4 go to dips (3 sets), indices 1,3 go to press (2 sets)
        XCTAssertTrue(result.contains("| Dips*"), "Expected Dips* in table, got: \(result)")
        XCTAssertTrue(result.contains("| Press*"), "Expected Press* in table, got: \(result)")
    }

    func testTransformMultiLineBasic() {
        let input = "Bench\n8x135\n8\n8\n8"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Bench"), "Expected Bench in table, got: \(result)")
        // Should have 4 sets of 8@135 (first set explicit, next 3 inherit weight)
        let count = result.components(separatedBy: "8@135").count - 1
        XCTAssertEqual(count, 4, "Expected 4 sets of 8@135, got \(count) in: \(result)")
    }

    func testTransformMultiLineInlinePlusContinuation() {
        let input = "shoulder press 7x100\n6\n5"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Shoulder Press"), "Expected Shoulder Press in table, got: \(result)")
        XCTAssertTrue(result.contains("7@100"), "Expected 7@100, got: \(result)")
        XCTAssertTrue(result.contains("6@100"), "Expected 6@100, got: \(result)")
        XCTAssertTrue(result.contains("5@100"), "Expected 5@100, got: \(result)")
    }

    func testTransformMultiLineBodyweight() {
        let input = "Dips\n8\n8\n8"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Dips"), "Expected Dips in table, got: \(result)")
        // Bodyweight - no @ symbol
        XCTAssertFalse(result.contains("@"), "Expected bodyweight (no @), got: \(result)")
    }

    func testTransformMultiLineMixedExercises() {
        let input = "Bench\n8x135\nshoulder press 7x100\n6"
        let result = transform(input)

        XCTAssertTrue(result.contains("| Bench"), "Expected Bench in table, got: \(result)")
        XCTAssertTrue(result.contains("| Shoulder Press"), "Expected Shoulder Press in table, got: \(result)")
        XCTAssertTrue(result.contains("8@135"), "Expected 8@135, got: \(result)")
        XCTAssertTrue(result.contains("7@100"), "Expected 7@100, got: \(result)")
        XCTAssertTrue(result.contains("6@100"), "Expected 6@100, got: \(result)")
    }
}
