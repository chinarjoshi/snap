grammar Paragraph;

paragraph : line+ EOF ;

line
    : exerciseLine NEWLINE?   # ExerciseLineAlt
    | proseLine NEWLINE?      # ProseLineAlt
    ;

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
WORD    : [a-zA-Z]+ ;
NEWLINE : '\n' ;
WS      : [ \t]+ -> skip ;
