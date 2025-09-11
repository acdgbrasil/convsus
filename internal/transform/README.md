Motor de Transformação (puro)

Responsável por mapear, achatar e validar registros conforme mapeamentos (DATASUS/IBGE), sem realizar I/O.

Requisitos de qualidade:
- Sem `panic`; retornar `error` com contexto.
- `context.Context` quando houver cancelamento/timeouts.
- Testes table‑driven; determinismo; considerar `-race`.

