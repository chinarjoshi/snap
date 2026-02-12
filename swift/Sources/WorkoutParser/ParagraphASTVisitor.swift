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

class WorkoutlinerASTVisitor: WorkoutlinerBaseVisitor<Any> {

    override func visitParagraph(_ ctx: WorkoutlinerParser.ParagraphContext) -> Any? {
        var result = WorkoutResult()

        // State for handling continuation lines
        var lastExercise: Exercise? = nil
        var pendingProse: String = ""

        func flush() {
            if let ex = lastExercise, !ex.sets.isEmpty {
                result.exercises.append(ex)
            } else if !pendingProse.isEmpty {
                result.proseLines.append(pendingProse)
            }
            lastExercise = nil
            pendingProse = ""
        }

        for lineCtx in ctx.line() {
            guard let lineResult = visit(lineCtx) else { continue }

            if let prose = lineResult as? String {
                // proseLine
                flush()
                pendingProse = prose
            } else if let exercise = lineResult as? Exercise {
                // exerciseLine
                flush()
                if !exercise.sets.isEmpty {
                    lastExercise = exercise
                } else {
                    pendingProse = exercise.name
                }
            } else if let exercises = lineResult as? [Exercise] {
                // supersetLine
                flush()
                for ex in exercises {
                    if !ex.sets.isEmpty {
                        result.exercises.append(ex)
                    }
                }
            } else if let tokenResults = lineResult as? [TokenResult] {
                // continuationLine
                if var ex = lastExercise {
                    // Extend existing exercise
                    let lastWeight = ex.sets.last?.weight ?? 0
                    let newSets = buildSetsFromResultsWithInitialWeight(tokenResults, initialWeight: lastWeight)
                    ex.sets.append(contentsOf: newSets)
                    lastExercise = ex
                } else if !pendingProse.isEmpty {
                    // Start new exercise from pending prose
                    let sets = buildSetsFromResults(tokenResults)
                    lastExercise = Exercise(name: titleCase(pendingProse), sets: sets)
                    pendingProse = ""
                }
                // If no pending prose or exercise, ignore the continuation line
            }
        }

        flush()
        return result
    }

    override func visitExerciseLineAlt(_ ctx: WorkoutlinerParser.ExerciseLineAltContext) -> Any? {
        guard let exerciseLine = ctx.exerciseLine() else { return nil }
        return visit(exerciseLine)
    }

    override func visitProseLineAlt(_ ctx: WorkoutlinerParser.ProseLineAltContext) -> Any? {
        guard let proseLine = ctx.proseLine() else { return nil }
        return visit(proseLine)
    }

    override func visitContinuationLineAlt(_ ctx: WorkoutlinerParser.ContinuationLineAltContext) -> Any? {
        guard let continuationLine = ctx.continuationLine() else { return nil }
        return visit(continuationLine)
    }

    override func visitContinuationLine(_ ctx: WorkoutlinerParser.ContinuationLineContext) -> Any? {
        var results: [TokenResult] = []

        if let numericToken = ctx.numericToken(),
           let numResult = visit(numericToken) as? TokenResult {
            results.append(numResult)
        }

        for tokenCtx in ctx.token() {
            if let result = visit(tokenCtx) as? TokenResult {
                results.append(result)
            }
        }

        return results
    }

    override func visitSupersetLineAlt(_ ctx: WorkoutlinerParser.SupersetLineAltContext) -> Any? {
        guard let supersetLine = ctx.supersetLine() else { return nil }
        return visit(supersetLine)
    }

    override func visitSupersetLine(_ ctx: WorkoutlinerParser.SupersetLineContext) -> Any? {
        let supersetNames = ctx.supersetName()
        guard supersetNames.count == 2 else { return nil }

        // Extract exercise names (keep * suffix)
        let name1 = extractSupersetName(supersetNames[0])
        let name2 = extractSupersetName(supersetNames[1])

        // Build all sets from tokens
        let tokens = ctx.token()
        let allSets = buildSetsFromTokens(tokens)

        // Alternate sets between two exercises
        var sets1: [WorkoutSet] = []
        var sets2: [WorkoutSet] = []
        for (i, set) in allSets.enumerated() {
            if i % 2 == 0 {
                sets1.append(set)
            } else {
                sets2.append(set)
            }
        }

        return [
            Exercise(name: name1, sets: sets1),
            Exercise(name: name2, sets: sets2)
        ]
    }

