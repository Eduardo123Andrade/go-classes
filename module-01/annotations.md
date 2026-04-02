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

    palavra reservada 'var' para declarar uma lista de variaveis

    ```
    	var nome, sobrenome string
    ```


    variavels de escopo de funcao precisam ser utilizadas, caso náo, vai ocorrer um erro de run time

    ```
        ./main.go:6:6: nome declared and not used
        ./main.go:6:12: sobrenome declared and not used
    ```

    no entando variavies de pacote não precisam ser utilizadas

    ```
    package main

    import "fmt"

    var idate int

    func main() {
        var nome, sobrenome string

        fmt.Println(nome, sobrenome)

    }
    ```

toda variavel é inicializada com o seu valor default caso não seja definido explicitamente

em go o tipo da variavel pode ser infreido de acordo com seu valor

    ```

    # o compilador entende que é do tipo string
    var nome, sobrenome = "Eduardo", "Andrade"

    ```

as variaveis podem ser declaradas de modo agrupado

    ```
        var (
            nome =  "Eduardo"
            sobrenome = "Andrade"
            idade = 29
        )
    ```

variaveis podem ser declaradas sem o 'var'

    ```

    nome := "Eduardo"
    sobrenome = "Andrade"

    ```

ao utilizar := estamos declarando de definindo o valor da variavel ao mesmo tempo.
Se usarmos apenas o = estaremos apenas alterando o valor, nesse caso a variavel sobrenome nao existe, logo lanćará um erro

    ```
        nome := "Eduardo"
        sobrenome = "Andrade"
    ```

esse tipo de atribuicão só pode ser feito no escopo de funćão, fora dele causará erro

    ```
        package main

        import "fmt"

        var idade := 29

        func main() {

            nome := "Eduardo"
            sobrenome := "Andrade"

            fmt.Println(nome, sobrenome, idade)

        }

        # myFirstGoProject

        ./main.go:5:11: syntax error: unexpected :=, expected =
    ```

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

ao utilizar o := para declarar uma variavel, não é possivel definir o tipo, o tipo vai ser inferido pelo valor da variavel

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

### tipode de variaveis

bool

int -> sigment int (negativo e positovo)
uint -> unsigment int (positivos)

o tamanho depende da arquitetura, se a arquitetura para o qual tiver sendo compilado for 32 bits, vao ter 32 bits, se for 64 vao ter 64b its

int8, int16, int32, int64 (8, 16, 32, 64 bit respectivamente)

uint8, uint16, uint32, uint64 (8, 16, 32, 64 bit respectivamente)

uintptr -> u inter pointer

Utilizado quando esta escrevendo codigo unsafe,
por exemplo, integrações low level, integrações de codigo go com C

byte -> o mesmo que uint8

rune -> o mesmo que int32
-> utilizado para representar caracteres

float32 float64

complex64 complex128

ao converter um inteiro para string é convertido para o valor na tabela ASCII, e não a string numerica

ˋˋˋ
func main(){
var x = 65
// conversion from int to string yields a string of one rune, not a string of digits
s := string(x)
fmt.Println(s) // A ->caracterer na tabela ASCII
}
ˋˋˋ

para converter um valor para a string literal deve usar o pacote strconv

ˋˋˋ
func main() {
x := 65
s := strconv.FormatInt(int64(x), 10)
fmt.Println(s)
}
ˋˋˋ

convertemos o numero para int 64 bits, na base decimal e pegamos esse valor e convertemos para string

## contants

utiliza a palavra reservada 'const' e seu valor nao pode ser mudado, e é possivel compilar o projeto mesmo se a constante não tiver sendo utilizada, diferentemente das variaveis que causa erro de compilação se não forem utilizadas

ˋˋˋ
const x int = 10
ˋˋˋ

apenas os tipos de caractere pode ser uma cosntante: rune, byte, string, bool, e qualquer tipo numerico
Não pode usar short sintax (:=) para declarar uma constante.

As constantes sao unsigned type, ou seja, não tem tipo explicito, elas definem o tipo de acordo com o contexto

ˋˋˋ
func main() {
const x = 10
takeInt32(x)
takeInt64(x)
}

func takeInt32(x int32) {
fmt.Println(x)
}

func takeInt64(x int64) {
fmt.Println(x)
}
ˋˋˋ

