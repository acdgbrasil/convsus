# CONVSUS — Interoperabilidade SUAS × DATASUS, por e para o Brasil

Quando um profissional de saúde ou de assistência social precisa analisar dados, a exigência costuma ser a mesma: “funcionar no que já existe”. No Brasil, isso significa falar fluentemente o idioma do DATASUS — TabWin e TabNet — sem exigir que ninguém reaprenda ferramentas, formatos ou fluxos.

O CONVSUS nasce para isso: utilidade pública, open source, que transforma dados socioassistenciais (SUAS) e correlatos em arquivos compatíveis com o ecossistema DATASUS (.DBF, .DEF, .CNV) e — em uma segunda camada — explica em linguagem natural o que significam os dados estatísticos e registros, com foco especial em doenças raras.

Construir tecnologias brasileiras para o Brasil não é um slogan: é uma necessidade. Queremos respeitar o legado, reduzir atrito no dia a dia e abrir caminho para inovação, sempre com utilidade real na ponta.

---

## Para quem é

- Gestores e times de saúde/assistência social em municípios, estados e União.
- ONGs, conselhos e coletivos que monitoram políticas públicas.
- Pesquisadores, jornalistas de dados e universidades.
- Comunidade open source interessada em interoperabilidade de dados públicos.

## ONG responsável

- Associação Brasileira de Profissionais Atuantes em Doenças Genéticas, Pacientes, Familiares e Voluntários – ACDG (ACDG)
- Site: https://acdgbrasil.com.br/
- Por que este projeto importa: interoperabilidade é inclusão. Ao reduzir atrito técnico (formatação, conversão, padronização), ampliamos a capacidade de análise e advocacy em doenças genéticas e raras, fortalecendo o SUS e o SUAS com base em evidências.

---

## O que é (agora)

Uma biblioteca/CLI open source focada em:

- Ler dados padronizados em JSON (pasta `Schemas/`).
- Transformar, achatar e validar conforme mapeamentos oficiais (DATASUS + IBGE).
- Exportar arquivos `.DBF` (dados), `.DEF` (estrutura), `.CNV` (dicionários) compatíveis com TabWin/TabNet — sem que ninguém precise reaprender ferramentas.

Uma camada de IA (opcional, próxima fase) para responder, em linguagem natural, perguntas sobre os dados — usando indexação e recuperação de contexto (RAG) com conjuntos limitados e gerenciáveis (ex.: 1000 registros de referência).

---

## Por que começar em Go

- Linguagem simples, compilada e multiplataforma — binários leves e fáceis de distribuir.
- Excelente performance para manipulação de dados/arquivos sem dependências pesadas.
- Ecossistema maduro para CLIs, testes e CI (gofmt, golangci-lint, govulncheck).
- Porta aberta para crescer depois:
  - Wrappers em múltiplas linguagens (Python/Node/Java/Rust).
  - API gRPC para SDKs.
  - FFI/C-ABI quando for necessário embutir.

Começamos pragmáticos com Go, mas queremos incentivar e, quando fizer sentido, adotar soluções totalmente brasileiras: serviços, modelos de IA e ferramentas mantidas pela comunidade do Brasil.

---

## Agradecimento à Kodus (AI Code Review)

Entramos em contato com a Kodus e eles aceitaram apoiar o projeto com Code Review automático (Kody). Nosso muito obrigado pela rapidez, pela abertura e por fortalecer um ecossistema de tecnologia feito no Brasil. Vamos integrar o serviço para acelerar a qualidade das contribuições e manter o padrão do core em Go e dos wrappers. Se por qualquer motivo a integração não estiver disponível em algum momento, seguimos com alternativas, mas estamos muito felizes em contar com a Kodus desde já.

---

## Roadmap (realista e incremental)

### Fase 1 — Core e Exportadores (0.x)

- Engine de transformação (Go), puro e testável: mapeia campos, achata estruturas (flatten) e aplica transformações/dicionários.
- Exportadores DATASUS: escrita de `.DBF` (dBase III), `.DEF` (descrição de campos) e `.CNV` (dicionários), compatíveis com TabWin/TabNet.
- CLI `convsus`: subcomandos `transform`, `export`, `validate`; entrada/saída via arquivos e `stdin/stdout`.
- Testes table‑driven, golden files e validação real no TabWin/TabNet.

### Fase 2 — Interoperabilidade e SDKs

- API gRPC e geração de SDKs (Python/Node/Java/Rust).
- Wrappers simples chamando a CLI como primeira etapa, evoluindo para clientes gRPC.

### Fase 3 — Camada de IA (RAG)

- Indexação de um universo de referência (ex.: 1000 registros) — foco em relevância, não em “Big Data”.
- API de perguntas/respostas que monta contexto enxuto para o LLM.
- PoC com modelo nacional (Magalu GPT) e, conforme custo/latência, migração opcional para modelos locais leves (TinyLLaMA, GPT‑J, LLaMA 2, Phi‑3 Mini).

### Fase 4 — Operacionalização

- Releases multiplataforma, exemplos reais e documentação ampliada.
- Observabilidade básica e automação de QA.

---

## Arquitetura (visão lógica)

