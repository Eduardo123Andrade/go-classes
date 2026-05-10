// Desafio Expert: O Analista de Logs
// Crie um programa que leia um arquivo de log gigante (que não cabe na RAM). O programa deve:

// Usar bufio.Reader para ler o arquivo em pedaços (chunks).

// Filtrar linhas que contenham a palavra "ERROR".

// Salvar essas linhas em um arquivo errors.log usando bufio.Writer.

// Ao final, imprimir quantos bytes totais foram processados e quanto tempo levou.