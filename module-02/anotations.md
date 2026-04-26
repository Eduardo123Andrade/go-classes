### Memoria

## stack

LIFO -> Last In First Out

é nessa memoria que são armazenada as variaveis, argumentos de funções, e tudo que for necessario para a função executar.

Quanro retornamos da função, removemos a função da stack

## heap

Utilizada quando precisa de mais memoria para o programa.
Deve ser feito uma chamada ao SO dizendo quanto precisa alocar na memoria, quando termiar de executar a função deve fazer outra chamada para o SO liberar aquela alocação de memoria para outros programas utilizarem.

Caso não seja feito vai vai causar o **memory leaky** e essa memoria vai ficar bloqueada enquanto o SO estiver sendo executado

### Ponteiros

Ponteiros são endereços de memoria

ˋˋˋ
func main() {
x := 10
p := &x

    fmt.Println(p, *p) //0xc000094000 10

}
ˋˋˋ

para definir o valor (endereço da memoria) do ponteiro usa (&x) onde x é a variavel

para saber o valor desse endereço deve ser feito uma **dereferencia** utilizando (\*p), onde p é o ponteiro.

quando é passado um parametro para uma função é passado o valor da variavel
entao:

ˋˋˋ
func main(){
x := 10
take(x)
fmt.Println(x) // 10
}

func take(x int) {
x = 100
}
ˋˋˋˋ

o valor de x não é alterado, pois foi passado apenas o valor de x durante a execução da função, e não a propria variavel. Nesse caso estamos fazendo uma copia do valor da variavel x, para o parametro da função

Para que seja possivel alterar o valor de X, deve usar ponteiros

ˋˋˋ
func main(){
x := 10
take(&x)
fmt.Println(x) // 100
}

func take(x *int) {
*x = 100
}
ˋˋˋ

qunado usar: - Quando queremos mudar o valor alocado na memoria - quando nao queremos passar o valor da variavel e sim a referencia como parametro - evitar copiar dados pesados (usamos a referencia da memoria para acessar esses dados) - e performance em alguns casos

\*int -> Ponteiro de inteiro (tipo de dado)

\*X -> derreferencia (pega o valor alocado na memoria daquele ponteiro)

### Struct

São coleções de dados definido da seguinte maneira:

ˋˋˋ
type User string {
Name string
ID uint64  
}
ˋˋˋ

podemos instanciar uma struct das seguintes maneiras

ˋˋˋ
user := User{}
ˋˋˋ
Caso não definamos os valores de cada atributo, vai ser setado o valor padrao ('', 0) respectivamente no exemplo acima

ˋˋˋ
user := User{"Teste", 10}
ˋˋˋ
onde passando apenas os valores **na ordem** definida na struct, logo, nao precisamos explicitar a qual tributo ele pertence

ˋˋˋ
user := User{Id: 10, Name: "Teste"}
ˋˋˋ

onde explicitamos a quais atributos vao ser atribuido os valores, porem **não** precisamos seguir uma ordem exata na atribuição

ˋˋˋ
user := User{Name: "Teste"}
ˋˋˋ

e por fim definir parcialmente os atributos com valores, o que nao foram definidos vao ter o valor padrao

OBS.:

Não é possivel misturar declaração explicita com declaração ralativa

ˋˋˋ
user := User{Name: "Teste", 0}
// mixture of field:value and value elements in
//struct literalcompilerMixedStructLit
ˋˋˋ

Para acessar algum dos valores basta apenas sar o .

ˋˋˋ
func main() {
user := User{Name: "Teste", Id: 0}
fmt.Println(user.Name)
}
ˋˋˋ

## Declaração de metodos para tipos

É possivel declarar um metodo para qualquer tipo, desde que, estejam no mesmo pacote.

a declaração do metodo é feita iqual a uma função, porem é passado o tipo que vai receber esse metodo

ˋˋˋ
func (User) print () {
fmt.Println("foo")
}
ˋˋˋ

