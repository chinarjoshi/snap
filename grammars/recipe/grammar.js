/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "recipe",

  extras: ($) => [/[ \t]/],

  conflicts: ($) => [
    [$.ingredient_line, $.text_line],
    [$.header_line, $.text_line],
    [$.quantity, $.text_line],
  ],

  rules: {
    recipe: ($) =>
      seq(
        $._line_content,
        repeat(seq("\n", optional($._line_content))),
      ),

    _line_content: ($) =>
      choice(
        $.header_line,
        $.ingredient_line,
        $.text_line,
      ),

    header_line: ($) => prec.dynamic(10, seq(repeat1($.word), ":")),

    ingredient_line: ($) =>
      prec.dynamic(5, seq($.quantity, repeat1(choice($.word, $.number)))),

    text_line: ($) =>
      prec.dynamic(0, repeat1(choice($.word, $.number, $.punctuation))),

    quantity: ($) =>
      choice(
        seq(field("min", $.number), "-", field("max", $.number)),
        field("value", $.number),
      ),

    word: (_) => /[a-zA-Z]+/,
    number: (_) => /[0-9]+/,
    punctuation: (_) => /[^\s\na-zA-Z0-9:]+/,
  },
});
