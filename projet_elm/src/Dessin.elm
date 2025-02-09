module Dessin exposing (Turtle, renderCommands)

import Svg exposing (Svg, line)
import Svg.Attributes exposing (x1, y1, x2, y2, stroke)
import Parsing exposing (Command(..))  -- Importation des commandes (Forward, Left, Right, Repeat)


-- Structure pour le "Turtle Graphics" 
type alias Turtle =
    { x : Float
    , y : Float
    , angle : Float  -- En degrés
    }


-- Fonction pour calculer la nouvelle position après un Forward
moveForward : Float -> Turtle -> Turtle
moveForward distance turtle =
    { turtle
        | x = turtle.x + distance * cos (degrees turtle.angle)
        , y = turtle.y + distance * sin (degrees turtle.angle)
    }


-- Fonctions pour calculer le nouvel angle après un Left ou un Right
turnLeft : Float -> Turtle -> Turtle
turnLeft angle turtle =
    { turtle | angle = turtle.angle - angle }

turnRight : Float -> Turtle -> Turtle
turnRight angle turtle =
    { turtle | angle = turtle.angle + angle }


renderCommands : List Command -> List (Svg msg)
renderCommands commands =
    let
        initialTurtle = { x = 250, y = 250, angle = 0 }  -- Position initiale au centre
        (_, svgElements) =
            List.foldl renderCommand (initialTurtle, []) commands 
    in
    svgElements


-- RENDER COMMANDS
renderCommand : Command -> (Turtle, List (Svg msg)) -> (Turtle, List (Svg msg))
renderCommand command (turtle, elements) =
    case command of
        Forward distance ->
            let
                newTurtle = moveForward distance turtle
                lineElement =
                    line
                        [ x1 (String.fromFloat turtle.x)
                        , y1 (String.fromFloat turtle.y)
                        , x2 (String.fromFloat newTurtle.x)
                        , y2 (String.fromFloat newTurtle.y)
                        , stroke "black"
                        ]
                        []
            in
            (newTurtle, lineElement :: elements)

        Left angle ->
            (turnLeft angle turtle, elements)

        Right angle ->
            (turnRight angle turtle, elements)

        Repeat n subCommands ->
            let
                -- Répète les sous-commandes n fois
                repeatedCommands = List.repeat n subCommands |> List.concat
            in
            -- Applique les commandes répétées
            List.foldl renderCommand (turtle, elements) repeatedCommands