caso deseje interagir com os valores dos tipo que o metodo foi atribuido devesse usar uma variavel, a recomendação da comunidade é que o nome da variavel seja as iniciais do tipo no nosso caso (u)

ˋˋˋ

func (u User) print() {
fmt.Println(u.Name)
}
ˋˋˋ

OBS.: Um metodo **não** pode ter mias de um recebedor

ˋˋˋ
// method has multiple receivers
func (u User, u2 User2) print() {
fmt.Println(u.Name)
}
ˋˋˋ

## Atualizando valores de um tipo atrave de um metodo declarado

ˋˋˋ
type User struct {
Name string
Id uint64
}

func (u User) UpdateName(name string) {
u.Name = name
}

func main() {
user := User{Name: "Teste", Id: 0}
user.UpdateName("Pedro")
fmt.Println(user.Name) // "Teste"
}
ˋˋˋ

O valor acima não mudou, pois a declaração ˋˋˋfunc (u User) UpdateName(name string) ˋˋˋ faz uma **copia** da variavel ˋˋˋuserˋˋˋ.

Para que possamos atualizar o valor da propria variavel que chama o metodo precisamos usar **pointer indirection**

ˋˋˋˋ
func (u \*User) UpdateName(name string)
ˋˋˋ

onde dizemos que vamos apontar para o local da memoria onde aquela variavel esta.

ˋˋˋ
func main() {
user := User{Name: "Teste", Id: 0}
user.UpdateName("Pedro")
fmt.Println(user.Name) // "Pedro"
}
ˋˋˋ

Desta maneira a função consegue alterar a variavel corretamente

## Embedded types

É a maneira atribuir os atributos e funcoes de uma struct em outra

ˋˋˋ
// foo/foo.go

type Foo struct {}

func (f Foo) Bar {}
ˋˋˋ

ˋˋˋ
// outro pacote
type User string {
foo.Foo // Embedded type
Name string
ID uint64  
}
ˋˋˋ

Nesse caso, o tipo **User** tem todos os atributos e metodos de **Foo**.

ˋˋˋ
user := User{}
user.Bar()
ˋˋˋ

##struct tags

São tags que indicam para outros pacote como "lidar" com os atributos

ˋˋˋ
type User struct {
Name string `json:"name"`
Id uint64 `json:"id"`
}
ˋˋˋ

sem as tags o resultado é

ˋˋˋ
res, err := json.Marshal(user)
if(err != nil){
panic(err)
}
fmt.Println(string(res))
//{"Name":"Teste","Id":0}
ˋˋˋ
com os atributos captalizado.

resultado apos as tags: ˋˋˋ{"name":"Teste","id":0}ˋˋˋ

# Interface

Contrato que diz como uma variavel deve ser atribuida

ˋˋˋˋ
type Animal interface {}
ˋˋˋ

a partir da versão 1.18, o go ganhou um alias para interface fazia chamado _any_

ˋˋˋ
func foo(x interface{}){}
func bar(x any){}
ˋˋˋ

A interfaces podem ser implementadas indiretamente, caso uma variavel tenha todos os metodos e atributos que a _itnerface_ exige, a variavel implementa indiretamente essa interface.

ˋˋˋ
type Animal interface {
Sound() string
}

type Dog struct{}

func (Dog) Sound() string {
return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
fmt.Println(a.Sound())
}

func main() {
dog := Dog{}

    whatDoesThisAnimalSay(dog) //Au! Au!

}
ˋˋˋ

mesmo sem uma keyword _implements_ o go implementa automaticamente a interface

ˋˋˋ
type Animal interface {
Sound() string
}

type Dog struct{}

func (Dog) Bark() string {
return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
fmt.Println(a.Sound())
}

func main() {
dog := Dog{}

    whatDoesThisAnimalSay(dog)

}
ˋˋˋ

porem nesse novo bloco, o metodo atribuido é o **Bark** e nao mais **Sound**, então o compilador alerta um erro: _cannot use dog (variable of struct type Dog) as Animal value in argument to whatDoesThisAnimalSay: Dog does not implement Animal (missing method Sound)_

