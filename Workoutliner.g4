grammar Workoutliner;

paragraph : line+ EOF ;

line
    : supersetLine NEWLINE?     # SupersetLineAlt
    | exerciseLine NEWLINE?     # ExerciseLineAlt
    | continuationLine NEWLINE? # ContinuationLineAlt
    | proseLine NEWLINE?        # ProseLineAlt
    ;

continuationLine : numericToken token* ;

supersetLine : supersetName supersetName token+ ;

supersetName : WORD+ STAR ;

exerciseLine : WORD+ numericToken token* ;

proseLine : WORD+ ;

numericToken
    : byExpr
    | multiplier
    | number
    ;

token
    : numericToken
    | note
    ;

byExpr
    : reps=NUMBER BY weight=NUMBER                # TwoPartBy
    | sets=NUMBER BY reps=NUMBER BY weight=NUMBER # ThreePartBy
    ;

multiplier
    : sets=NUMBER X reps=NUMBER                   # PartialMultiplier
    | sets=NUMBER X reps=NUMBER X weight=NUMBER   # FullMultiplier
    ;

number : NUMBER ;

note : WORD+ ;

NUMBER  : [0-9]+ ;
X       : [xX] ;
BY      : [bB][yY] ;
STAR    : '*' ;
WORD    : [a-zA-Z]+ ;
NEWLINE : '\n' ;
WS      : [ \t]+ -> skip ;
