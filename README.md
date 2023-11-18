# MethodhJulia
MethodJulia, é um progama escrito em Go (golang), que utiliza o método de Regressão de Júlia, para calcular a Raiz Quadrada de números inteiros positivos.

## Regressão de Júlia
A Regressão de Júlia é uma fórmula inédita para o cálculo de raiz quadrada, desenvolvida pelo professor de Matemática da Unidade Cidade Nova do Colégio Santa Maria, Frederico Ferreira de Pinho Tavares, a partir de questionamentos de sua aluna, Júlia Helena Pimenta Ferreira.

A Regressão de Júlia é uma fórmula recursiva, ou seja, ela se baseia na repetição de um mesmo cálculo para chegar ao resultado final. A fórmula é a seguinte:

```xn = xn-1 - (xn-1)^2 / (2 * xn-2)```

onde:

- xn é a raiz quadrada de n
- xn-1 é a estimativa da raiz quadrada de n, no passo anterior
- xn-2 é a estimativa da raiz quadrada de n, no passo anterior ao anterior

## Instalação
### Linux
### Windows
## Como Compilar

```Requisitos: golang 1.21+```

```1. git clone https://github.com/erickmerys/MethodJulia.git```

```2. cd MethodJulia```

```3. go build .```
