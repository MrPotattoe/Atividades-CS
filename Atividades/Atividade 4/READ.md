Para a quarta questão eu fiz três abordagens diferentes, já que eu não pensei em como resolver ela usando a ideia de cache que foi exigida.

1 Abordagem
A primeira abordagem foi minha primeira ideia de como fazer. A priori eu não sabia como fazer algo como um "cache", então eu simplesmente coloquei uma variável "global" para servir como um cache. De resto, a função tem um algoritmo que verifica se um número é primo tentando dividir ele por todos os outros primos, o que, presumivelmente, demora bastante e tem uma complexidades difícil até mesmo de calcular, devido a algumas pequenas otimizações que eu fiz.
Prós:
- Cria uma slice só com números primos;
Contas:
- Complexidade alta.

2 Abordagem
A segunda abordagem é quase igual a primeira, a única diferença é que nela eu criei uma slice de booleanos na qual os índices primos são "true" e os números compostos são "false", que serve também como um "cache". O algoritmo para descobrir se um número é primo é igual o da primeira abordagem e, consequentemente, os problemas também. 
Prós:
- Nenhum;
Contras:
- Complexidade alta.

3 Abordagem
Depois de pesquisar um pouco, cheguei a conclusão de que, dado o nível e o intuito da questão, o melhor modo de fazer o que é desejado é usar o Crivo de Erastotenes. Então na terceira abordagem tem uma slice de primos e uma array com uma certa quantidade de números pré determinada. Daí, para você usa o Crivo em cada elemento da array maior ou igual do que 2 e cada primo determinado pelo crivo é colocado na slice de números primos.
Prós:
- Complexidade menor que a das outras duas abordagens;
Contras:
- É necessário determinar um número de "teto" para o crivo, o que é limitante tendo em vista que arrays tem tamanho limitado e não MUITO grande.

OBS: As duas primeiras abordagens são completamente "originais" e, provavelmente por isso, são as que mais demoram para ser executadas. O modo que o algoritmo que descobre os primos funciona nelas limita as possibilidades de cache.
