## Importações

ˋˋimport "fmt"ˋˋ -> importa o pacote fmt
com a variavel fmt

referencia js:

ˋˋimport fmt from "fmt"ˋˋ

ˋˋimport (. "fmt")ˋˋ

importa tudo de fmt separadamente

referencia js:

ˋˋimport \* from "fmt"ˋˋ

import (\_ "fmt")
-> Carrega todo o pacote fmt antes de carregar o pacote atual

ˋˋimport (meuPacote "fmt")ˋˋ
-> renomeia a variavel de import

referencia js:

ˋˋimport ftm as meuPacote from "fmt"ˋˋ

Em go o que é exportado é o pacote e não um arquivo em si

ˋˋˋ
import (
"fmt"
"myFirstGoProject/pacote"
)
ˋˋˋ
o import começa com o nome do projeto "myFirstGoProject" e depois o nome do pacote "pacote"

apois esse import todas as variavies e funções publicas do pacote "pacote" serão exibidas

OBS: Se tentar declarar o mesmo nome dentro de arquivos diferentes mais no mesmo pacote, será exibido erro de duplicação de nome (variavel, funcao ou afim)

## Nomes

### privacidade

nomes iniciado em maisculos são publicos

ˋˋˋ
package example
var Foo string

##

import "example"

example.Foo

ˋˋˋ

nomes iniciado com minusculo são privados

ˋˋˋ
package example
var foo string

##

import "example"

# example.foo # nao existe

ˋˋˋ

todos os arquivos dentro de um mesmo diretorio devem estar no mesmo pacote

pacote
-> a.go -> package pacote
-> b.go -> package pacote
-> c.go -> package outro # vai da erro deve percentencer ao package pacote

## Variaveis

### internal

pacote
-> internal
    -> foo -> package foo
        -> variavel
-> a.go -> package pacote
-> b.go -> package pacote
-> c.go -> package pacote

a pasta internal é utilizada para funções e/ou variaveis que serão usadas dentro do mesmo pacote.

porem se tentar importar algo da pastar internal fora do pacote, irá disparar um erro

### main

não é possivel importar nada do pacote main, ele serve apenas para criar o executavel