```mermaid
flowchart LR
  subgraph Sources
    S1[JSON padronizado\nSchemas/]:::file
    S2[Mapeamentos oficiais\nDATASUS/IBGE]:::file
  end

  subgraph Core Go
    C1[Transform Engine (puro)\nmapeia/achata/valida]:::go
    C2[Exporters\n.DBF .DEF .CNV]:::go
    CLI[CLI convsus]:::cli
  end

  subgraph Interoperabilidade
    T1[TabWin]:::tool
    T2[TabNet]:::tool
  end

  subgraph IA (RAG)
    I1[Normalização p/ índice]:::proc
    I2[Indexação\nvetorial/keyword]:::store
    I3[Q&A API / Wrappers]:::api
    I4[LLM Provider\nMagalu GPT / Local]:::llm
  end

  subgraph Operação & Qualidade
    Q1[pre-commit / linters]:::qa
    Q2[Kody Code Review]:::qa
    Q3[govulncheck / tests -race]:::qa
  end

  S1 --> C1
  S2 --> C1
  C1 --> C2
  C2 --> T1
  C2 --> T2

  C1 --> |registros achatados| I1
  I1 --> I2
  I2 --> I3
  I3 --> I4

  CLI --> C1
  CLI --> C2
  W1[[Wrappers: Python/Node/Rust/Java]]:::api --> |chama CLI/gRPC| C1
  W1 --> |export| C2
  I3 --> |SDKs| W1

  Q1 -.-> C1
  Q1 -.-> C2
  Q2 -.-> C1
  Q3 -.-> C1

  classDef go fill:#e8f5e9,stroke:#2e7d32,color:#1b5e20;
  classDef cli fill:#e3f2fd,stroke:#1565c0,color:#0d47a1;
  classDef tool fill:#fff3e0,stroke:#ef6c00,color:#e65100;
  classDef file fill:#f3e5f5,stroke:#6a1b9a,color:#4a148c;
  classDef api fill:#ede7f6,stroke:#4527a0,color:#311b92;
  classDef store fill:#fbe9e7,stroke:#bf360c,color:#bf360c;
  classDef proc fill:#f1f8e9,stroke:#558b2f,color:#33691e;
  classDef llm fill:#ffebee,stroke:#c62828,color:#b71c1c;
  classDef qa fill:#eceff1,stroke:#455a64,color:#263238;
```

---

## Qualidade e governança

- Formatação e validação:
  - JSON e Markdown padronizados (2 espaços, LF, newline no EOF).
  - pre-commit: `trailing-whitespace`, `end-of-file-fixer`, `mixed-line-ending`, `check-json`, `pretty-format-json`.
- Go (core):
  - `gofmt`/`goimports` obrigatório; `golangci-lint` e `govulncheck` no CI.
  - Tratamento de erros sem `panic`; `context.Context` em operações canceláveis; testes table‑driven e `-race` no CI.
  - Fixtures em `Schemas/` e golden files em `testdata/`.
- Code Review automático (Kodus/Kody):
  - Integração aceita e em andamento. Regras editoriais (Kody Rules) e revisão automática em PRs para manter padrão e velocidade.

---

## Estado atual do repositório

- `Schemas/` com exemplos reais de estrutura de dados.
- Documentos e mapeamentos exploratórios (DATASUS/IBGE).
- `docs/tasks/` com critérios BDD para core e exportadores.
- Core em Go em desenvolvimento (estrutura sendo preparada).

---

## Como contribuir

- Issues: descreva problema/feature com contexto e exemplo.
- PRs: pequenos, com motivação e forma de validação; preferir testes table‑driven.
- Padrões (resumo):
  - Go: `gofmt`/`goimports`; `golangci-lint` limpo; erros sem `panic`; `context` onde couber; `-race` nos testes.
  - JSON/Markdown padronizados; rodar pre‑commit antes do push.
- Boas primeiras contribuições:
  - Writers de `.DEF`/`.CNV` e testes de compatibilidade com TabWin.
  - CLI `transform/export/validate`.
  - Mapeamentos para domínios adicionais (com dicionários oficiais).
  - Exemplos de wrappers (Python/Node) chamando a CLI.

---

## Licenças e governança

- Mantido pela ACDG (https://acdgbrasil.com.br/) e comunidade.
- Código (software): PolyForm Noncommercial License 1.0.0 — uso, modificação e distribuição permitidos para fins não comerciais. Uso comercial (inclusive SaaS) requer licença separada com a ACDG. Veja `LICENSE-CODE.txt`.
- Documentação, textos e imagens: CC BY‑NC‑SA 4.0 — reuso não comercial com atribuição e compartilhamento pela mesma licença. Veja `LICENSE-DOCS.txt`.
- Código de Conduta: em preparação; seguimos etiqueta comunitária e respeito absoluto a contribuidoras(es).

---

## FAQ

- É só para saúde?
  - Não. O foco é SUAS × Saúde, mas o pipeline é extensível a outras bases públicas.
- Funciona sem GPU?
  - Sim. O core em Go é CPU‑bound. IA é opcional e incremental.
- Windows/Linux/macOS?
  - Sim. Go compila binários multiplataforma.

---

## Agradecimentos

- Profissionais, pacientes, familiares e voluntários que inspiram este trabalho.
- Comunidade open source brasileira.
- Kodus/Kody, pela parceria e por investir em tecnologia feita no Brasil.
