import CSnap

func parseRecipe(_ input: String) -> RecipeResult? {
    let cResult = UnsafeMutablePointer<CSnap.RecipeResult>.allocate(capacity: 1)
    UnsafeMutableRawPointer(cResult)
        .initializeMemory(as: UInt8.self, repeating: 0,
                          count: MemoryLayout<CSnap.RecipeResult>.size)
    defer { cResult.deallocate() }

    guard recipe_parse(input, cResult) != 0 else { return nil }

    let title: String = withUnsafePointer(to: &cResult.pointee.title) { ptr in
        ptr.withMemoryRebound(to: CChar.self, capacity: Int(RECIPE_NAME_MAX)) {
            String(cString: $0)
        }
    }

    let sectionCount = Int(cResult.pointee.section_count)
    var sections: [RecipeSection] = []

    withUnsafeMutablePointer(to: &cResult.pointee.sections) { secTuplePtr in
        secTuplePtr.withMemoryRebound(to: CSnap.RecipeSection.self, capacity: Int(RECIPE_MAX_SECTIONS)) { secBuf in
            for i in 0..<sectionCount {
                let cSec = secBuf + i

                let name: String = withUnsafePointer(to: &cSec.pointee.name) { ptr in
                    ptr.withMemoryRebound(to: CChar.self, capacity: Int(RECIPE_NAME_MAX)) {
                        String(cString: $0)
                    }
                }

                let ingCount = Int(cSec.pointee.ingredient_count)
                var ingredients: [Ingredient] = []

                withUnsafeMutablePointer(to: &cSec.pointee.ingredients) { ingTuplePtr in
                    ingTuplePtr.withMemoryRebound(to: CSnap.Ingredient.self, capacity: Int(RECIPE_MAX_INGREDIENTS)) { ingBuf in
                        for j in 0..<ingCount {
                            let cIng = ingBuf + j

                            let qty: String = withUnsafePointer(to: &cIng.pointee.quantity) { ptr in
                                ptr.withMemoryRebound(to: CChar.self, capacity: Int(RECIPE_QTY_MAX)) {
                                    String(cString: $0)
                                }
                            }

                            let text: String = withUnsafePointer(to: &cIng.pointee.text) { ptr in
                                ptr.withMemoryRebound(to: CChar.self, capacity: Int(RECIPE_NAME_MAX)) {
                                    String(cString: $0)
                                }
                            }

                            ingredients.append(Ingredient(quantity: qty, text: text))
                        }
                    }
                }

                let instrCount = Int(cSec.pointee.instruction_count)
                var instructions: [String] = []

                withUnsafeMutablePointer(to: &cSec.pointee.instructions) { instrTuplePtr in
                    let base = UnsafeRawPointer(instrTuplePtr).assumingMemoryBound(to: CChar.self)
                    for j in 0..<instrCount {
                        let linePtr = base.advanced(by: j * Int(RECIPE_TEXT_MAX))
                        instructions.append(String(cString: linePtr))
                    }
                }

                sections.append(RecipeSection(name: name, ingredients: ingredients, instructions: instructions))
            }
        }
    }

    return RecipeResult(title: title, sections: sections)
}
