Mapas são uma estrutra **chave-valor**

o valor inicial de um mapa é _nulo_ (**nil**)

ˋˋˋ
func main(){
var m map[string]string
fmt.Println(m == nill) // true
}
ˋˋˋ

ˋˋˋ
func main() {
var m map[string]string
// nil dereference in map update
m["Pedro"] = "Pessoa" // panic
fmt.Println(m)
}
ˋˋˋ

um mapa precisa ser **inicializado** antes de adicionar um valor, no codigo acima o mapa esta _nulo_

## Inicialziando mapas

ˋˋˋ
func main() {
m := make(map[string]string, x)
fmt.Println(m == nil) //false
}
ˋˋˋ

para criar mapas usando a função _make_, deve **obrigatoriamente** passar o tipo _map[string]string_ e, opcionalmente, o tamanho do mapa (_x_)

ˋˋˋ
func main() {
m := map[string]string{}
fmt.Println(m == nil) //false
}
ˋˋˋ
inicializa um mapa com valor vazio (**diferente de nulo**)

ˋˋˋ
func main() {
m := map[string]string{
"Pedro": "Pessoa",
"Carlos": "Cachorro",
}
fmt.Println(m == nil) //false
}
ˋˋˋ
inicializa um mapa com valores definidos

## mapa de slices

ˋˋˋ
func main() {
m := map[string][]int{
//redundant type from array, slice, or map composite literal
"Pedro": []int{1, 2, 3},
}
fmt.Println(m)
}
ˋˋˋ
Por ja ser definido previamente que o meu mapa é uma chave _string_ para um valor _slice de int_ não é necessario explicitar no valor que é um slice de int

ˋˋˋ
func main() {
m := map[string][]int{
"Pedro": {1, 2, 3},
}
fmt.Println(m)
}
ˋˋˋ

# set e get

ˋˋˋ
func main() {
m := make(map[string]string)
// seta a chave Pedro e o valor Pessoa
m["Pedro"] = "Pessoa"
// busca o valor da chave "Pedro"
valor := m["Pedro"]
fmt.Println(valor)
}
ˋˋˋ

ˋˋˋ
func main() {
m := make(map[string]string)
m["Pedro"] = "Pessoa"
valor, ok := m["Pedro"]
fmt.Println(valor, ok) // Pedro true

valor2, ok2 := m["foo"]
fmt.Println(valor2, ok2) // "" false
}
ˋˋˋ
Nosso map foi definido como um **valor string**, então, caso tentarmos buscar um chave que não existe no map (_foo_) vai retornar uma string vazia, então por ao buscarmos um valor em um map **valor, sucesso = m[chave]** é retornado um segundo valor _booleano_ para dizer se a chave existe ou não no map

# delete

ˋˋˋ
func main() {
m := make(map[string]string)
m["Pedro"] = "Pessoa"
valor, ok := m["Pedro"]
fmt.Println(valor, ok) // Pedro true

    delete(m, "Pedro")
    valor, ok = m["Pedro"]
    fmt.Println(valor, ok) // "" false

}
ˋˋˋ
Para apagar uma chave de um mapa usamos a função **delete** passando o mapa e a chave a ser deletada

ˋˋˋ
func main() {
m := make(map[string]string)
m["Pedro"] = "Pessoa"
m["foo"] = "bar"
fmt.Println(m) // map[Pedro:Pessoa foo:bar]

    clear(m)
    fmt.Println(m) // map[]

}
ˋˋˋ
Limpa todas as chaves do mapa

# loop

ˋˋˋ
func main() {
m := make(map[string]string)
m["Pedro"] = "Pessoa"
m["foo"] = "bar"

    for k, v := range m {
    	fmt.Println(k, v)
    }

}
ˋˋˋ
a funcao range, quando usada com mapas, retorna os valores **key** _k_ e **value** _v_
