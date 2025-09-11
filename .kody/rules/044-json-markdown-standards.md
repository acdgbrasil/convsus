---
title: "JSON/Markdown padronizados: 2 espaços, LF, válido"
scope: "pull_request"
path: ["**/*.json", "**/*.md"]
severity_min: "low"
bu  ckets: ["automation", "style"]
enabled: true
---

## Instructions
- JSON: identação 2 espaços, válido (sem trailing commas).
- Markdown: títulos consistentes, listas com `-`, linhas até ~120 colunas.
- Sem espaços ao final de linha; sempre newline ao final do arquivo.

## Examples

### Bad example
```
Schemas/example.json // com tabs e vírgula sobrando
```

### Good example
```
Schemas/example.json // formatado com 2 espaços e JSON válido
```

