---
title: "PRs pequenos, descritos e validados"
scope: "pull_request"
path: ["**/*"]
severity_min: "medium"
bu  ckets: ["quality-gates", "process"]
enabled: true
---

## Instructions
- Prefira PRs pequenos e temáticos, com descrição do problema, solução e como validar.
- Inclua passos de validação e, quando aplicável, artefatos de teste (fixtures/golden).
- Se o PR for grande por necessidade, explique o porquê e como revisar por partes.

## Examples

### Bad example
```
"feat: big refactor"
// sem descrição, sem validação
```

### Good example
```
"feat: writer .DEF e testes"
- Problema: export não gera .DEF
- Solução: implementa writer c/ validações
- Validação: abrir em TabWin (passos listados), testes passados
```

