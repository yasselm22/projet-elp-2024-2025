module Main exposing (..)

import Browser
import Html exposing (Html, div, input, text)
import Html.Attributes exposing (..)
import Html.Events exposing (onInput)
import Svg exposing (..)
import Svg.Attributes exposing (..)

-- MODEL
type alias Model =
    { commands : String }

init : Model
init =
    { commands = "" }


-- UPDATE
type Msg
    = UpdateCommands String

update : Msg -> Model -> Model
update msg model =
    case msg of
        UpdateCommands newCommands ->
            { model | commands = newCommands }


-- VIEW
view : Model -> Html Msg
view model =
    div []
        [ input 
            [ placeholder "Entrez des commandes de tracé"
            , onInput UpdateCommands
            , ariaLabel "Champ de commandes"
            ] 
            []
        , svg [ width "500", height "500", viewBox "0 0 500 500" ]
            (renderCommands model.commands)
        ]

-- RENDER COMMANDS
renderCommands : String -> List (Svg msg)
renderCommands commands =
    if commands /= "" then
        [ circle [ cx "250", cy "250", r "50", fill "red" ] [] ]
    else
        []
        

-- MAIN
main =
    Browser.sandbox { init = init, update = update, view = view }
