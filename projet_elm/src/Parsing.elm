module Parsing exposing (Command(..), programParser)

import Parser exposing (Parser, (|.), (|=), succeed, symbol, int, float, spaces, lazy, sequence)

type Command
    = Forward Float
    | Left Float
    | Right Float
    | Repeat Int (List Command)

commandParser : Parser Command
commandParser =
    Parser.oneOf
        [ repeatParser  -- Tester d'abord les commandes complexes
        , forwardParser
        , leftParser
        , rightParser   
        ]

forwardParser : Parser Command
forwardParser =
    succeed Forward
        |. symbol "Forward"
        |. spaces
        |= float

leftParser : Parser Command
leftParser =
    succeed Left
        |. symbol "Left"
        |. spaces
        |= float

rightParser : Parser Command
rightParser =
    succeed Right
        |. symbol "Right"
        |. spaces
        |= float

repeatParser : Parser Command
repeatParser =
    succeed Repeat
        |. symbol "Repeat"
        |. spaces
        |= int
        |. spaces
        |. symbol "["
        |= lazy (\_ -> sequence { start = "", separator = ",", end = "", spaces = spaces, item = commandParser, trailing = Parser.Forbidden }) -- Sans lazy, on aurait une récursion infinie car commandParser dépend de lui-même (pour les commandes imbriquées)
        |. symbol "]"

-- Pour parser une liste complète de commandes
programParser : Parser (List Command)
programParser =
    sequence
        { start = "["
        , separator = ","
        , end = "]"
        , spaces = spaces
        , item = commandParser
        , trailing = Parser.Forbidden
        }


-- Faire un message d'erreur si input autre chose que les commandes données
