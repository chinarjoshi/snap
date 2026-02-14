import Foundation

public struct WorkoutSet: Equatable {
    public let reps: Int
    public let weight: Int
    public var note: String

    public init(reps: Int, weight: Int, note: String = "") {
        self.reps = reps
        self.weight = weight
        self.note = note
    }
}

public struct Exercise: Equatable {
    public let name: String
    public var sets: [WorkoutSet]

    public init(name: String, sets: [WorkoutSet]) {
        self.name = name
        self.sets = sets
    }
}

public struct WorkoutResult: Equatable {
    public var proseLines: [String]
    public var exercises: [Exercise]

    public init(proseLines: [String] = [], exercises: [Exercise] = []) {
        self.proseLines = proseLines
        self.exercises = exercises
    }

    public var hasExercises: Bool { !exercises.isEmpty }
}
