Chat em Go
==========

Este é um projeto simples de chat em Go que inclui um servidor, um cliente e um bot. O servidor permite que vários clientes se conectem e conversem em uma sala de chat.

Executando o Servidor
---------------------

Para executar o servidor, siga os passos abaixo:

1.  Navegue até a pasta `server`:
    
    `cd server`
    
2.  Execute o servidor com o comando:
    
    `go run server.go`
    

O servidor agora está em execução e aguardando conexões de clientes.

Executando o Cliente
--------------------

Você pode executar um ou mais clientes para se conectar ao servidor. Siga os passos abaixo para executar um cliente:

1.  Navegue até a pasta `client`:
    
    `cd client`
    
2.  Execute o cliente com o comando:
    
    `go run client.go`
    

Repita os passos acima para executar múltiplos clientes e permitir que eles conversem entre si no servidor.

Executando o Bot
----------------

O bot é um cliente especial que responde a comandos específicos. Siga os passos abaixo para executar o bot:

1.  Navegue até a pasta `bot`:
    
    `cd bot`
    
2.  Execute o bot com o comando:
    
    `go run bot.go`
    

Agora, o bot "BobinhoBot" está em execução e pode responder a comandos como `/oi`, `/piada`, `/hora` e `/inverte`.

Ordem Recomendada para a Execução
---------------------------------

Recomenda-se seguir a seguinte ordem para executar os componentes:

1.  Execute o servidor primeiro com `go run server.go`.
2.  Em seguida, execute pelo menos um cliente com `go run client.go`.
3.  Se desejar, execute o bot com `go run bot.go`.

Certifique-se de que o servidor esteja em execução antes de iniciar os clientes ou o bot. Os clientes podem se conectar ao servidor e interagir entre si, e o bot pode participar da conversa e responder a comandos específicos.

Funções do Usuário e do Bot
---------------------------

*   **Usuário (Cliente)**: Os usuários são clientes que se conectam ao servidor para enviar mensagens, ler mensagens de outros usuários e interagir no chat. Eles podem conversar e usar comandos como `/oi`, `/hora`, `/piada` e `/inverte`.
*   **Bot ("BobinhoBot")**: O bot é um cliente especial que responde a comandos específicos. Ele pode interagir com outros usuários e responder a comandos como `/oi` (saudação), `/hora` (informar a hora atual), `/piada` (contar uma piada aleatória) e `/inverte` (inverter texto).

Divirta-se explorando este projeto de chat em Go!
