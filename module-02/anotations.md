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
