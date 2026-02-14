/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "workout",

  extras: ($) => [/[ \t]/],

  conflicts: ($) => [],

  rules: {
    // A paragraph is one or more newline-separated lines.
    // The top-level rule uses newline as a line separator/terminator.
    paragraph: ($) =>
      seq(
        $._line_content,
        repeat(seq($._separator, optional($._line_content))),
      ),

    _separator: (_) => choice("\n", "."),

    // A line's content (without the newline) is one of these forms.
    _line_content: ($) =>
      choice(
        $.superset_line,
        $.exercise_line,
        $.continuation_line,
        $.prose_line,
      ),

    // A continuation line starts with a numeric token.
    continuation_line: ($) =>
      prec.left(2, seq($.numeric_token, repeat($.token))),

    // A superset line has two superset names (WORD+ *) followed by tokens.
    superset_line: ($) =>
      prec.left(3, seq($.superset_name, $.superset_name, repeat1($.token))),

    superset_name: ($) => prec.left(seq(repeat1($.word), $.star)),

    // An exercise line starts with words, then a numeric token, then more tokens.
    exercise_line: ($) =>
      prec.left(1, seq(repeat1($.word), $.numeric_token, repeat($.token))),

    // A prose line is just words.
    prose_line: ($) => prec.left(0, repeat1($.word)),

    numeric_token: ($) => choice($.by_expr, $.multiplier, $.number),

    token: ($) => choice($.numeric_token, $.note),

    by_expr: ($) => choice($.three_part_by, $.two_part_by),

    two_part_by: ($) =>
      prec(2, seq(
        field("reps", $.number_literal),
        $.by_keyword,
        field("weight", $.number_literal),
      )),

    three_part_by: ($) =>
      prec(3, seq(
        field("sets", $.number_literal),
        $.by_keyword,
        field("reps", $.number_literal),
        $.by_keyword,
        field("weight", $.number_literal),
      )),

    multiplier: ($) => choice($.full_multiplier, $.partial_multiplier),

    partial_multiplier: ($) =>
      seq(
        field("sets", $.number_literal),
        $.x_separator,
        field("reps", $.number_literal),
      ),

    full_multiplier: ($) =>
      prec(1, seq(
        field("sets", $.number_literal),
        $.x_separator,
        field("reps", $.number_literal),
        $.x_separator,
        field("weight", $.number_literal),
      )),

    number: ($) => $.number_literal,

    note: ($) => prec(-1, repeat1($.word)),

    number_literal: (_) => /[0-9]+/,
    x_separator: (_) => token(prec(2, /[xX]/)),
    by_keyword: (_) => token(prec(2, /[bB][yY]/)),
    star: (_) => "*",
    word: (_) => /[a-zA-Z]+/,
  },
});
