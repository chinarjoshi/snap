import XCTest
@testable import Snap

final class WorkoutTests: XCTestCase {

    // MARK: - Helpers

    private func exercise(_ result: WorkoutResult, named name: String, file: StaticString = #file, line: UInt = #line) -> Exercise? {
        guard let ex = result.exercises.first(where: { $0.name == name }) else {
            XCTFail("Expected exercise '\(name)', found: \(result.exercises.map { $0.name })", file: file, line: line)
            return nil
        }
        return ex
    }

    // MARK: - Basic Parsing

    func testParsePureWorkout() {
        guard let w = parse("squat 8 8 8 135\nbench 3x8 95") else {
            return XCTFail("Expected workout")
        }

        guard let squat = exercise(w, named: "Squat") else { return }
        XCTAssertEqual(squat.sets.count, 3)
        XCTAssertEqual(squat.sets[0], WorkoutSet(reps: 8, weight: 135))

        guard let bench = exercise(w, named: "Bench") else { return }
        XCTAssertEqual(bench.sets.count, 3)
        XCTAssertEqual(bench.sets[0], WorkoutSet(reps: 8, weight: 95))
    }

    func testParseProseBeforeWorkout() {
        guard let w = parse("ate a banana beforehand\nsquat 8 8 8 135") else {
            return XCTFail("Expected workout")
        }

        XCTAssertEqual(w.proseLines, ["ate a banana beforehand"])
        XCTAssertNotNil(exercise(w, named: "Squat"))
    }

    func testParseMultipleProseLines() {
        guard let w = parse("morning workout\nfelt strong\nsquat 8 8 8 135") else {
            return XCTFail("Expected workout")
        }

        XCTAssertTrue(w.proseLines.contains("morning workout"))
        XCTAssertTrue(w.proseLines.contains("felt strong"))
        XCTAssertNotNil(exercise(w, named: "Squat"))
    }

    func testParseInlineNotes() {
        guard let w = parse("squat 8 8 8 135 light headed\nbench 3x8 95") else {
            return XCTFail("Expected workout")
        }

        guard let squat = exercise(w, named: "Squat") else { return }
        let notes = squat.sets.compactMap { $0.note.isEmpty ? nil : $0.note }
        XCTAssertTrue(notes.contains("light headed"), "Expected 'light headed' note, got: \(notes)")
    }

    func testParseMultiplierSyntax() {
        let cases: [(String, Int, Int)] = [
            ("squat 3x8 135", 8, 135),
            ("squat 3x8x135", 8, 135),
            ("squat 8 by 135", 8, 135),
            ("squat 3 by 8 by 135", 8, 135),
        ]

        for (input, expectedReps, expectedWeight) in cases {
            guard let w = parse(input) else {
                XCTFail("Expected workout for: \(input)")
                continue
            }
            guard let squat = exercise(w, named: "Squat") else { continue }
            XCTAssertEqual(squat.sets[0].reps, expectedReps, "Input: \(input)")
            XCTAssertEqual(squat.sets[0].weight, expectedWeight, "Input: \(input)")
        }
    }

    func testParseTitleCase() {
        guard let w = parse("BARBELL SQUAT 8 8 8 135") else {
            return XCTFail("Expected workout")
        }
        XCTAssertNotNil(exercise(w, named: "Barbell Squat"))
    }

    func testParseTrailingReps() {
        guard let w = parse("squat 8 8 8 135 8 8") else {
            return XCTFail("Expected workout")
        }
        guard let squat = exercise(w, named: "Squat") else { return }

        XCTAssertEqual(squat.sets.count, 5)
        for set in squat.sets {
            XCTAssertEqual(set, WorkoutSet(reps: 8, weight: 135))
        }
    }

    func testParseDefaultSetsForBySyntax() {
        guard let w = parse("squat 8 by 135") else {
            return XCTFail("Expected workout")
        }
        guard let squat = exercise(w, named: "Squat") else { return }
        XCTAssertEqual(squat.sets.count, 3)
    }