    private func extractSupersetName(_ ctx: WorkoutlinerParser.SupersetNameContext) -> String {
        let words = ctx.WORD().map { $0.getText() }
        let name = titleCase(words.joined(separator: " "))
        return name + "*"
    }

    private func buildSetsFromTokens(_ tokens: [WorkoutlinerParser.TokenContext]) -> [WorkoutSet] {
        var sets: [WorkoutSet] = []

        for tokenCtx in tokens {
            guard let result = visit(tokenCtx) as? TokenResult else { continue }

            switch result {
            case .sets(let newSets):
                // Fully specified sets (e.g., 3x8x135) - add each as individual set
                sets.append(contentsOf: newSets)
            case .pendingReps(let pending):
                // Partial multiplier (e.g., 3x8) - add each as bodyweight
                for r in pending {
                    sets.append(WorkoutSet(reps: r, weight: 0))
                }
            case .weight(let w):
                // Single weight - assume default reps
                sets.append(WorkoutSet(reps: defaultReps, weight: w))
            case .reps(let r):
                // Single rep count - bodyweight
                sets.append(WorkoutSet(reps: r, weight: 0))
            case .note:
                // Notes don't create sets in superset context
                break
            }
        }

        return sets
    }

    override func visitExerciseLine(_ ctx: WorkoutlinerParser.ExerciseLineContext) -> Any? {
        guard let numericToken = ctx.numericToken() else { return nil }

        let words = ctx.WORD()
        let tokens = ctx.token()

        // Collect first exercise name
        let firstName = extractName(words: words, numericToken: numericToken)

        // Collect all token results including notes
        var allTokenResults: [TokenResult] = []
        if let numResult = visit(numericToken) as? TokenResult {
            allTokenResults.append(numResult)
        }
        for tokenCtx in tokens {
            if let result = visit(tokenCtx) as? TokenResult {
                allTokenResults.append(result)
            }
        }

        // Split into exercises by detecting note+numericToken pattern
        let exercises = splitIntoExercises(firstName: firstName, tokens: allTokenResults)
        if exercises.count == 1 {
            return exercises[0]
        }
        return exercises
    }

    private func splitIntoExercises(firstName: String, tokens: [TokenResult]) -> [Exercise] {
        var exercises: [Exercise] = []
        var currentName = firstName
        var currentTokens: [TokenResult] = []

        for i in 0..<tokens.count {
            let tok = tokens[i]

            // Check if this note starts a new exercise (note followed by numeric token)
            if case .note(let text) = tok,
               i + 1 < tokens.count,
               isNumericToken(tokens[i + 1]) {
                // Flush the current exercise
                if !currentName.isEmpty {
                    let sets = buildSetsFromResults(currentTokens)
                    exercises.append(Exercise(name: currentName, sets: sets))
                }
                // Start new exercise with this note as name
                currentName = titleCase(text)
                currentTokens = []
            } else {
                currentTokens.append(tok)
            }
        }

        // Flush last exercise
        if !currentName.isEmpty {
            let sets = buildSetsFromResults(currentTokens)
            exercises.append(Exercise(name: currentName, sets: sets))
        }

        return exercises
    }

    private func isNumericToken(_ tok: TokenResult) -> Bool {
        switch tok {
        case .sets, .pendingReps, .weight, .reps:
            return true
        case .note:
            return false
        }
    }

    override func visitProseLine(_ ctx: WorkoutlinerParser.ProseLineContext) -> Any? {
        let words = ctx.WORD().map { $0.getText() }
        return words.joined(separator: " ")
    }

    private func extractName(words: [TerminalNode], numericToken: WorkoutlinerParser.NumericTokenContext) -> String {
        let firstTokenStart = numericToken.getStart()?.getStartIndex() ?? 0
        var nameWords: [String] = []

        for word in words {
            if let symbol = word.getSymbol(), symbol.getStartIndex() < firstTokenStart {
                nameWords.append(word.getText())
            }
        }

        return titleCase(nameWords.joined(separator: " "))
    }

