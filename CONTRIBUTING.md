Contribuindo com o CONVSUS

Obrigado por apoiar uma tecnologia feita no Brasil, para o Brasil. Este projeto é mantido pela ACDG (Associação Brasileira de Profissionais Atuantes em Doenças Genéticas, Pacientes, Familiares e Voluntários) e pela comunidade.

Visão rápida
- Começamos pequeno, com foco no essencial: interoperar SUAS × DATASUS gerando arquivos compatíveis (.DBF, .DEF, .CNV) e, depois, adicionar uma camada de IA (RAG) para explicar dados em linguagem natural.
- O core é em Go (binários leves, multiplataforma, testes simples). Wrappers em outras linguagens são bem‑vindos.
- Valorizamos contribuições incrementais, bem testadas e com documentação clara.

Regras de conduta e licenças
- Respeito em primeiro lugar. Seguimos etiqueta comunitária em todas as interações.
- Licenças:
  - Código: PolyForm Noncommercial 1.0.0 (uso não comercial). Uso comercial (incluindo SaaS) exige acordo separado com a ACDG. Veja `LICENSE-CODE.txt`.
  - Documentação/imagens: CC BY‑NC‑SA 4.0. Veja `LICENSE-DOCS.txt`.
- Ao contribuir, você concorda em licenciar sua contribuição sob os mesmos termos aplicáveis à parte do repositório que você alterou (código ou documentação).

Como começar
1) Abra uma issue
- Descreva problema/feature, contexto, exemplos e impacto.
- Use os templates em `.github/ISSUE_TEMPLATE/` para agilizar a triagem.

2) Crie um PR pequeno e objetivo
- Siga o template em `.github/PULL_REQUEST_TEMPLATE.md`.
- Explique o “por quê”, o “como” e como validar.
- Inclua testes (table‑driven no Go) e exemplos quando possível.

3) Padrões técnicos
- Go (core):
  - Formatação: `gofmt`/`goimports`.
  - Lint: manter `golangci-lint` limpo (quando configurado no CI).
  - Erros: sem `panic` no core; prefira `error`, `errors.Is/As` e mensagens claras.
  - Concurrency: evitar globais; use `-race` nos testes locais quando possível.
  - `context.Context` em operações canceláveis.
- JSON/Markdown:
  - 2 espaços, LF, newline ao final de arquivo.
  - JSON válido (sem trailing comma). Markdown claro e conciso.

4) Estrutura de diretórios (proposta)
- `cmd/convsus/` — CLI.
- `internal/transform/` — motor de transformação (puro, sem I/O).
- `internal/format/{dbf,def,cnv}/` — I/O de formatos.
- `bindings/{python,node,rust,java}/` — wrappers.
- `testdata/` — fixtures/golden (usa `Schemas/` como fonte de exemplos).
- `docs/` — guias e notas técnicas.

5) Qualidade e revisão
- PRs passam por revisão humana e pelo Kody (Kodus) — Code Review automático.
- Se necessário, comente `@kody start-review` no PR para disparar uma revisão.

Contato
- E‑mail: gabriel.aderaldo@acdgbrasil.com.br
- Site ACDG: https://acdgbrasil.com.br/
- Preferiu Git? Abra uma issue neste repositório.

Solicitação de licença comercial / exceções
- Caso sua organização precise de permissão de uso comercial (por exemplo, oferecer SaaS), siga o modelo em `docs/COMMERCIAL_EXCEPTION_REQUEST.md` e envie para gabriel.aderaldo@acdgbrasil.com.br.

Obrigado por construir tecnologia pública e útil com a gente.
