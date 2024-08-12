Aqui está a versão atualizada da documentação com os comandos solicitados:

---

# Web Crawler para Extração de Dados

![Example Workflow](https://github.com/gildemberg-santos/webcrawlerurl_v2/actions/workflows/go.yml/badge.svg)

## Visão Geral

Este projeto fornece uma ferramenta automatizada para coletar informações específicas de páginas da web. Utilizando técnicas avançadas de mapeamento e extração de dados, o crawler analisa a estrutura das páginas e identifica padrões na apresentação dos dados, organizando-os em uma estrutura útil para análise ou processamento posterior.

## Pré-requisitos

- **Git**: Para clonar o repositório.
- **Docker**: Para construir e executar o container do projeto.
- **Make**: Para simplificar os comandos de execução.

## Instalação

Clone o repositório em sua máquina local:

```bash
git clone git@github.com:gildemberg-santos/webcrawlerurl_v2.git
cd webcrawlerurl_v2
```

## Configuração do Docker

### Construção da Imagem Docker

Construa a imagem Docker do projeto:

```bash
docker build -t webcrawlerurl_v2 .
```

### Executando o Container

Inicie o container em segundo plano:

```bash
docker run -d --name webcrawlerurl_v2 -p 8080:8080 webcrawlerurl_v2
```

Interagindo com o container:

```bash
docker exec -it webcrawlerurl_v2 /bin/bash
```

### Outros Comandos Docker

Para matar o container em execução:

```bash
docker container kill webcrawlerurl_v2
```

Para remover todos os containers parados e liberar espaço:

```bash
docker container prune -f
```

### Exibindo Logs do Container

Para visualizar os logs do container em execução:

```bash
docker logs webcrawlerurl_v2
```

## Uso do Makefile

### Iniciando o Container

Para iniciar o container utilizando o `make`, execute:

```bash
make start
```

### Parando o Container

Para parar o container:

```bash
make stop
```

### Executando o Serviço

Para rodar o serviço de web crawler:

```bash
make run
```

Modo de desenvolvimento:

```bash
make dev
```

### Executando os Testes

Para rodar todos os testes automatizados:

```bash
make test
```

### Realizando o Build do Projeto

Para construir o projeto:

```bash
make build
```

## Considerações Finais

Com esses comandos, você pode facilmente instalar, configurar e executar o web crawler para extrair dados de qualquer site. Sinta-se à vontade para explorar e customizar conforme as necessidades do seu projeto.
