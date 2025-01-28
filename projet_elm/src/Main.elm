module Main exposing (..)

import Browser
import Html exposing (Html, div, text, textarea, button, input, pre)
import Html.Attributes exposing (style, value, placeholder)
import Html.Events exposing (onInput, onClick)
import Parsing exposing (Command(..), programParser)
import Parser
import Svg exposing (Svg, svg, line, circle)
import Svg.Attributes exposing (viewBox, width, height, x1, y1, x2, y2, stroke, cx, cy, r, fill)

-- MODEL
type alias Model =
    { input : String
    , commands : Result String (List Command)  -- Stocke le résultat du parsing (succès/erreur)
    , drawing : List (Svg Msg)  -- Liste des éléments SVG à dessiner
    }

init : Model
init =
    { input = ""
    , commands = Ok []
    , drawing = []
    }

-- UPDATE
type Msg
    = InputChanged String
    | Draw

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        InputChanged newInput ->
            let
                parsedCommands =
                    Parser.run programParser newInput
                        |> Result.mapError (\_ -> "Erreur de syntaxe")
            in
            ( { model
                | input = newInput
                , commands = parsedCommands
              }
            , Cmd.none
            )

        Draw ->
            case model.commands of
                Ok commands ->
                    ( { model
                        | drawing = renderCommands commands
                      }
                    , Cmd.none
                    )

                Err errorMsg ->
                    ( { model
                        | drawing = [text errorMsg]
                      }
                    , Cmd.none
                    )

{-
-- VIEW
view : Model -> Html Msg
view model =
    div [ style "padding" "20px" ]
        [ -- Zone de saisie
          textarea 
            [ value model.input -- La valeur actuelle de la zone de texte est liée au champ input du modèle. Cela garantit que la zone de texte affiche toujours la valeur stockée dans le modèle
            , onInput InputChanged -- Cet événement déclenche un message InputChanged à chaque fois que l'utilisateur modifie le texte dans la zone. Le message contiendra le nouveau texte saisi par l'utilisateur.
            , style "width" "400px"
            , style "height" "150px"
            , style "margin-bottom" "20px"
            ] 
            []
          
          -- Affichage des commandes parsées (pour débogage)
        , case model.commands of
            Ok commands ->
                div []
                    [ text "Commandes parsées avec succès :"
                    , pre [] [ text (Debug.toString commands) ]  -- Affiche la structure Elm
                    ]

            Err errorMsg ->
                div [ style "color" "red" ] [ text errorMsg ]
        ]

-- VIEW
view : Model -> Html Msg
view model =
    div []
        [ input
            [ placeholder "example: [Repeat 360 [Forward 1, Left 1]]"
            , value model.input
            , onInput InputChanged
            ]
            []
        , button [ onClick Draw ] [ text "Draw" ]
        , svg
            [ viewBox "0 0 500 500"
            , width "500"
            , height "500"
            ]
            model.drawing
        ]-}


-- VIEW
view : Model -> Html Msg
view model =
    div [ style "padding" "20px" ]
        [ -- Zone de saisie
          input
            [ placeholder "example: [Repeat 360 [Forward 1, Left 1]]"
            , value model.input
            , onInput InputChanged
            , style "width" "400px"
            , style "margin-bottom" "20px"
            ]
            []
          
          -- Bouton pour dessiner
        , button
            [ onClick Draw
            , style "margin-bottom" "20px"
            ]
            [ text "Draw" ]
          
          -- Zone de dessin SVG
        , svg
            [ viewBox "0 0 500 500"
            , width "500"
            , height "500"
            ]
            model.drawing
          
          -- Affichage des commandes parsées pour le débogage
        , case model.commands of
            Ok commands ->
                div []
                    [ text "Commandes parsées avec succès :"
                    , pre [] [ text (Debug.toString commands) ]
                    ]

            Err errorMsg ->
                div [ style "color" "red" ] [ text errorMsg ]
        ]



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
        initialTurtle = { x = 250, y = 250, angle = 0 }  -- Position initiale au centre du SVG
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

        Repeat _ _ ->
            (turtle, elements)  -- Ignore Repeat pour l'instant



-- RENDER COMMANDS
--renderCommands : String -> List (Svg msg)
--renderCommands commands =
    --if commands /= "" then
        --[ circle [ cx "250", cy "250", r "50", fill "red" ] [] ]
    --else
        --[]
        

-- MAIN : Ce code Elm est la configuration principale de votre application. Il utilise la fonction Browser.element pour définir une application Elm qui s'intègre dans le navigateur.
main : Program () Model Msg -- () : Représente les données initiales fournies par l'environnement JavaScript (généralement ignorées ici, d'où le type vide ()).
-- Model : Le type du modèle qui représente l'état de l'application.
-- Msg : Le type des messages qui représentent les événements pouvant modifier l'état.

main =
    Browser.element -- Permet de créer une application Elm qui interagit directement avec le DOM du navigateur.
        { init = \_ -> ( init, Cmd.none ) -- init : Initialise l'état de l'application.
        , update = update -- update : Met à jour le modèle en réponse à des messages.
        , view = view -- view : Génère la vue (HTML) basée sur le modèle.
        , subscriptions = \_ -> Sub.none -- subscriptions : Définit les abonnements aux événements externes
        }
--main =
    --Browser.sandbox
        --{ init = init
        --, update = update
        --, view = view
        --}
