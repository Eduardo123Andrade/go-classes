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

### Funções

ˋˋˋ
func somar(a int, b int) int {
return a + b
}
ˋˋˋ

func -> define que é uma função
somar -> nome da funcao
(a int, b int) -> parametros da funcao
int -> tipo de retorno da funcao

caso o tipo dos parametros sejam o mesmo pode ser escrito assim

ˋˋˋ
func somar(a, b int) int {
return a + b
}
ˋˋˋ

nesse caso a e b sao do tipo inteiro
essa abreviação o tipo deve estar no ultimo parametro.

Funções em Go podem retornar mais de um valor

ˋˋˋ
func swap(a, b int) (int, int) {
return b, a
}
ˋˋˋ

ˋˋˋ
func main() {
a, b := swap(10, 20)
fmt.Println(a, b) # a -> 20, b -> 10
}
ˋˋˋ

ˋˋˋ
func dividir(a, b int) (res int, rem int) {
res := a / b
rem := a % b
return res, rem
}
ˋˋˋ
a variavei res e rem sao declaradas na assinatura da função e tem seus valores atualizados no corpo

# naked return

ˋˋˋ
func dividir(a, b int) (res int, rem int) {
res := a / b
rem := a % b
return
}
ˋˋˋ

Como as variaveis foram declaradas na assinatura, e atualizadas no corpo o return vazio automaticamente retorna as variaveis res e rem respectivamente

Não é uma boa pratica. Em codigos grandes fica dificil de entender

### Atribuição

a = 10 // atribui o valor 10 a variavel a

a := 5 // atualiza o valor da variavel a para 5

### High order function

ˋˋˋ
func main(){
// x := somar(2)(1)
f := somar(2)
x := f(1)
fmt.Println(x) // 3
}

func somar(a int) func(int) int {
return func(b int) int {
return a + b
}
}
ˋˋˋ

Funções podem retornar outras funcoes ou receber outras funções como parametro

# clousures

ˋˋˋ
func somar(a int) func(int) int {
return func(b int) int {
return a + b
}
}
ˋˋˋ

são funções que capturam (usam) variaveis de um escopo acima do seu proprio.
Nesse caso a função que esta sendo retornada usa a variavel b, do proprio escopo e a variavel a do escopo acima

# funções anonimas

ˋˋˋ
func main(){
// função anonima
f := func(a, b int) int {
return a + b
}

    x := f(10, 20)
    y := subtrair(30, 50)
    fmt.Println(x, y)

}

    // função nomeada
    func subtrair(a, b int) int {
    	return a - b
    }

ˋˋˋ
São funções que não tem nome, nesse caso atribuimos a função anonima a variavel f e com isso, f passa a ser uma função

referencia js:

ˋˋˋ
const soma = (a, b) => a + b
ˋˋˋ
OBS: Não é possivel declarar uma função nomeada dentro de outra função

# Argumentos variaveis

ˋˋˋ
func sum(nums ...int) int {
var out int
for \_, n := range nums {
out += n
}

    return out

}
ˋˋˋ

a variavel nums pode ter 0 ou N valores, será tratado como "array" e deve ser o ultimo argumento da função