No codigo abaixo a variavel _a_ é iniciada como nula e depois passa a servir de referencia a um ponteiro que aponta para nil, "mantendo" a variavel _a_ como "nula"

OBS.: Só é possivel pois estamos lidando com ponteiros

ˋˋˋ
type Animal interface {
Sound() string
}

type Dog struct{
Name string
}

func (d \*Dog) Sound() string {
fmt.Println(d.Name)
return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
fmt.Println(a.Sound())
}

func main() {
var a Animal // nil
var dog \*Dog // nil
a = dog // not nil
whatDoesThisAnimalSay(a) // Au! Au!
}
ˋˋˋ

Mas em **Go** podemos executar funções de _reciveirs_ nulos, desde que não utilizemos internamente algum atributo desse receiver, pois dessa maneira dipararia um _panic_: "panic: runtime error: invalid memory address or nil pointer dereference"

Isso só é possivel pois estamos lidando com ponteiros, caso tentarmos executar a funcao diretamente da interface nula teremos um panic:

ˋˋˋ
func main(){
var a Animal
//panic: runtime error: invalid memory address or nil pointer dereference
fmt.Println(a.Sound())
}
ˋˋˋ

Interfaces vazias são implementadas por qualquer variavel, pois essas interfaces não tem atributos

# Type assertion

em Go podemos verificar o tipo da variavel usando o seguinte codigo ˋˋstr, ok := a.(string)ˋˋ

o ok serve para garantir que o tipo que estamos tentando fazer o _assertion_ é realmente uma string, se nao for o ok vai ser _false_, caso nao seja feita a validação do ok, vai causar panic: _interface conversion: interface {} is int, not string_

## Switch case de tipos

ˋˋˋ
type Animal interface {
Sound() string
}

type Dog struct{
Name string
}

type Cat struct {}

func (d \*Dog) Sound() string {
return "Au! Au!"
}

func (c \*Cat) Sound() string {
return "Miau!"
}

func takeAnimal(a Animal) (x string) {
x = "<nil>"
switch t := a.(type) {
case *Dog:
x = t.Sound()
case *Cat:
x = t.Sound()
}

    return

}
ˋˋˋ

usando a keyword _type_ podemos pegar dinamicamente o tipo da variavel _a_ para utilizarmos nos cases do swtich

# Erros

Erros são interfaces que existe por padrão na linguagem, essa funcõe guardam valores

```
// Interface padrão de erro em Go
type error interface {
    Error() string
}
```

Quando uma funcão erra não é disparado um erro, nem um _panic_ , ou seja não existe `throw new Error()` como em outras linguagens
ou `panic()`

_panic só pode ser usado quando o programa não pode mais continuar e idealmente não deve ser usado, deve retornar um erro para que, quem chamou a funcão
decisa o que deve ser feito_

uma funćao em go pode retonar mais um valor, como já foi dito anteriormente, então os erros são _retornados_ na funcão e não disparam uma excecao

```
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não pode dividir por 0")
	}

	return a / b, nil
}
```

Pela convencão, caso seja necessario retornar um erro, devemos retornar a _valor default_ do valor, nesse caso `return 0, erro`

```
func main(){
	a := 10
	b := 0

	res, err := dividir(a, b)

	if err != nil {
		 fmt.Println(err)
		 return
	}

	fmt.Println(res)

}
```

e a funćão que chamou a funcão dividir, deve tratar com erro retornado da maneria que desejar

## Erros customizados

```
type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {return s.msg}

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"Não exite raiz quadrada de numero negativo"}
	}

	return math.Sqrt(x), nil
}
```

em outras linguagens teriamos `class CustomError extends Error {}`, porem em Go, eu crio uma estrutura para o meu erro customizado `type CustomError struct {}`

e atribuo a funcão Error a essa struct `func (s SqrtError) Error() string {return "Erro customizado"}`, como um erro deve implemetnar a inteface erro do Go que possui o metodo `Error() string`, então apos a a atribuićão do erro a minha struct CustomError, essa struct pode ser usada em "validacões" de erros em Go

## Validacóes

