import CSnap

/// Parse a workout string into a structured WorkoutResult.
/// Returns nil if the input is not recognized as a workout.
public func parse(_ input: String) -> WorkoutResult? {
    // Heap-allocate the C result struct (~544KB) to avoid stack overflow
    let cResult = UnsafeMutablePointer<CSnap.WorkoutResult>.allocate(capacity: 1)
    UnsafeMutableRawPointer(cResult)
        .initializeMemory(as: UInt8.self, repeating: 0,
                          count: MemoryLayout<CSnap.WorkoutResult>.size)
    defer { cResult.deallocate() }

    guard workout_parse(input, cResult) != 0 else { return nil }

    let exerciseCount = Int(cResult.pointee.exercise_count)
    let proseCount = Int(cResult.pointee.prose_count)

    // Extract exercises using pointer arithmetic on the heap-allocated struct
    var exercises: [Exercise] = []

    withUnsafeMutablePointer(to: &cResult.pointee.exercises) { exTuplePtr in
        exTuplePtr.withMemoryRebound(to: CSnap.Exercise.self, capacity: Int(WORKOUT_MAX_EXERCISES)) { exBuf in
            for i in 0..<exerciseCount {
                let cExPtr = exBuf + i

                // Extract name
                let name: String = withUnsafePointer(to: &cExPtr.pointee.name) { namePtr in
                    namePtr.withMemoryRebound(to: CChar.self, capacity: Int(WORKOUT_NAME_MAX)) {
                        String(cString: $0)
                    }
                }

                let setCount = Int(cExPtr.pointee.set_count)
                var sets: [WorkoutSet] = []

                withUnsafeMutablePointer(to: &cExPtr.pointee.sets) { setsTuplePtr in
                    setsTuplePtr.withMemoryRebound(to: CSnap.WorkoutSet.self, capacity: Int(WORKOUT_MAX_SETS)) { setBuf in
                        for j in 0..<setCount {
                            let cSetPtr = setBuf + j

                            let note: String = withUnsafePointer(to: &cSetPtr.pointee.note) { notePtr in
                                notePtr.withMemoryRebound(to: CChar.self, capacity: Int(WORKOUT_NOTE_MAX)) {
                                    String(cString: $0)
                                }
                            }

                            sets.append(WorkoutSet(
                                reps: Int(cSetPtr.pointee.reps),
                                weight: Int(cSetPtr.pointee.weight),
                                note: note
                            ))
                        }
                    }
                }

                exercises.append(Exercise(name: name, sets: sets))
            }
        }
    }

    // Extract prose lines
    var proseLines: [String] = []
    if proseCount > 0 {
        withUnsafeMutablePointer(to: &cResult.pointee.prose_lines) { proseTuplePtr in
            let base = UnsafeRawPointer(proseTuplePtr).assumingMemoryBound(to: CChar.self)
            for i in 0..<proseCount {
                let linePtr = base.advanced(by: i * Int(WORKOUT_NAME_MAX))
                proseLines.append(String(cString: linePtr))
            }
        }
    }

    return WorkoutResult(proseLines: proseLines, exercises: exercises)
}
