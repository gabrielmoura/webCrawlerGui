# Perguntas Frequentes

## O que é Agente de Usuário?

O cabeçalho de requisição User-Agent é uma cadeia de caracteres que identifica a aplicação, sistema operacional,
fornecedor e/ou versão do agente de usuário que está fazendo a requisição. Mais informações podem ser
encontradas [aqui](https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Headers/User-Agent).

## Por que um link que incluí na busca é inicialmente aceito e depois removido?

Isso ocorre porque o sistema não permite a inclusão de links duplicados. Se você tentar adicionar um link que já existe
na fila, ele será automaticamente removido. Além disso, um link pode ser removido se for inválido ou se não atender aos
critérios de aceitação da busca.

## Consigo acessar o site más o site não é buscado? Erro 429

O erro 429 é um erro de status HTTP que indica que o servidor está sobrecarregado ou que o limite de solicitações foi
excedido.
Isso ocorre porque o site está bloqueando o acesso do WebCrawler.
Para resolver esse problema, você pode trocar de proxy.
Erro comum usando I2P para acessar a rede Tor.

## Uso com I2P

O sistema é compatível com o I2P, mas é necessário configurar o proxy http corretamente.

Habilite o proxy e Insirá `http://localhost:4444` em Proxy URL.

Para mais informações, consulte a [documentação oficial](https://geti2p.net/pt/docs/api).

## Uso com Tor

- Instale o Tor e privoxy

```bash
sudo apt install tor privoxy
```

- Configure o privoxy para usar o Tor

```bash
echo 'forward-socks5t / 127.0.0.1:9050 .' | sudo tee -a /etc/privoxy/config
```

- Inicie os serviços

```bash
sudo systemctl start tor
sudo systemctl start privoxy
```

- Configure o WebCrawler para usar o proxy

Habilite o proxy e Insirá `http://localhost:8118` em Proxy URL.

## Como usar o privoxy com Tor e I2p?

- Após a instalação configure

```conf
# /etc/privoxy/config
forward-socks5t   .onion  127.0.0.1:9050 .
forward          .i2p    127.0.0.1:4444
# As linhas abaixo especificam que qualquer requisição para endereços IP nas faixas de rede locais (192.168.*.*, 10.*.*.* e 127.*.*.*) devem ser tratadas localmente e não precisam ser encaminhadas para um proxy externo.
forward         192.168.*.*/     .
forward            10.*.*.*/     .
forward           127.*.*.*/     .
```

## Onde ficam os dados e arquivos de configuração?

- Os dados são armazenados em `~/.local/share/webCrawlerGui/`
- Os arquivos de configuração são armazenados em `~/.config/webCrawlerGui/config.yml`