### errors.is

Esse tipo de validacão verifica se o erro passado é um erro originado do `errors.New()` (_erros padrões de Go_)

```
type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

var ErrNotFound = errors.New("not found")

func main() {

	err := foo()

	if err != nil && errors.Is(err, ErrNotFound) {
		fmt.Println("Error not found")
		return
	}

	fmt.Println("Fora do if")

}

func foo() error { return ErrNotFound }
```

Nesse caso eu estou comparando o erro em si e não o tipo, ou seja _ErrNotFound_ é um tipo não uma interface, entao eu verifico no meu if, se o erro retornado pela funcão é um erro igual ao ErrNotFound "parecido com 1 == 1".

### errors.as

```
type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

var ErrNotFound = errors.New("not found")

func main() {

	err := foo()
	var sqrtError SqrtError

	if err != nil && errors.As(err, &sqrtError) {
		fmt.Println(sqrtError.msg)
		return
	}

	fmt.Println("Fora do if")

}

func foo() error { return SqrtError{"teste"} }
```

a funcão As precisa receber o erro e um ponteiro, para que, caso encontre o erro na error tree (em outras linguagens: _Error stack_) , vai mutar a variavel para o erro encontrado.
ou seja, caso enconter algum erro _SqrtError_ na arvore, vai apontar esse erro para a variavel _sqrtError_, para que possamos acessar seus atributos e funcões se necessario `srqtError.msg`

OBS.: Mesmo que minha funcao seja um _pointer receiver_ `func (s *SqrtError) Error() string { return s.msg }`, ainda sim, precisamos passar um um ponteiro para o `errors.As(err, &somePointerError)`

### Quando usar CustomError struct ou erros.New

Sempre que precisar armazenar valores no erro: _input, messagem, codigo, descricão_, usamos uma _struct_

caso precisamos apenas de um erro simples, usamos _errors.New()_

### Error wraping

Erros que contem varios erros

```
func main() {
	err := foo()

	if err != nil && errors.Is(err, CustomError) { // false
		fmt.Println("deu erro:", err)
		return
	}

}

var CustomError = errors.New("error")

func foo() error {
	err := bar()
	if err != nil {
		return errors.New("deu erro em foo " + err.Error())
	}
	return nil
}

func bar() error { return CustomError }
```

Nesse acaso err não é um CustomErro, por mais que seja retornado um CustomError pela funcao _bar_, o CustomErro esta **incluso** no `errors.New("deu erro em foo " + err.Error())` na retornado pela funcao _foo_.

```
func foo() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("deu erro em foo: %w", err)
	}
	return nil
}
```

para contornar essa situaćão o pacote _fmt_ tem a funcao _Errorf_ que possuir um verbo **%w** de _wrapper_ para dizer que o erro _err_ retornado por bar, está dentro do erro retornando pela funcao _foo_, dessa maneria o `errors.Is()` consegue fazer o match com CustomError

Por isso a comunidade enconraja fortemente pela comunidade, fazer wrapper de erros para que, quem chamar essa funcao possa usar`errors.Is()` e identificar erros com precisão.

```
unc main() {
	err := foo()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))

}

var (
	ErrQualquer  = errors.New("Erro")
	ErrQualquer2 = errors.New("Erro 2")
)

func a() error { return ErrQualquer }
func b() error { return ErrQualquer2 }

func foo() error {
	var errorRetults error

	if err := a(); err != nil {
		errorRetults = errors.Join(errorRetults, err)
	}

	if err := b(); err != nil {
		errorRetults = errors.Join(errorRetults, err)

	}

	return errorRetults

}
```

Outra maneira de fazer wrapper é usando o `errors.Join(groupper, error)`

o resultado impresso na funcão main é:

fmt.Println(err):
Erro
Erro 2

onde o `fmt.Println(err)` concatena os erros com `\n`

e consguimos fazer o match do erro com o `errors.Is()`

`fmt.Println(errors.Is(err, ErrQualquer)) //true`
`fmt.Println(errors.Is(err, ErrQualquer2)) //true`
