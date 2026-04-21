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