    override func visitNumericToken(_ ctx: WorkoutlinerParser.NumericTokenContext) -> Any? {
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

    override func visitToken(_ ctx: WorkoutlinerParser.TokenContext) -> Any? {
        if let numToken = ctx.numericToken() {
            return visit(numToken)
        }
        if let note = ctx.note() {
            return visit(note)
        }
        return nil
    }

    override func visitTwoPartBy(_ ctx: WorkoutlinerParser.TwoPartByContext) -> Any? {
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0
        let weight = Int(ctx.weight?.getText() ?? "0") ?? 0

        let sets = (0..<defaultSets).map { _ in WorkoutSet(reps: reps, weight: weight) }
        return TokenResult.sets(sets)
    }

    override func visitThreePartBy(_ ctx: WorkoutlinerParser.ThreePartByContext) -> Any? {
        let numSets = Int(ctx.sets?.getText() ?? "0") ?? 0
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0
        let weight = Int(ctx.weight?.getText() ?? "0") ?? 0

        let sets = (0..<numSets).map { _ in WorkoutSet(reps: reps, weight: weight) }
        return TokenResult.sets(sets)
    }

    override func visitFullMultiplier(_ ctx: WorkoutlinerParser.FullMultiplierContext) -> Any? {
        let numSets = Int(ctx.sets?.getText() ?? "0") ?? 0
        let reps = Int(ctx.reps?.getText() ?? "0") ?? 0
        let weight = Int(ctx.weight?.getText() ?? "0") ?? 0

        let sets = (0..<numSets).map { _ in WorkoutSet(reps: reps, weight: weight) }
        return TokenResult.sets(sets)
    }

    override func visitPartialMultiplier(_ ctx: WorkoutlinerParser.PartialMultiplierContext) -> Any? {
        let first = Int(ctx.sets?.getText() ?? "0") ?? 0
        let second = Int(ctx.reps?.getText() ?? "0") ?? 0

        // If second number looks like a weight, treat as reps@weight
        if second > repThreshold {
            return TokenResult.sets([WorkoutSet(reps: first, weight: second)])
        }

        // Otherwise, it's sets x reps waiting for weight
        let pending = (0..<first).map { _ in second }
        return TokenResult.pendingReps(pending)
    }

    override func visitNumber(_ ctx: WorkoutlinerParser.NumberContext) -> Any? {
        let n = Int(ctx.NUMBER()?.getText() ?? "0") ?? 0
        if n > repThreshold {
            return TokenResult.weight(n)
        }
        return TokenResult.reps(n)
    }

    override func visitNote(_ ctx: WorkoutlinerParser.NoteContext) -> Any? {
        let words = ctx.WORD().map { $0.getText() }
        return TokenResult.note(words.joined(separator: " "))
    }

    private func buildSetsFromResultsWithInitialWeight(_ results: [TokenResult], initialWeight: Int) -> [WorkoutSet] {
        return buildSetsFromResultsImpl(results, initialWeight: initialWeight)
    }

    private func buildSetsFromResults(_ results: [TokenResult]) -> [WorkoutSet] {
        return buildSetsFromResultsImpl(results, initialWeight: 0)
    }

    private func buildSetsFromResultsImpl(_ results: [TokenResult], initialWeight: Int) -> [WorkoutSet] {
        var sets: [WorkoutSet] = []
        var pendingReps: [Int] = []
        var lastWeight = initialWeight

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
                // Resolve pending reps before attaching note
                if !pendingReps.isEmpty {
                    for r in pendingReps {
                        sets.append(WorkoutSet(reps: r, weight: lastWeight))
                    }
                    pendingReps.removeAll()
                }
                if !sets.isEmpty {
                    var lastSet = sets.removeLast()
                    lastSet.note = joinNotes(lastSet.note, text)
                    sets.append(lastSet)
                }
            }
        }

        for r in pendingReps {
            sets.append(WorkoutSet(reps: r, weight: lastWeight))
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
