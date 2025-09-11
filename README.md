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
