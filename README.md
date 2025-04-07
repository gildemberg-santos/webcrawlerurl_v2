# Web Crawler para Extração de Dados

![Example Workflow](https://github.com/gildemberg-santos/webcrawlerurl_v2/actions/workflows/go.yml/badge.svg)

## **Visão Geral**  
Este projeto fornece uma **solução automatizada** para coletar e organizar informações específicas de páginas da web. Utilizando **técnicas avançadas de mapeamento e extração de dados**, o crawler analisa a estrutura das páginas, identifica padrões na apresentação dos dados e os organiza em uma estrutura útil para análise ou processamento posterior.  

### **Principais Benefícios**  
✅ **Automação eficiente** – Reduz esforço manual na coleta de informações.  
✅ **Precisão na extração** – Identificação e estruturação automática dos dados.  
✅ **Flexibilidade e personalização** – Configuração adaptável conforme as necessidades do usuário.  

---

## **Pré-requisitos**  
Antes de iniciar a instalação, certifique-se de ter os seguintes componentes:  

- **Git** – Necessário para clonar o repositório.  
- **Docker** – Utilizado para construir e executar o container do projeto.  
- **Make** – Facilita a execução de comandos automatizados. 

---
## Instalação e Configuração  

### 1. Clonando o Repositório  
Para obter o código-fonte, execute os seguintes comandos:  

```sh
git clone git@github.com:gildemberg-santos/webcrawlerurl_v2.git
cd webcrawlerurl_v2
```

### 2. Construção e Execução da Aplicação com Docker
**Construindo a Imagem Docker**
Execute o seguinte comando para criar a imagem Docker:

```sh
docker build -t webcrawlerurl_v2 .
```

**Executando o Container**
Para iniciar o container e executar a aplicação, utilize:

```sh
docker run -d --name webcrawlerurl_v2 -p 8080:8080 webcrawlerurl_v2
```

**Interagindo com o Container**
Caso precise acessar o ambiente interno do container:

```sh
docker exec -it webcrawlerurl_v2 /bin/bash
```

**Parar e Remover Containers**
Finalizar um container ativo:

```sh
docker container kill webcrawlerurl_v2
```

Remover containers parados e liberar espaço:

```sh
docker container prune -f
```

**Visualizando Logs do Container**
Para consultar logs do serviço em execução:

```sh
docker logs webcrawlerurl_v2
```
