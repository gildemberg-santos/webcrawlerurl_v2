# webcrawlerurl_v2

[![Go](https://github.com/gildemberg-santos/webcrawlerurl_v2/actions/workflows/go.yml/badge.svg)](https://github.com/gildemberg-santos/webcrawlerurl_v2/actions/workflows/go.yml)

Serviço HTTP em Go que mapeia sites, lê sitemaps e feeds do Google Shopping e extrai o texto das páginas em um formato pronto para alimentar bases de conhecimento de IA (RAG, chatbots, busca semântica).

## O que ele faz

- **Mapeamento de URLs** a partir de uma página inicial, com limite de páginas, filtro por padrão de URL e lista de URLs descartadas.
- **Leitura de sitemaps**, incluindo sitemaps aninhados e versões em vários idiomas.
- **Leitura de feeds do Google Shopping** (RSS/Atom), com imagens, preço e parcelamento.
- **Extração de texto e meta tags**, com renderização via Chrome headless ([chromedp](https://github.com/chromedp/chromedp)) para páginas que dependem de JavaScript.
- **Normalização de URLs**: resolve links relativos e padroniza barras para evitar visitas duplicadas.

## Arquitetura

```
cmd/
  serve/     servidor HTTP (ponto de entrada em produção)
  cli/       execução local para testes manuais
  drive/     utilitário de download/descompactação do Google Drive
pkg/         casos de uso expostos pelas rotas (mapeamento, leitura, e-commerce...)
util/        blocos reutilizáveis: carregamento de página, extração, sitemap,
             Google Shopping, normalização e filtro de URLs, chunking
```

## Rodando

Com Docker (a imagem final usa o Chrome headless do `browserless/chrome`):

```sh
docker build -t webcrawlerurl_v2 .
docker run -d --name webcrawlerurl_v2 -p 8080:8080 webcrawlerurl_v2
```

Sem Docker (requer Go 1.24 e um Chrome/Chromium instalado):

```sh
make dev     # sobe o servidor na porta 4747 (ou na definida em PORT)
make test    # roda a suíte de testes
```

## API

| Método | Rota | Entrada | Descrição |
|---|---|---|---|
| GET | `/mappingurl` | `?url=&limit=` | Lista as URLs encontradas a partir da página inicial |
| GET | `/readtext` | `?url=` | Extrai texto e meta tags de uma página |
| GET | `/smart_call` | `?url=` | Título, parágrafo principal e meta description da página |
| POST | `/crawler_leadster_ai` | JSON | Mapeia e extrai o texto de várias páginas de um site (sitemap ou navegação) |
| POST | `/e-commerce-sitemap` | JSON | Lista URLs de produtos a partir do sitemap de uma loja |
| POST | `/e-commerce-google-shopping` | JSON | Lê os produtos de um feed do Google Shopping |
| POST | `/e-commerce` | JSON | Extrai o texto de uma lista de URLs de produtos |

### Exemplos

Mapear até 50 URLs de um site:

```sh
curl "http://localhost:8080/mappingurl?url=https://example.com&limit=50"
```

```json
{
  "success": true,
  "failure": false,
  "message": "",
  "data": ["https://example.com/", "https://example.com/sobre", "https://example.com/contato"],
  "ts": 1.42
}
```

Extrair o texto de uma página:

```sh
curl "http://localhost:8080/readtext?url=https://example.com/sobre"
```

Ler um site inteiro pelo sitemap, respeitando limite de páginas e tempo:

```sh
curl -X POST http://localhost:8080/crawler_leadster_ai \
  -H "Content-Type: application/json" \
  -d '{
        "url": "https://example.com",
        "url_pattern": "https://example.com/blog",
        "max_url_limit": 100,
        "max_timeout": 120,
        "is_load_fast": true,
        "is_sitemap": true,
        "is_complete": false,
        "discarded_urls": ["https://example.com/login"]
      }'
```

A resposta traz o total de caracteres extraídos e, para cada página, a URL, o texto e as meta tags.

## Testes

```sh
go test ./...
```

Os testes usam [httpmock](https://github.com/jarcoal/httpmock) e [testify](https://github.com/stretchr/testify) e rodam no GitHub Actions a cada push e pull request.

## Licença

[MIT](LICENSE)