    func testParseDefaultRepsForWeight() {
        guard let w = parse("squat 135 135") else {
            return XCTFail("Expected workout")
        }
        guard let squat = exercise(w, named: "Squat") else { return }
        XCTAssertEqual(squat.sets[0], WorkoutSet(reps: 8, weight: 135))
    }

    func testParseMultipleExercises() {
        guard let w = parse("squat 8 8 8 135\nbench 3x8 95\npress 8 by 35") else {
            return XCTFail("Expected workout")
        }

        XCTAssertEqual(w.exercises.count, 3)
        XCTAssertNotNil(exercise(w, named: "Squat"))
        XCTAssertNotNil(exercise(w, named: "Bench"))
        XCTAssertNotNil(exercise(w, named: "Press"))
    }

    // MARK: - Returns nil

    func testParseReturnsNilForEmpty() {
        XCTAssertNil(parse(""))
    }

    func testParseReturnsNilForPureProse() {
        XCTAssertNil(parse("Just some notes today.\nNothing special."))
    }

    func testParseReturnsNilForNoDigits() {
        XCTAssertNil(parse("Just words here\nno numbers at all"))
    }

    // MARK: - Bodyweight

    func testParseBodyweight() {
        guard let w = parse("dips 8 8 8\nbench 8x135 8x130") else {
            return XCTFail("Expected workout")
        }

        guard let dips = exercise(w, named: "Dips") else { return }
        for set in dips.sets {
            XCTAssertEqual(set.weight, 0, "Bodyweight exercise should have weight 0")
        }
    }

    // MARK: - Supersets

    func testParseSupersetBasic() {
        guard let w = parse("dips* press* 8 135 8 135 8 135") else {
            return XCTFail("Expected workout")
        }

        XCTAssertNotNil(exercise(w, named: "Dips*"))
        XCTAssertNotNil(exercise(w, named: "Press*"))
    }

    func testParseSupersetMixedBodyweightWeighted() {
        guard let w = parse("dips* shoulder press* 8 8x135 8 8x125 8 8x100") else {
            return XCTFail("Expected workout")
        }

        guard let dips = exercise(w, named: "Dips*") else { return }
        XCTAssertEqual(dips.sets.count, 3)
        for set in dips.sets {
            XCTAssertEqual(set.weight, 0)
        }

        guard let press = exercise(w, named: "Shoulder Press*") else { return }
        XCTAssertEqual(press.sets.count, 3)
        XCTAssertEqual(press.sets[0].weight, 135)
        XCTAssertEqual(press.sets[1].weight, 125)
        XCTAssertEqual(press.sets[2].weight, 100)
    }

    func testParseSupersetOddSets() {
        guard let w = parse("dips* press* 8 135 8 135 8") else {
            return XCTFail("Expected workout")
        }

        guard let dips = exercise(w, named: "Dips*") else { return }
        XCTAssertEqual(dips.sets.count, 3) // indices 0,2,4

        guard let press = exercise(w, named: "Press*") else { return }
        XCTAssertEqual(press.sets.count, 2) // indices 1,3
    }

    // MARK: - Multi-Line

    func testParseMultiLineBasic() {
        guard let w = parse("Bench\n8x135\n8\n8\n8") else {
            return XCTFail("Expected workout")
        }

        guard let bench = exercise(w, named: "Bench") else { return }
        XCTAssertEqual(bench.sets.count, 4)
        for set in bench.sets {
            XCTAssertEqual(set, WorkoutSet(reps: 8, weight: 135))
        }
    }

    func testParseMultiLineInlinePlusContinuation() {
        guard let w = parse("shoulder press 7x100\n6\n5") else {
            return XCTFail("Expected workout")
        }

        guard let press = exercise(w, named: "Shoulder Press") else { return }
        XCTAssertEqual(press.sets.count, 3)
        XCTAssertEqual(press.sets[0], WorkoutSet(reps: 7, weight: 100))
        XCTAssertEqual(press.sets[1], WorkoutSet(reps: 6, weight: 100))
        XCTAssertEqual(press.sets[2], WorkoutSet(reps: 5, weight: 100))
    }

