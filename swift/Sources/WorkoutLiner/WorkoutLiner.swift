import Foundation
import Antlr4

public func parse(_ input: String) -> WorkoutResult? {
    guard !input.isEmpty, containsDigit(input) else { return nil }

    do {
        let inputStream = ANTLRInputStream(input)
        let lexer = WorkoutlinerLexer(inputStream)
        let tokenStream = CommonTokenStream(lexer)
        let parser = try WorkoutlinerParser(tokenStream)

        let tree = try parser.paragraph()

        let visitor = WorkoutlinerASTVisitor()
        guard let result = visitor.visit(tree) as? WorkoutResult else { return nil }
        guard isWorkout(result) else { return nil }

        return result
    } catch {
        return nil
    }
}

private func containsDigit(_ s: String) -> Bool {
    for line in s.components(separatedBy: "\n") {
        if line.contains("|") {
            continue
        }
        if line.contains(where: { $0.isNumber }) {
            return true
        }
    }
    return false
}

private func isWorkout(_ result: WorkoutResult) -> Bool {
    guard result.hasExercises else { return false }
    var totalSets = 0
    var hasWeight = false
    for exercise in result.exercises {
        totalSets += exercise.sets.count
        for set in exercise.sets {
            if set.weight > 0 {
                hasWeight = true
            }
        }
    }
    return totalSets >= 2 && hasWeight
}
