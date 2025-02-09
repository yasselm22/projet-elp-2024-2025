module Parsing exposing (Command(..), programParser)

import Parser exposing (Parser, (|.), (|=), succeed, symbol, int, float, spaces, lazy, sequence)

-- Définition des types de commandes possibles
type Command
    = Forward Float
    | Left Float
    | Right Float
    | Repeat Int (List Command)

-- Parser principal pour une commande
commandParser : Parser Command
commandParser =
    Parser.oneOf
        [ repeatParser  -- Tester d'abord les commandes complexes
        , forwardParser
        , leftParser
        , rightParser   
        ]

-- Parser pour la commande "Forward"
forwardParser : Parser Command
forwardParser =
    succeed Forward
        |. symbol "Forward"
        |. spaces
        |= float

-- Parser pour la commande "Left"
leftParser : Parser Command
leftParser =
    succeed Left
        |. symbol "Left"
        |. spaces
        |= float

-- Parser pour la commande "Right"
rightParser : Parser Command
rightParser =
    succeed Right
        |. symbol "Right"
        |. spaces
        |= float

-- Parser pour la commande "Repeat"
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


