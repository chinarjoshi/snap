import XCTest
@testable import Snap

final class RecipeTests: XCTestCase {

    // MARK: - Helpers

    private func recipe(_ input: String, file: StaticString = #file, line: UInt = #line) -> RecipeResult? {
        guard let result = snap(input) else { return nil }
        guard case .recipe(let r) = result else {
            XCTFail("Expected .recipe, got \(result)", file: file, line: line)
            return nil
        }
        return r
    }

    private func section(_ result: RecipeResult, named name: String, file: StaticString = #file, line: UInt = #line) -> RecipeSection? {
        guard let sec = result.sections.first(where: { $0.name == name }) else {
            XCTFail("Expected section '\(name)', found: \(result.sections.map { $0.name })", file: file, line: line)
            return nil
        }
        return sec
    }

    // MARK: - Basic

    func testBasicRecipeWithSections() {
        guard let r = recipe("chicken biryani:\nchicken marinade:\n1 tbsp jeera\n1 bay leaf\nrice water:\n2 cups rice") else {
            return XCTFail("Expected recipe")
        }
        XCTAssertEqual(r.title, "Chicken Biryani")
        XCTAssertEqual(r.sections.count, 2)

        guard let cm = section(r, named: "Chicken Marinade") else { return }
        XCTAssertEqual(cm.ingredients.count, 2)
        XCTAssertEqual(cm.ingredients[0], Ingredient(quantity: "1", text: "tbsp jeera"))
        XCTAssertEqual(cm.ingredients[1], Ingredient(quantity: "1", text: "bay leaf"))

        guard let rw = section(r, named: "Rice Water") else { return }
        XCTAssertEqual(rw.ingredients.count, 1)
        XCTAssertEqual(rw.ingredients[0], Ingredient(quantity: "2", text: "cups rice"))
    }

    func testRecipeWithInstructions() {
        guard let r = recipe("pasta:\n2 cups flour\n1 egg\nMix well\nKnead for 10 min") else {
            return XCTFail("Expected recipe")
        }
        XCTAssertEqual(r.title, "Pasta")
        XCTAssertEqual(r.sections.count, 1)
        XCTAssertEqual(r.sections[0].ingredients.count, 2)
        XCTAssertEqual(r.sections[0].instructions.count, 2)
    }

    func testRangeQuantity() {
        guard let r = recipe("spice mix:\n3-5 pieces cardamon\n8 cloves") else {
            return XCTFail("Expected recipe")
        }
        XCTAssertEqual(r.sections[0].ingredients[0], Ingredient(quantity: "3-5", text: "pieces cardamon"))
        XCTAssertEqual(r.sections[0].ingredients[1], Ingredient(quantity: "8", text: "cloves"))
    }

    // MARK: - Returns nil

    func testNotRecipeProse() {
        XCTAssertNil(snap("Just some notes today"))
    }

    func testNotRecipeNoIngredients() {
        XCTAssertNil(snap("dinner plan:\nGo to the store"))
    }

    func testNotRecipeEmpty() {
        XCTAssertNil(snap(""))
    }

    // MARK: - Formatting

    func testTitleCase() {
        guard let r = recipe("CHICKEN TIKKA:\n1 cup yogurt\n1 tsp turmeric") else {
            return XCTFail("Expected recipe")
        }
        XCTAssertEqual(r.title, "Chicken Tikka")
    }

    func testRecipeNoSubsections() {
        guard let r = recipe("smoothie:\n1 banana\n2 cups milk\n1 tbsp honey") else {
            return XCTFail("Expected recipe")
        }
        XCTAssertEqual(r.title, "Smoothie")
        XCTAssertEqual(r.sections.count, 1)
        XCTAssertEqual(r.sections[0].name, "")
        XCTAssertEqual(r.sections[0].ingredients.count, 3)
    }

    func testMixedInstructionsAndIngredients() {
        guard let r = recipe("soup:\n1 onion\nDice finely\n2 cups broth\nSimmer for 20 min") else {
            return XCTFail("Expected recipe")
        }
        XCTAssertEqual(r.sections[0].ingredients.count, 2)
        XCTAssertEqual(r.sections[0].instructions.count, 2)
    }

    // MARK: - Cross-parser

    func testWorkoutNotRecipe() {
        let result = snap("squat 3x8x135")
        if let result = result {
            if case .recipe = result {
                XCTFail("Workout should not parse as recipe")
            }
        }
    }
}
