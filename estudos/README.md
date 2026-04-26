# Estudos de Go — Documentação Completa

Repositório de exercícios e estudos da linguagem Go. Cada arquivo aborda um conceito específico, documentado abaixo com explicações, exemplos e diferenças em relação a outras linguagens.

---

## Índice

1. [main.go — Conversão de Tipos](#maingo--conversão-de-tipos)
2. [main1.go — Condicionais if/else](#main1go--condicionais-ifelse)
3. [main2.go — Loop for com continue](#main2go--loop-for-com-continue)
4. [main3.go — Loops aninhados](#main3go--loops-aninhados)
5. [main4.go — Slices: criação e iteração](#main4go--slices-criação-e-iteração)
6. [main5.go — Slices: make, append e filtragem](#main5go--slices-make-append-e-filtragem)
7. [main6.go — Slices: fatiamento](#main6go--slices-fatiamento)
8. [main7.go — Maps](#main7go--maps)
9. [exe.go — Exercício: soma de slice](#exego--exercício-soma-de-slice)
10. [exe2.go — Exercício: soma por condição](#exe2go--exercício-soma-por-condição)
11. [test-10.go — Exercício: último elemento](#test-10go--exercício-último-elemento)
12. [test11.go — Map vazio com make](#test11go--map-vazio-com-make)

---

## main.go — Conversão de Tipos

**Arquivo:** `main.go`

Em Go, **conversões de tipo nunca são implícitas**. Diferente de JavaScript ou Python, onde tipos são convertidos automaticamente, em Go você precisa fazer isso de forma explícita.

### Conversão numérica (type casting)

```go
a := 10            // Go infere como int
var numero2 float32
numero2 = float32(a)  // converte int → float32 explicitamente
fmt.Printf("%T\n", numero2) // imprime o tipo: "float32"
```

> `%T` no `Printf` imprime o **tipo** da variável, não o valor.

### Conversão string → número (`strconv.ParseInt`)

```go
b := "20"
i, _ := strconv.ParseInt(b, 10, 64)
// Argumentos: (string, base numérica, tamanho em bits)
// base 10 = decimal | bitSize 64 = int64
// Retorna (int64, error) — o "_" descarta o erro
```

| Função | De → Para | Observação |
|---|---|---|
| `strconv.ParseInt` | string → int64 | Especifica base e tamanho |
| `strconv.ParseFloat` | string → float64 | Especifica tamanho |
| `strconv.Atoi` | string → int | Atalho para base 10 |

### Conversão número → string (`strconv.Itoa`)

```go
c := 20
s := strconv.Itoa(c)   // "Integer to ASCII"
// Alternativa: strconv.FormatInt(int64(c), 10)
```

### Declaração explícita vs inferida

```go
var d int = 10   // tipo explícito
d := 10          // tipo inferido pelo compilador (equivalente)
```

---

## main1.go — Condicionais if/else

**Arquivo:** `main1.go`

Go usa `if/else if/else` igual a outras linguagens, com duas diferenças sintáticas importantes:
- **Sem parênteses** na condição
- **Chaves `{}` obrigatórias** mesmo com uma única linha

```go
salario := 1000.00
tipo := "CLT"

if salario < 1000.00 && tipo == "CLT" {
    salario -= (salario * 0.1)   // desconto 10%
} else if salario <= 2000.00 {
    salario -= (salario * 0.15)  // desconto 15%
} else {
    salario -= (salario * 0.2)   // desconto 20%
}
```

### Operadores lógicos

| Operador | Significado | Exemplo |
|---|---|---|
| `&&` | AND — ambas verdadeiras | `a > 0 && b > 0` |
| `\|\|` | OR — pelo menos uma verdadeira | `a > 0 \|\| b > 0` |
| `!` | NOT — inverte o booleano | `!ativo` |

### Operadores de atribuição composta

```go
salario -= (salario * 0.1)
// equivale a:
salario = salario - (salario * 0.1)
```

> **Atenção com a lógica:** neste exemplo `salario == 1000.00`, então `salario < 1000.00` é **falso**. O código cai no `else if` e aplica 15% de desconto. Resultado: `850.00`.

---

## main2.go — Loop for com continue

**Arquivo:** `main2.go`

Go tem **apenas uma estrutura de repetição: o `for`**. Ele substitui `while`, `do-while` e `for` clássico de outras linguagens.

### For estilo "while"

```go
i := 0
for i < tamanho {   // condição apenas, sem init nem pós-execução
    // ...
    i++
}
```

### `continue` — pula para a próxima iteração

```go
if string(texto[i]) == "r" {
    continue   // volta para o topo do loop sem executar o restante
}
fmt.Println(string(texto[i]))
i++
```

> **Bug de estudo:** quando `texto[i] == "r"`, o `continue` é executado antes do `i++`. O índice nunca avança e o loop trava **infinitamente** na letra "r". A correção seria mover o `i++` para antes do `continue`.

### Indexação de strings

```go
texto := "palavra"
len(texto)       // → 7 (número de bytes, não caracteres Unicode)
texto[i]         // retorna uint8 (byte)
string(texto[i]) // converte byte → string legível
```

---

## main3.go — Loops aninhados

**Arquivo:** `main3.go`

### For clássico (3 partes)

```go
for inicialização; condição; pós-execução {
    // ...
}
```

### Tabuada com loops aninhados

```go
for numbase := 1; numbase <= 10; numbase++ {
    for multiplicador := 1; multiplicador <= 10; multiplicador++ {
        fmt.Println(numbase, " x ", multiplicador, " = ", numbase*multiplicador)
    }
}
```

- O loop **interno** executa completamente a cada iteração do **externo**
- Total: 10 × 10 = **100 iterações**
- Complexidade: **O(N²)** — cresce quadraticamente com N

> Loops aninhados têm acesso às variáveis dos loops externos, mas não o inverso.

---

## main4.go — Slices: criação e iteração

**Arquivo:** `main4.go`

Slice é a estrutura de lista dinâmica mais usada em Go.

### Array vs Slice

| | Array | Slice |
|---|---|---|
| Tamanho | Fixo, parte do tipo | Dinâmico |
| Sintaxe | `[4]int{1,2,3,4}` | `[]int{1,2,3,4}` |
| Redimensionável | Não | Sim (com `append`) |
| Uso comum | Raramente | Frequente |

### Criação de slices

```go
// Forma 1: literal com valores iniciais
lista2 := []int{1, 2, 3, 4}

// Forma 2: make — tamanho definido, valores = zero do tipo
lista := make([]int, 1000)  // 1000 ints, todos zero

// Forma 3: strings
listaString := []string{"golang", "javascript", "php"}
```

### Iteração e append

```go
for i := 0; i < len(lista2); i++ {
    fmt.Printf("%v\n", lista2[i])
}

lista = append(lista, 6)   // adiciona elemento ao final
```

> `%v` imprime o **valor** padrão do tipo (funciona para qualquer tipo).

---

## main5.go — Slices: make, append e filtragem

**Arquivo:** `main5.go`

### nil slice vs slice vazio

```go
var i []int           // nil slice — valor zero, sem alocação
listaMenor := make([]int, 0)  // slice vazio, inicializado

// Ambos imprimem "[]" mas têm comportamentos diferentes:
// nil slice: i == nil → true
// make slice: listaMenor == nil → false
```

### Filtragem com append

```go
listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3}
listaMenor := make([]int, 0)

for i := 0; i < len(listaToda); i++ {
    if listaToda[i] < 5 {
        listaMenor = append(listaMenor, listaToda[i])
    }
}
// listaMenor → [2, 4, 1, 3]
```

> **append sempre retorna um novo slice** — pode realocar memória internamente. O retorno **deve** ser reatribuído: `listaMenor = append(listaMenor, ...)`.

### Variable shadowing

```go
var i []int           // i externo (slice)
for i := 0; i < len(listaToda); i++ {   // i interno (int) oculta o externo
    // aqui "i" é o int do for, não o slice declarado acima
}
```

---

## main6.go — Slices: fatiamento

**Arquivo:** `main6.go`

### Notação de fatiamento

```go
listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3}
//                  0   1  2  3  4  5  6  7

segundaLista  := listaToda[:3]              // [2, 10, 9]     → índices 0, 1, 2
terceiraLista := listaToda[3:]              // [4, 5, 8, 1, 3] → índice 3 até o fim
ultimoItem    := listaToda[len(listaToda)-1:] // [3]           → só o último
```

### Regra geral: `slice[inicio:fim]`

| Notação | Significado |
|---|---|
| `[2:5]` | Posições 2, 3, 4 (fim é exclusivo) |
| `[:3]` | Do início até posição 2 |
| `[3:]` | Da posição 3 até o fim |
| `[:]` | Cópia superficial do slice inteiro |
| `[len-1:]` | Apenas o último elemento (como slice) |

> **Cuidado:** sub-slices **compartilham o mesmo array subjacente**. Modificar um pode alterar o original. Use `copy()` se quiser uma cópia independente.

---

## main7.go — Maps

**Arquivo:** `main7.go`

Map é a estrutura chave-valor do Go (equivalente a `dict` do Python ou objeto do JavaScript).

### Criação de maps

```go
// Forma 1: literal com valores iniciais
m1 := map[string]int{"sp": 100, "rj": 200, "mg": 300}

// Forma 2: make — map vazio para preencher depois
m2 := make(map[string]int)
m2["sp"] = 100
m2["rj"] = 200
```

### Leitura — com e sem verificação

```go
// Leitura simples (retorna zero value se chave não existir)
valor, _ := m2["sp"]

// Leitura com verificação de existência (idiomático em Go)
if valor, ok := m2["sp"]; ok {
    fmt.Printf("Valor: %v\n", valor)
    // ok == true → chave existe
    // ok == false → chave não existe
}
```

> Sem o `ok`, não é possível distinguir "chave inexistente" de "valor igual ao zero value do tipo".

### Deletar e iterar

```go
delete(m2, "sp")   // remove a chave (seguro se não existir)

for chave, valor := range m2 {
    fmt.Println(chave, valor)
}
```

> **A ordem de iteração de maps é aleatória em Go** — propositalmente randomizada para evitar dependência de ordem específica.

### nil map vs map vazio

```go
var m map[string]int    // nil map — PANIC ao escrever
m := make(map[string]int) // map vazio — seguro para leitura e escrita
```

---

## exe.go — Exercício: soma de slice

**Arquivo:** `exe.go`

Exercício clássico: somar todos os elementos de um slice.

```go
lista := []int{1, 2, 3, 4}
soma := 0

for i := 0; i < len(lista); i++ {
    soma += lista[i]
}

fmt.Println("Soma:", soma) // → 10
```

### Alternativa idiomática com range

```go
for _, v := range lista {
    soma += v
}
// "_" descarta o índice quando não é necessário
// "range" é preferido quando não precisamos manipular o índice
```

---

## exe2.go — Exercício: soma por condição

**Arquivo:** `exe2.go`

Exercício com duas abordagens para separar e somar elementos de um slice.

### Abordagem 1 (comentada): fatiamento antes de somar

```go
slice1 := slice[:5]  // primeiros 5 elementos
slice2 := slice[5:]  // restante

// loop 1: soma slice1
// loop 2: soma slice2
```

### Abordagem 2 (ativa): condição dentro do loop

```go
slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
somaPrimeira := 0  // elementos <= 5
somaSegunda := 0   // elementos > 5

for i := 0; i < len(slice); i++ {
    if slice[i] <= 5 {
        somaPrimeira += slice[i]  // 2+3+5+4+1 = 15
    } else {
        somaSegunda += slice[i]   // 8+10+7+9 = 34
    }
}
```

### Comparação das abordagens

| | Abordagem 1 (fatiamento) | Abordagem 2 (condicional) |
|---|---|---|
| Loops | 2 | 1 |
| Alocação extra | Sim (sub-slices) | Não |
| Complexidade | O(n) | O(n) |
| Clareza | Alta (separa etapas) | Alta (compacta) |

---

## test-10.go — Exercício: último elemento

**Arquivo:** `test-10.go`

Demonstra como acessar o último elemento de um slice.

```go
listaToda := []int{12, 2, 3, 3, 3}

// Retorna um SLICE com o último elemento
segundaLista := listaToda[len(listaToda)-1:]
fmt.Print(segundaLista) // → [3]

// Para retornar o VALOR (int) diretamente:
// ultimoValor := listaToda[len(listaToda)-1]  → 3
```

> Go não suporta índices negativos como Python (`lista[-1]`). É necessário calcular `len-1` manualmente.

---

## test11.go — Map vazio com make

**Arquivo:** `test11.go`

Exemplo mínimo da criação de um map vazio.

```go
m := make(map[string]int)
fmt.Println(m) // → map[]
```

### Nil map vs make map

```go
var m1 map[string]int       // nil — m1 == nil → true
m2 := make(map[string]int)  // vazio — m2 == nil → false

m1["chave"] = 1  // PANIC: assignment to entry in nil map
m2["chave"] = 1  // OK
```

> **Sempre use `make` para criar maps que serão escritos.** A declaração com `var` cria um nil map que causa panic em runtime ao tentar inserir valores.

---

## Resumo dos Conceitos Estudados

| Conceito | Arquivo(s) | Pontos-chave |
|---|---|---|
| Conversão de tipos | `main.go` | Sempre explícita; `strconv` para string↔número |
| Condicionais | `main1.go` | Sem parênteses; chaves obrigatórias; `&&`, `\|\|`, `!` |
| Loop for (while) | `main2.go` | Único loop em Go; `continue` pula iteração |
| Loop for clássico | `main3.go` | `for init; cond; post { }` |
| Loops aninhados | `main3.go` | Loop interno executa N vezes por iteração do externo |
| Slices (criação) | `main4.go` | Literal `[]int{}` ou `make([]int, n)` |
| Slices (filtragem) | `main5.go` | `append` retorna novo slice; nil vs make |
| Slices (fatiamento) | `main6.go` | `[inicio:fim]` compartilha array base |
| Maps | `main7.go` | `make` obrigatório; verificar com ok; range é aleatório |
| Soma acumulada | `exe.go` | Padrão acumulador com `+=` |
| Soma condicional | `exe2.go` | Distribuir em acumuladores por condição |
| Último elemento | `test-10.go` | `slice[len-1:]` retorna slice; `slice[len-1]` retorna valor |
