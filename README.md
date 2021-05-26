# Carcereiro

CLI para liberação de acesso a banco de dados.

Hoje tenho necessidade de fazer liberações constantes de acesso a base de dados, isso aqui é o primeiro passo para uma automação mais completa e incremental do fluxo.

Primeiro projeto após [curso da Ellen Korbers](https://www.youtube.com/channel/UCxD5EE0H7qOhRr0tIVsOZPQ) de Golang. Fazer em shellscript seriam bem tranquilo, mas vamos aprender melhor uma nova linguagem.

## Estrutura

carcereiro<br>
&emsp;&emsp;configure<br>
&emsp;&emsp;liberar<br>
&emsp;&emsp;&emsp;&emsp;select<br>
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;database.tabela usuario<br>


Versões futuras
>&emsp;&emsp;&emsp;&emsp;update<br>
>&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;database.tabela usuario<br>
>&emsp;&emsp;&emsp;&emsp;insert<br>
>&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;database.tabela usuario<br>
>&emsp;&emsp;&emsp;&emsp;delete<br>
>&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;database.tabela usuario<br>
>&emsp;&emsp;prender<br>
>&emsp;&emsp;&emsp;&emsp;database.tabela usuario<br>