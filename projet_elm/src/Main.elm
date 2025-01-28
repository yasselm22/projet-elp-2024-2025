module Main exposing (..)

import Browser
import Html exposing (Html, div, text, textarea, pre)
import Html.Attributes exposing (style, value)
import Html.Events exposing (onInput)
import Parsing exposing (Command, programParser)  -- Importez programParser depuis Parsing.elm
import Parser
--import Svg exposing (..)
--import Svg.Attributes exposing (..)

-- MODEL
type alias Model =
    { input : String
    , commands : Result String (List Command)  -- Stocke le résultat du parsing (succès/erreur)
    }

init : Model
init =
    { input = "Exemple :[Forward 100, Repeat 4 [Forward 50, Left 90], Forward 100]"  -- C'est meiux le champ vide au départ ?
        , commands = Ok []
        }

-- UPDATE
type Msg
    = InputChanged String

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        InputChanged newInput ->
            let
                parsedCommands = 
                    Parser.run programParser newInput
                        |> Result.mapError (\_ -> "Erreur de syntaxe")
                _ = Debug.log "Parsed Commands" parsedCommands
            in
            ( { model 
                | input = newInput
                , commands = parsedCommands
              }
            , Cmd.none
            )


-- VIEW
view : Model -> Html Msg
view model =
    div [ style "padding" "20px" ]
        [ -- Zone de saisie
          textarea 
            [ value model.input
            , onInput InputChanged
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

-- RENDER COMMANDS
--renderCommands : String -> List (Svg msg)
--renderCommands commands =
    --if commands /= "" then
        --[ circle [ cx "250", cy "250", r "50", fill "red" ] [] ]
    --else
        --[]
        

-- MAIN
main : Program () Model Msg
main =
    Browser.element
        { init = \_ -> ( init, Cmd.none )
        , update = update
        , view = view
        , subscriptions = \_ -> Sub.none
        }
--main =
    --Browser.sandbox
        --{ init = init
        --, update = update
        --, view = view
        --}