Nesse caso funciona pq ambos do da familia do int

constant literal -> Magic number

ˋˋˋ
func main() {
takeInt32(10)
takeInt64(10)
}

func takeInt32(x int32) {
fmt.Println(x)
}

func takeInt64(x int64) {
fmt.Println(x)
}
ˋˋˋ

Não é possivel passar uma const float para funções do tipo int, mas inverso é permitido

ˋˋˋ
func main() {
const x = 3.14
#cannot use x (untyped float constant 3.14) as int64 value in argument to takeInt64
takeInt32(x)
takeInt64(x)
}

func takeInt32(x int32) {
fmt.Println(x)
}

func takeInt64(x int64) {
fmt.Println(x)
}
ˋˋˋ
ˋˋˋ
func main() {
const x = 3
takeInt32(x)
takeInt64(x)
}

func takeFloat32(x float32) {
fmt.Println(x)
}

func takeInt32(x int32) {
fmt.Println(x)
}

func takeInt64(x int64) {
fmt.Println(x)
}
ˋˋˋ

### arrays

o tamnho tem que ser constante

ˋˋˋ
arr1 := [3]int{}
arr2 := [3]int{1, 2, 3}
ˋˋˋ

é possivel definir qual posição vai receber determiado elemento

ˋˋˋ
arr3 := [10]int{5: 400, 7: 300}
ˋˋˋ

**NAO** podemos definir o tamanho de um array usando uma **variavel**. O tamanho do array deve ser sabido em tempo de compilação

ˋˋˋ
x:= 10
arry :=[x]int{}
ˋˋˋ

Porem **PODE-SE** usar **constantes** para definir o tamanho do array

ˋˋˋˋ
const x = 10
arry := [x]int{}
ˋˋˋ

### Loops

## for

para o for, não é necessario parenteses (causa erro de compilaćao)

```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

todos o statmetns sao opcionais

```
i :=  1
for ; i < 10;{
    fmt.Println(i)
    i++
}
```

Inclusive a condićão que, caso nao seja definida, vai executar o for infinitamente

```
i :=  1
for {
    fmt.Println(i)
    i++
}
```

## range loop

```
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
for range arry {
    // Codigo executado para cada elemento dentro do array
    fmt.Println("run")
}

```

o range retorna dois elementos, o index e o elemento, respectivamente.

```
for i, elem range arry {
    fmt.Println(i, elem)
}
```

porem, utilizando **blank identifier** (\_), podemos ignorar o primeiro elemento

```
for _, elem range arry {
    fmt.Println(elem)
}
```

## version 1.22+

range sob um valor inteiro

```
for i := range 10 {
    fmt.Println(i)
}
```

dessa forma o range vai fazer o loop executar 10x, porem so retorna um elemento que é o valor da execucao atual, 1 - 10 nesse exemplo

### if

assim como no for, não é necessario parenteses

```
x := 10
if x > 5 {
    fmt.Println(x)
}
```

e caso seja necessario é possivel declarar variaveis dentro do statemnet do if

```
if x := 10; x > 5 {
    fmt.Println(x)
}
```
porem so vai existir no escopo do if

### switch

tambem nao precisa de parenteses e o break é implicito, caso coloque receberar um alerta de duplicidade, mas o compilador nao quebra

```
switch x {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
default:
    fmt.Println("default")
}
```

e é possivel fazer com o que o case subsequente ao selecionado seja executado, utilizando a palavra reservada 'fallthrough'

```
x := 1
switch x {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")
default:
    fmt.Println("default")
}
``` 
no caso acima o 1 e o 2 serao impressos


é possivel, tambem, fazer com que o switch nao tenha uma variavel e sim uma expressao em cada case

```
switch {
case x > 5:
    fmt.Println("x é maior que 5")
case x < 5:
    fmt.Println("x é menor que 5")
default:
    fmt.Println("x é igual a 5")
}
```

é possivel declarar variaveis dentro do statemnet do switch

```
switch x := 10; x > 5 {
case true:
    fmt.Println("x é maior que 5")
case false:
    fmt.Println("x é menor que 5")
}
```

tambem é possivel ter mais de uma expressao em um mesmo case

```
switch x {
case 1, 2, 3:
    fmt.Println("1, 2 ou 3")
default:
    fmt.Println("default")
}
``` 
