# push-swap

Projet en Go (bibliothèque standard uniquement) avec deux programmes :

- `push-swap` : affiche la liste d'instructions qui trie la pile `a`.
- `checker` : lit des instructions sur l'entrée standard, les exécute, puis affiche `OK` ou `KO`.

## Compiler

```
go build -o push-swap ./programme/push-swap
go build -o checker ./programme/checker
```

Sous Windows, ajoute `.exe` : `go build -o push-swap.exe ./programme/push-swap`.

## Utiliser

```
$ ./push-swap "4 67 3 87 23"
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
$ ./push-swap "0 one 2 3"
Error
```

## Organisation

| Dossier | Rôle |
|---|---|
| `parsing/` | Transforme les arguments en liste d'entiers et détecte les erreurs (pas un entier, doublon). |
| `instructions/` | Les deux piles et les 11 instructions (`sa`, `pb`, `rra`...). |
| `tri/` | L'algorithme qui choisit les instructions. |
| `programme/checker/` | Le programme `checker`. |
| `programme/push-swap/` | Le programme `push-swap`. |

## Algorithme

- Liste déjà triée : aucune instruction.
- 2 et 3 nombres : règles simples (on envoie le plus grand en bas, puis `sa` si besoin).
- 4 à 6 nombres : parcours en largeur. On essaie toutes les suites de 1 coup, puis de 2 coups, etc. La première qui trie est la plus courte.
- Plus de 6 nombres : tri par paquets. On envoie les nombres dans `b` par paquets de petits nombres, puis on ramène toujours le plus grand de `b` en haut de `a`.

## Tests

```
go test ./...
```
