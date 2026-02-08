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

// "8 by 35" (reps by weight, default 3 sets)
// "2 by 8 by 35" (sets by reps by weight)
byExpr
    : reps=NUMBER BY weight=NUMBER                      # TwoPartBy
    | sets=NUMBER BY reps=NUMBER BY weight=NUMBER       # ThreePartBy
    ;

// "3x8" (sets x reps, weight comes next)
// "3x8x135" (sets x reps x weight)
multiplier
    : sets=NUMBER X reps=NUMBER                         # PartialMultiplier
    | sets=NUMBER X reps=NUMBER X weight=NUMBER         # FullMultiplier
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
