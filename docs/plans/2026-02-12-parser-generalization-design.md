# Parser Generalization Design

## Context

WorkoutLiner is one of several "special type" parsers in an iOS app (and future emacs integration). The app detects paragraph breaks (double-enter), runs parsers on the preceding paragraph, and overlays a rendered view on top of the plain text. Backspace dissolves the overlay back to plain text. Plain text is always the source of truth.

Future parsers include: recipes, bucket lists, TODO lists, calendar events.

## Design

### Parser Library Contract

Each parser is a standalone module with one job:

```
paragraph: String -> TypeResult? (nil = not my type)
```

Parsers are pure functions: text in, structured data out. They know nothing about search bars, overlays, rendering, or platforms. They share ANTLR grammars across Swift/Go but are otherwise independent packages.

No shared protocol or interface across parsers. Each parser defines its own return type. The return types are fundamentally different (workout sets vs recipe ingredients vs todo items), and the app already needs type-specific code for rendering, search actions, and index mapping. Polymorphism adds indirection without value here.

### App Responsibilities

The consuming app (iOS, emacs) handles:

- Paragraph splitting (split on `\n\n`)
- Running paragraphs through registered parsers
- Building the KV search index from parser output
- Defining per-type search actions
- Rendering per-type overlays

### Search Index

The app maps each parser's domain objects to a uniform `SearchEntry`:

```swift
SearchEntry(
    key: item.name.lowercased(),  // lookup key
    name: item.name,              // display name
    rawLine: item.rawLine,        // original text
    type: "workout"               // discriminator for actions
)
```

The app adds `date` and `filePath` from context. The `type` field tells the search bar which actions to show.

### WorkoutParser Changes

- Rename `ParseResult` to `WorkoutResult`
- Remove `ParagraphResult` enum (app's concern)
- Rename `transform` to `parse`, return `WorkoutResult?` instead of `[ParagraphResult]`
- Remove paragraph splitting logic from the parser
- `WorkoutResult` keeps `exercises: [Exercise]` and `proseLines: [String]`

### Repo Structure

Rename repo from `workoutliner` to reflect multi-parser purpose.

```
repo/
  WorkoutParser/
    grammar/        <- ANTLR .g4
    swift/          <- Swift package
    go/             <- Go module
  RecipeParser/     <- future
  TodoParser/       <- future
```

### What Stays the Same

- ANTLR grammar, visitor logic, models (Exercise, WorkoutSet)
- The Go implementation (same renames applied)
- All existing parsing behavior

### What Moves to the App

- Paragraph splitting (split on `\n\n`)
- Prose vs structured data classification
- Search index building
- Rendering and search actions
