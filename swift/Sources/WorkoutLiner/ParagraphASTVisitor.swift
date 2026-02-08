import Foundation
import Antlr4

let defaultReps = 8
let defaultSets = 3
let repThreshold = 20

enum TokenResult {
    case sets([WorkoutSet])
    case pendingReps([Int])
    case weight(Int)
    case reps(Int)
    case note(String)
}

class ParagraphASTVisitor: ParagraphBaseVisitor<Any> {

    override func visitParagraph(_ ctx: ParagraphParser.ParagraphContext) -> Any? {
        var result = ParseResult()

        for lineCtx in ctx.line() {
            guard let lineResult = visit(lineCtx) else { continue }

            if let prose = lineResult as? String {
                result.proseLines.append(prose)
            } else if let exercise = lineResult as? Exercise {
                if exercise.sets.isEmpty {
                    result.proseLines.append(exercise.name)
                } else {
                    result.exercises.append(exercise)
                }
            }
        }

        return result
    }

    override func visitExerciseLineAlt(_ ctx: ParagraphParser.ExerciseLineAltContext) -> Any? {
        guard let exerciseLine = ctx.exerciseLine() else { return nil }
        return visit(exerciseLine)
    }

    override func visitProseLineAlt(_ ctx: ParagraphParser.ProseLineAltContext) -> Any? {
        guard let proseLine = ctx.proseLine() else { return nil }
        return visit(proseLine)
    }

    override func visitExerciseLine(_ ctx: ParagraphParser.ExerciseLineContext) -> Any? {
        guard let numericToken = ctx.numericToken() else { return nil }

        let words = ctx.WORD()
        let tokens = ctx.token()

        let name = extractName(words: words, numericToken: numericToken)
        let sets = buildSetsFromExercise(numericToken: numericToken, tokens: tokens)

        return Exercise(name: name, sets: sets)
    }

    override func visitProseLine(_ ctx: ParagraphParser.ProseLineContext) -> Any? {
        let words = ctx.WORD().map { $0.getText() }
        return words.joined(separator: " ")
    }

    private func extractName(words: [TerminalNode], numericToken: ParagraphParser.NumericTokenContext) -> String {
        let firstTokenStart = numericToken.getStart()?.getStartIndex() ?? 0
        var nameWords: [String] = []

        for word in words {
            if let symbol = word.getSymbol(), symbol.getStartIndex() < firstTokenStart {
                nameWords.append(word.getText())
            }
        }

        return titleCase(nameWords.joined(separator: " "))
    }

    private func buildSetsFromExercise(numericToken: ParagraphParser.NumericTokenContext, tokens: [ParagraphParser.TokenContext]) -> [WorkoutSet] {
        var allResults: [TokenResult] = []

        if let numResult = visit(numericToken) as? TokenResult {
            allResults.append(numResult)
        }

        for tokenCtx in tokens {
            if let result = visit(tokenCtx) as? TokenResult {
                allResults.append(result)
            }
        }

        return buildSetsFromResults(allResults)
    }

    override func visitNumericToken(_ ctx: ParagraphParser.NumericTokenContext) -> Any? {
        if let byExpr = ctx.byExpr() {
            return visit(byExpr)
        }
        if let mult = ctx.multiplier() {
            return visit(mult)
        }
        if let num = ctx.number() {
            return visit(num)
        }
        return nil
    }

    override func visitToken(_ ctx: ParagraphParser.TokenContext) -> Any? {
        if let numToken = ctx.numericToken() {
            return visit(numToken)
        }
        if let note = ctx.note() {
            return visit(note)
        }
        return nil
    }

    override func visitTwoPartBy(_ ctx: ParagraphParser.TwoPartByContext) -> Any? {
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0
        let weight = Int(ctx.weight?.getText() ?? "0") ?? 0

        let sets = (0..<defaultSets).map { _ in WorkoutSet(reps: reps, weight: weight) }
        return TokenResult.sets(sets)
    }

    override func visitThreePartBy(_ ctx: ParagraphParser.ThreePartByContext) -> Any? {
        let numSets = Int(ctx.sets?.getText() ?? "0") ?? 0
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0
        let weight = Int(ctx.weight?.getText() ?? "0") ?? 0

        let sets = (0..<numSets).map { _ in WorkoutSet(reps: reps, weight: weight) }
        return TokenResult.sets(sets)
    }

    override func visitFullMultiplier(_ ctx: ParagraphParser.FullMultiplierContext) -> Any? {
        let numSets = Int(ctx.sets?.getText() ?? "0") ?? 0
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0
        let weight = Int(ctx.weight?.getText() ?? "0") ?? 0

        let sets = (0..<numSets).map { _ in WorkoutSet(reps: reps, weight: weight) }
        return TokenResult.sets(sets)
    }

    override func visitPartialMultiplier(_ ctx: ParagraphParser.PartialMultiplierContext) -> Any? {
        let numSets = Int(ctx.sets?.getText() ?? "0") ?? 0
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0

        let pending = (0..<numSets).map { _ in reps }
        return TokenResult.pendingReps(pending)
    }

    override func visitNumber(_ ctx: ParagraphParser.NumberContext) -> Any? {
        let n = Int(ctx.NUMBER()?.getText() ?? "0") ?? 0
        if n > repThreshold {
            return TokenResult.weight(n)
        }
        return TokenResult.reps(n)
    }

    override func visitNote(_ ctx: ParagraphParser.NoteContext) -> Any? {
        let words = ctx.WORD().map { $0.getText() }
        return TokenResult.note(words.joined(separator: " "))
    }

    private func buildSetsFromResults(_ results: [TokenResult]) -> [WorkoutSet] {
        var sets: [WorkoutSet] = []
        var pendingReps: [Int] = []
        var lastWeight = 0

        for result in results {
            switch result {
            case .sets(let newSets):
                sets.append(contentsOf: newSets)
                if let last = newSets.last {
                    lastWeight = last.weight
                }
            case .pendingReps(let pending):
                pendingReps.append(contentsOf: pending)
            case .weight(let w):
                if !pendingReps.isEmpty {
                    for r in pendingReps {
                        sets.append(WorkoutSet(reps: r, weight: w))
                    }
                    pendingReps.removeAll()
                } else {
                    sets.append(WorkoutSet(reps: defaultReps, weight: w))
                }
                lastWeight = w
            case .reps(let r):
                pendingReps.append(r)
            case .note(let text):
                if !sets.isEmpty {
                    var lastSet = sets.removeLast()
                    lastSet.note = joinNotes(lastSet.note, text)
                    sets.append(lastSet)
                }
            }
        }

        for r in pendingReps {
            if lastWeight > 0 {
                sets.append(WorkoutSet(reps: r, weight: lastWeight))
            }
        }

        return sets
    }

    private func titleCase(_ s: String) -> String {
        let words = s.trimmingCharacters(in: .whitespaces).components(separatedBy: .whitespaces)
        return words.map { word in
            guard !word.isEmpty else { return word }
            return word.prefix(1).uppercased() + word.dropFirst().lowercased()
        }.joined(separator: " ")
    }

    private func joinNotes(_ existing: String, _ new: String) -> String {
        if existing.isEmpty { return new }
        if new.isEmpty { return existing }
        return "\(existing) \(new)"
    }
}
