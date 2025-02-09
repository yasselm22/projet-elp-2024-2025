module Main exposing (..)

import Browser
import Html exposing (Html, div, text, textarea, button, input, pre)
import Html.Attributes exposing (style, value, placeholder)
import Html.Events exposing (onInput, onClick)
import Parsing exposing (Command(..), programParser)
import Dessin exposing (renderCommands)
import Parser
import Svg exposing (Svg, svg, line, circle)
import Svg.Attributes exposing (viewBox, width, height, x1, y1, x2, y2, stroke, cx, cy, r, fill)

-- MODEL
type alias Model =
    { input : String
    , commands : Result String (List Command)  -- Stocke le résultat du parsing
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
                        | drawing = Dessin.renderCommands commands
                      }
                    , Cmd.none
                    )

                Err errorMsg ->
                    ( { model
                        | drawing = [text errorMsg]
                      }
                    , Cmd.none
                    )

-- VIEW
view : Model -> Html Msg
view model =
    div [ style "padding" "5px", style "display" "flex", style "flex-direction" "column", style "align-items" "center" ]
    -- active le mode Flexbox, pour organiser les éléments dans l'interface, éléments affichés en colonne, éléments centrés
        [ -- Zone de saisie et bouton Draw dans un conteneur
          div [ style "margin-bottom" "10px", style "text-align" "center" ]
            [ input
                [ placeholder "example: [Repeat 360 [Forward 1, Left 1]]"
                , value model.input
                , onInput InputChanged
                , style "width" "500px"
                , style "margin-bottom" "10px"
                ]
                []
            , button
                [ onClick Draw
                , style "display" "block"  -- prend seulement la largeur de son contenu
                , style "margin" "auto"  -- centre l'élément horizontalement, mais seulement si l'élément est en display: block ou flex
                ]
                [ text "Draw" ]
            ]
          
          -- Affichage des commandes parsées pour le débogage
        , case model.commands of
            Ok commands ->
                div [ style "margin-bottom" "10px", style "text-align" "center" ]
                    [ text "Commandes parsées avec succès :"
                    , pre [] [ text (Debug.toString commands) ]
                    ]

            Err errorMsg ->
                div [ style "color" "red", style "margin-bottom" "10px", style "text-align" "center" ]
                    [ text errorMsg ]

          -- Zone de dessin SVG
        , svg
            [ viewBox "0 0 500 500"
            -- , Svg.Attributes.width "500"
            -- , Svg.Attributes.height "500"
            , width "500"
            , height "500"
            , style "border" "1px solid black"
            ]
            model.drawing
        ]


-- MAIN
main : Program () Model Msg
main =
    Browser.element -- Crée une application Elm qui interagit directement avec le DOM du navigateur
        { init = \_ -> ( init, Cmd.none )
        , update = update
        , view = view
        , subscriptions = \_ -> Sub.none
        }
