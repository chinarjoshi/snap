import Foundation

public enum SnapResult: Equatable {
    case workout(WorkoutResult)
    case recipe(RecipeResult)
}

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

public struct Ingredient: Equatable {
    public let quantity: String
    public let text: String

    public init(quantity: String, text: String) {
        self.quantity = quantity
        self.text = text
    }
}

public struct RecipeSection: Equatable {
    public let name: String
    public var ingredients: [Ingredient]
    public var instructions: [String]

    public init(name: String, ingredients: [Ingredient] = [], instructions: [String] = []) {
        self.name = name
        self.ingredients = ingredients
        self.instructions = instructions
    }
}

public struct RecipeResult: Equatable {
    public let title: String
    public var sections: [RecipeSection]

    public init(title: String, sections: [RecipeSection] = []) {
        self.title = title
        self.sections = sections
    }
}
