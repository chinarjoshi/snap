import Foundation
import Antlr4

public func transform(_ input: String) -> String {
    let paragraphs = input.components(separatedBy: "\n\n")
    var result: [String] = []

    for para in paragraphs {
        let transformed = transformParagraph(para)
        result.append(transformed)
    }

    return result.joined(separator: "\n\n")
}

private func containsDigit(_ s: String) -> Bool {
    s.contains(where: { $0.isNumber })
}

private func transformParagraph(_ para: String) -> String {
    guard containsDigit(para) else {
        return para
    }

    guard let parseResult = parse(para), parseResult.hasExercises else {
        return para
    }

    return formatResult(parseResult)
}

func parse(_ input: String) -> ParseResult? {
    do {
        let inputStream = ANTLRInputStream(input)
        let lexer = ParagraphLexer(inputStream)
        let tokenStream = CommonTokenStream(lexer)
        let parser = try ParagraphParser(tokenStream)

        let tree = try parser.paragraph()

        let visitor = ParagraphASTVisitor()
        guard let result = visitor.visit(tree) as? ParseResult else { return nil }

        return result
    } catch {
        return nil
    }
}

private func formatResult(_ result: ParseResult) -> String {
    var parts: [String] = []

    if !result.proseLines.isEmpty {
        parts.append(result.proseLines.joined(separator: "\n"))
    }

    let tableStr = formatExercises(result.exercises)
    parts.append(tableStr)

    return parts.joined(separator: "\n\n")
}

private func formatExercises(_ exercises: [Exercise]) -> String {
    guard !exercises.isEmpty else { return "" }

    let maxSets = exercises.map { $0.sets.count }.max() ?? 1
    let maxNameLen = exercises.map { $0.name.count }.max() ?? 0

    var lines: [String] = []

    // Header
    var header = "| " + "".padding(toLength: maxNameLen, withPad: " ", startingAt: 0) + "|"
    for i in 0..<maxSets {
        header += " " + String(i).padding(toLength: 5, withPad: " ", startingAt: 0) + " |"
    }
    lines.append(header)

    // Separator
    var separator = "|" + String(repeating: "-", count: maxNameLen + 2) + "|"
    for _ in 0..<maxSets {
        separator += "-------|"
    }
    lines.append(separator)

    // Data rows
    for exercise in exercises {
        var row = "| " + exercise.name.padding(toLength: maxNameLen, withPad: " ", startingAt: 0) + "|"
        for i in 0..<maxSets {
            if i < exercise.sets.count {
                let set = exercise.sets[i]
                let setStr = "\(set.reps)@\(set.weight)"
                row += " " + setStr.padding(toLength: 5, withPad: " ", startingAt: 0) + " |"
            } else {
                row += "       |"
            }
        }
        lines.append(row)
    }

    // Notes as description list
    var notes: [String] = []
    for exercise in exercises {
        let noteText = exercise.sets.compactMap { $0.note.isEmpty ? nil : $0.note }.joined(separator: ", ")
        if !noteText.isEmpty {
            notes.append("- \(exercise.name) :: \(noteText)")
        }
    }
    if !notes.isEmpty {
        lines.append("")
        lines.append(contentsOf: notes)
    }

    return lines.joined(separator: "\n")
}
