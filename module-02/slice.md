### Slice

Array dinamico, vai crescendo a medida que for adicionando novos itens.
slices usam ponterios para referenciar os valores dentro do array

como criar um slice apartir de um array,

ˋˋˋ
func main(){
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]
fmt.Println(slice) // [2, 3, 4]
}

define o indice de inicio (indece 1 no exemplo), e o indice de finalização mais 1, pois o intervalo é aberto a esquerda, ou seja, o valor definido no é incluido. Nesse caso, queremos do indicie 1 ou 3, mas como o 3 nao seria aplicado, adicionamos 1 para incluir o 3.

## instanciando um slice

ˋˋˋ
slice := []int{1, 2, 3}
fmt.Println(slice)
ˋˋˋ

para definir um slice usasse a mesma denotação de um array, porem, sem definir o tamanho inicial. o slice vai crescendo a medida que for inserindo novos valores

### slice literals

ˋˋˋ
slice := []int{} // slice vazio
fmt.Println(slice == nil) // false
ˋˋˋ

ˋˋˋ
var slice []int{} // slince nulo
fmt.Println(slice == nil) // true
ˋˋˋ

# funções builtin

ˋˋˋ
len([]) // tamanho do slice ou array
cap([]) // tamanho do slice ou array
append([], valor) // adicionar novo valor ao slice
make(x, y, z) // criar um slice do tipo x, com o tamanho y e a capacidade z
ˋˋˋ

tamanho -> quantidade de espaços com valores alocados
capacidade -> quantidade de elementos que aquele array ou slice pode armazenar

append -> sempre que o slice atinge a **capacidade** maxima, é criado um novo slice com o dobro do **tamanho**

## contruindo slice apartir de um array

ˋˋˋ
arr := [5]int{1, 2, 3, 4, 5}
slice1 := arr[x:y]
slice2 := arr[x:y:z]
ˋˋˋ

x -> indicie inicial, valor default = 0
y -> indice final, valor default = length do array
z -> capacidade maxima do slice, deve ser igual a quantidade de elementos do que esta sendo definido no slice, z = y

ˋˋˋ
arr := [5]int{1, 2, 3, 4, 5}
slice1 := arr[:]
ˋˋˋˋ
nesse caso esta sendo feito um slice do indicie **0** ate o **len(arr)** por padrão

ˋˋˋ
arr := [5]int{1, 2, 3, 4, 5}
slice2 := arr[:2:2]
ˋˋˋ
nesse caso esta sendo feito um slice do indice **0** ate o **indice 1** com **capacidade** de **2** elementos

# bound check

ˋˋˋ
func main() {
slice := []int{1, 2, 3}
foo(slice)
}

func foo(slice []int) {
\_ = slice[4]
fmt.Println(slice[0])
fmt.Println(slice[1])
fmt.Println(slice[2])
fmt.Println(slice[3])
}
ˋˋˋ

para cada chamada de fmt.Println(slice[i]) o compilador faz uma checagem para saber se o index passado é menor que a **capacidade** do slice/array, no exemplo acima seriam feitas 4 checagens

porem ao adicionar **\_ = slice[3]** o compilador verifica se o index **3** esta dentro do capacidade do array, se estiver segue o codigo normalmente, se não é lançado o **panic**, porem com isso o compilador so precisa checar apenas uma vez ao invez de 4 vezes.

### valor x referencia

# array

ˋˋˋ
func main(){
arr := [4]int{1, 2, 3, 4}
foo(arr)
fmt.Println(arr) // [1, 2, 3, 4]
}

func foo(arr [4]int) {
arr[0] = 123
}
ˋˋˋ

Quando passamos um array como parametro de uma função, passamos o valor dele, ou seja, **copiamos todo o conteudo e criamos um novo valor**. Nesse caso, mesmo que dentro da função _foo_ mudarmos o valor do parametro _arr_ não mudamos o array original

# slice

ˋˋˋ
func main(){
slice := []int{1, 2, 3, 4}
foo(slice)
fmt.Println(slice) // [123, 2, 3, 4]
}

func foo(slice [4]int) {
slice[0] = 123
}
ˋˋˋ

Quando passamos um slice como parametro de uma função, nos **copiamos todo o conteudo e criamos um novo slice** assim como no array, porem, o slice é um **ponteiro** para um array, ou seja, quando copiamos para a função nos **criamos um novo slice que aponta para as mesmas posições de memoria** que o slice original

# ponteiro de array

ˋˋˋ
func main(){
arr := [4]int{1, 2, 3, 4}
foo(&arr)
fmt.Println(arr) // [123, 2, 3, 4]
}

func foo(arr \*[4]int) {
arr[0] = 123
}
ˋˋˋ
Nesse caso temos um comportamento _parecido_ com o do slice já que **estamos passando a referencia da memoria** daquele array como parametro da função, ao invez do array em si
