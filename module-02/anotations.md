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
