# push-swap

## Présentation

push-swap est un projet en Go qui trie une liste d'entiers avec deux piles, `a` et `b`, et un jeu de 11 instructions.
Le but est de trouver la suite d'instructions la plus courte possible.

Le projet contient deux programmes :

- `push-swap` calcule et affiche les instructions qui trient la pile `a`.
- `checker` lit des instructions, les exécute sur la pile `a`, puis dit si le tri est réussi.

## Fonctionnalités

- Lire une liste d'entiers en arguments (`"3 2 1"` ou `3 2 1`). Le premier nombre est le sommet de la pile `a`.
- Détecter les erreurs : argument qui n'est pas un entier, doublon, instruction inconnue. Le programme affiche alors `Error` sur la sortie d'erreur.
- Les 11 instructions : `sa`, `sb`, `ss`, `pa`, `pb`, `ra`, `rb`, `rr`, `rra`, `rrb`, `rrr`.
- `push-swap` : afficher une instruction par ligne. Si la liste est déjà triée, rien n'est affiché.
- `checker` : afficher `OK` si `a` est triée et `b` vide, sinon `KO`.
- Sans argument, les deux programmes n'affichent rien.

## Technologies

- Go, avec seulement la bibliothèque standard.

## Prérequis

- Go 1.21 ou plus récent (le package `slices` est utilisé).

## Installation

```
git clone https://github.com/BoWhite50/push-swap.git
cd push-swap
go build -o push-swap ./programme/push-swap
go build -o checker ./programme/checker
```

Sous Windows, ajoute `.exe` au nom : `go build -o push-swap.exe ./programme/push-swap`.

## Configuration

Aucune configuration ni variable d'environnement n'est nécessaire.

## Lancement

```
$ ./push-swap "2 1 3"
sa
$ ./push-swap "0 one 2 3"
Error
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
6
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
$ ./checker "3 2 1 0"
sa
rra
pb
KO
```

Dans `checker`, les instructions sont tapées une par ligne. Pour finir la saisie : `Ctrl+D` (Linux/Mac) ou `Ctrl+Z` puis `Entrée` (Windows).

## Tests

```
go test ./...
```

Les tests vérifient les 11 instructions, la lecture des arguments (erreurs comprises) et le tri : tous les ordres possibles jusqu'à 5 nombres, les exemples du sujet, et des listes au hasard de 100 et 500 nombres.

## Architecture

```
push-swap/
├── parsing/        transforme les arguments en liste d'entiers et détecte les erreurs
├── instructions/   les deux piles et les 11 instructions
├── tri/            l'algorithme qui choisit les instructions
└── programme/
    ├── checker/    le programme checker
    └── push-swap/  le programme push-swap
```

L'algorithme de tri dépend de la taille de la liste :

| Taille | Méthode |
|---|---|
| déjà triée | aucune instruction |
| 2 ou 3 nombres | règles simples : le plus grand va en bas, puis `sa` si besoin |
| 4 à 6 nombres | parcours en largeur : on essaie toutes les suites de 1 coup, puis de 2 coups, etc. La première qui trie est la plus courte. |
| plus de 6 nombres | tri par paquets : on envoie les nombres dans `b` par paquets de petits nombres, puis on ramène toujours le plus grand de `b` en haut de `a` |

Nombre moyen d'instructions : environ 575 pour 100 nombres, environ 5 050 pour 500 nombres.
