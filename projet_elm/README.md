# 🖌️ TC-LOGO : Application Web de Dessin en Elm

## Description
TC-LOGO est une application web permettant de **visualiser un dessin** généré à partir de **commandes de tracé TcTurtle** saisies par l'utilisateur.  
Ce projet est développé en **Elm**, un langage fonctionnel conçu pour les applications web.

## Fonctionnalités
- **Interprétation des commandes TcTurtle** (Forward, Left, Right, Repeat).
- **Génération de dessins dynamiques en SVG**.
- **Exécution optimisée via Elm Architecture**.
- **Interface utilisateur simple avec un champ de saisie et un bouton Draw**.

## Prérequis
Avant de commencer, assure-toi d’avoir installé :
- [Elm](https://elm-lang.org/) (compilateur)
- Un navigateur web moderne (Chrome, Firefox, Edge)

## Installation & Exécution
**Clone le projet avec la commande suivante :**  
```Bash
git clone https://github.com/yasselm22/projet-elp-2024-2025.git  
cd ./projet_elm/src  
```

**Pour compiler le fichier Elm en JavaScript, utilise la commande suivante :**  
```Bash
elm make Main.elm --output=main.js  
```

**Ouvre le fichier index2.html dans ton navigateur pour voir l'application Elm en action.**

## Syntaxe TcTurtle
L'application interprète le langage TcTurtle. Les dessins sont générés à partir de commandes comme :  
[Repeat 6 [Repeat 6 [Forward 100, Right 60], Repeat 12 [Forward 50, Left 150], Right 60]]  
[Repeat 36 [Repeat 12 [Forward 50, Left 150],Right 10]]  
[Repeat 180 [Forward 5, Left 10, Repeat 6 [Forward 60, Right 60]]]  

## Contributions
Les contributions sont les bienvenues !

## Auteurs
Collet Marine  
El Moukri Yassmine  
Muccini Bianca  