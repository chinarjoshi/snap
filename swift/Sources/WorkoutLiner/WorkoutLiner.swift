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
    for line in s.components(separatedBy: "\n") {
        // Skip lines with table formatting
        if line.contains("|") {
            continue
        }
        if line.contains(where: { $0.isNumber }) {
            return true
        }
    }
    return false
}

private func transformParagraph(_ para: String) -> String {
    guard containsDigit(para) else {
        return para
    }

    // Check if paragraph starts with an existing table
    let lines = para.components(separatedBy: "\n")
    if !lines.isEmpty && lines[0].trimmingCharacters(in: .whitespaces).hasPrefix("|") {
        return transformTableWithTrailing(lines)
    }

    guard let parseResult = parse(para), parseResult.hasExercises else {
        return para
    }

    return formatResult(parseResult)
}

private func transformTableWithTrailing(_ lines: [String]) -> String {
    var tableLines: [String] = []
    var trailingLines: [String] = []
    var descListLines: [String] = []
    var inTable = true

    for line in lines {
        let trimmed = line.trimmingCharacters(in: .whitespaces)
        if inTable {
            if trimmed.hasPrefix("|") {
                tableLines.append(line)
            } else if trimmed.hasPrefix("- ") && trimmed.contains(" :: ") {
                descListLines.append(line)
            } else if trimmed.isEmpty {
                continue
            } else {
                inTable = false
                trailingLines.append(line)
            }
        } else {
            trailingLines.append(line)
        }
    }

    // If no trailing lines, return original
    guard !trailingLines.isEmpty else {
        return lines.joined(separator: "\n")
    }

    // Parse existing table into exercises
    var existingExercises = parseTableLines(tableLines)

    // Parse existing description list into notes
    let existingNotes = parseDescList(descListLines)
    for i in 0..<existingExercises.count {
        if let note = existingNotes[existingExercises[i].name] {
            for j in 0..<existingExercises[i].sets.count {
                if existingExercises[i].sets[j].note.isEmpty {
                    existingExercises[i].sets[j].note = note
                    break
                }
            }
        }
    }

    // Parse trailing lines as new exercises
    let trailingPara = trailingLines.joined(separator: "\n")
    guard let parseResult = parse(trailingPara), parseResult.hasExercises else {
        return lines.joined(separator: "\n")
    }

    // Merge exercises
    let allExercises = existingExercises + parseResult.exercises

    // Format combined result
    let result = ParseResult(proseLines: parseResult.proseLines, exercises: allExercises)
    return formatResult(result)
}

private func parseTableLines(_ lines: [String]) -> [Exercise] {
    var exercises: [Exercise] = []

    for line in lines {
        // Skip header and separator
        if line.contains("---") {
            continue
        }

        let cells = line.components(separatedBy: "|")
        guard cells.count >= 3 else { continue }

        // First cell (after empty) is the name
        let name = cells[1].trimmingCharacters(in: .whitespaces)
        guard !name.isEmpty else { continue }

        var sets: [WorkoutSet] = []
        for cell in cells.dropFirst(2) {
            let trimmedCell = cell.trimmingCharacters(in: .whitespaces)
            guard !trimmedCell.isEmpty && trimmedCell.contains("@") else { continue }

            let parts = trimmedCell.components(separatedBy: "@")
            if parts.count == 2,
               let reps = Int(parts[0].trimmingCharacters(in: .whitespaces)),
               let weight = Int(parts[1].trimmingCharacters(in: .whitespaces)),
               reps > 0 && weight > 0 {
                sets.append(WorkoutSet(reps: reps, weight: weight))
            }
        }

        if !sets.isEmpty {
            exercises.append(Exercise(name: name, sets: sets))
        }
    }

    return exercises
}

private func parseDescList(_ lines: [String]) -> [String: String] {
    var notes: [String: String] = [:]
    for line in lines {
        var trimmed = line.trimmingCharacters(in: .whitespaces)
        guard trimmed.hasPrefix("- ") else { continue }
        trimmed = String(trimmed.dropFirst(2))
        if let range = trimmed.range(of: " :: ") {
            let name = String(trimmed[..<range.lowerBound])
            let note = String(trimmed[range.upperBound...])
            notes[name] = note
        }
    }
    return notes
}

func parse(_ input: String) -> ParseResult? {
    do {
        let inputStream = ANTLRInputStream(input)
        let lexer = WorkoutlinerLexer(inputStream)
        let tokenStream = CommonTokenStream(lexer)
        let parser = try WorkoutlinerParser(tokenStream)

        let tree = try parser.paragraph()

        let visitor = WorkoutlinerASTVisitor()
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
                let setStr = set.weight == 0 ? "\(set.reps)" : "\(set.reps)@\(set.weight)"
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
