---
title: "Go Core: erros, concorrência e testes confiáveis"
scope: "pull_request"
path: ["cmd/**", "internal/**"]
severity_min: "high"
bu  ckets: ["go", "quality-gates", "reliability"]
enabled: true
---

## Instructions
- Sem `panic` no core. Propague `error` com contexto; use `errors.Is/As` quando fizer sentido.
- Operações canceláveis devem aceitar `context.Context` como primeiro parâmetro.
- Evite estado global; preferir injeção de dependência e valores imutáveis.
- Testes table-driven; execute com `-race` localmente e garanta determinismo.

## Examples

### Bad example
```
func WriteDBF(path string, rows []Row) {
    if len(rows) == 0 { panic("no data") }
    // ...
}
```

### Good example
```
func WriteDBF(ctx context.Context, path string, rows []Row) error {
    if len(rows) == 0 { return fmt.Errorf("no data to write") }
    // ...
    return nil
}
```