    func testParseMultiLineBodyweight() {
        guard let w = parse("Dips\n8\n8\n8\nbench 8x135 8x130") else {
            return XCTFail("Expected workout")
        }

        guard let dips = exercise(w, named: "Dips") else { return }
        for set in dips.sets {
            XCTAssertEqual(set.weight, 0)
        }
    }

    func testParseMultiLineMixedExercises() {
        guard let w = parse("Bench\n8x135\nshoulder press 7x100\n6") else {
            return XCTFail("Expected workout")
        }

        guard let bench = exercise(w, named: "Bench") else { return }
        XCTAssertEqual(bench.sets[0], WorkoutSet(reps: 8, weight: 135))

        guard let press = exercise(w, named: "Shoulder Press") else { return }
        XCTAssertEqual(press.sets.count, 2)
        XCTAssertEqual(press.sets[0], WorkoutSet(reps: 7, weight: 100))
        XCTAssertEqual(press.sets[1], WorkoutSet(reps: 6, weight: 100))
    }

    // MARK: - Single-Line Multi-Exercise

    func testParseSingleLineMultipleExercises() {
        guard let w = parse("squat 8x135 bench 8x95 press 8x45") else {
            return XCTFail("Expected workout")
        }

        XCTAssertEqual(w.exercises.count, 3)
        XCTAssertEqual(exercise(w, named: "Squat")?.sets[0].weight, 135)
        XCTAssertEqual(exercise(w, named: "Bench")?.sets[0].weight, 95)
        XCTAssertEqual(exercise(w, named: "Press")?.sets[0].weight, 45)
    }

    func testParseSingleLineMultiWordWithModifier() {
        guard let w = parse("squat 8x135 heavy bench 8x95") else {
            return XCTFail("Expected workout")
        }

        XCTAssertNotNil(exercise(w, named: "Squat"))
        XCTAssertNotNil(exercise(w, named: "Heavy Bench"))
    }

    func testParseSingleLineMultiWordExercise() {
        guard let w = parse("squat 8x135 bench press 8x95") else {
            return XCTFail("Expected workout")
        }

        XCTAssertNotNil(exercise(w, named: "Squat"))
        XCTAssertNotNil(exercise(w, named: "Bench Press"))
    }

    // MARK: - Heuristics

    func testHeuristicRejectsProseWithNumber() {
        let cases = [
            "I ran 5 miles today",
            "ate 3 eggs",
            "rested 10 minutes between sets",
            "dips 8",
            "squat 135",
            "bench 8x135",
            "dips 8 8 8",
            "rested 30 minutes",
        ]
        for input in cases {
            XCTAssertNil(parse(input), "'\(input)' should not parse as workout")
        }
    }

    func testHeuristicAcceptsWorkout() {
        let cases = [
            "squat 135 135 135",
            "bench 8x135 8x130 8x125",
            "squat 135 bench 95",
            "squat 3x8x135",
        ]
        for input in cases {
            XCTAssertNotNil(parse(input), "'\(input)' should parse as workout")
        }
    }

    // MARK: - Notes

    func testParseMultiLineNotesPreserved() {
        guard let w = parse("Squat\n8x225 ouch\n7 dang\n5 fuck") else {
            return XCTFail("Expected workout")
        }

        guard let squat = exercise(w, named: "Squat") else { return }
        XCTAssertEqual(squat.sets.count, 3)
        XCTAssertEqual(squat.sets[0], WorkoutSet(reps: 8, weight: 225, note: "ouch"))
        XCTAssertEqual(squat.sets[1], WorkoutSet(reps: 7, weight: 225, note: "dang"))
        XCTAssertEqual(squat.sets[2], WorkoutSet(reps: 5, weight: 225, note: "fuck"))
    }
}
