Em GO, as funções que estão no mesmo package não precisam de importação. 
Eles são visíveis para todos do mesmo package.

1 - go mod init github.com/juniorleaoo/learn-go/gopportunities
2 - no import do GO, se você não definir nada, o nome da variável será o nome do pacote
3 - go mod tidy baixa as dependências utilizadas e remove as não utilizadas
4 - sempre que criarmos uma pasta no GO, estamos assumindo que essa pasta é um package
4.1 - se eu tenho uma pasta router, eu vou ter um arquivo router, com o mesmo nome da pasta
no javascript é o arquivo "index.js", que assume que é o arquivo principal daquela pasta
mas não é obrigatório, é uma convenção 
5 - Qualquer variável, função, etc que começa com a letra maiúscula, ela automáticamente é exportada
se for minuscula ela não é exportada (privada)
Essa regra são para packages diferentes, se eu estiver no mesmo package eu não preciso importar o arquivo
tudo que está no package é acessível para o package inteiro