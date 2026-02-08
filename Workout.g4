grammar Workout;

// Parser rules
workoutLog
    : exercise (DELIMITER exercise)* EOF
    ;

exercise
    : exerciseName token*
    ;

exerciseName
    : WORD+
    ;

token
    : byExpr
    | multiplier
    | number
    | note
    ;

// "8 by 35" or "2 by 8 by 35"
byExpr
    : NUMBER BY NUMBER (BY NUMBER)?
    ;

// "3x8" or "3x8x135"
multiplier
    : NUMBER X NUMBER (X NUMBER)?
    ;

number
    : NUMBER
    ;

note
    : WORD+
    ;

// Lexer rules (order matters - more specific first)
NUMBER : [0-9]+ ;
X      : [xX] ;
BY     : [bB][yY] ;
WORD   : [a-zA-Z]+ ;
DELIMITER : [.\n]+ ;
WS     : [ \t]+ -> skip ;
