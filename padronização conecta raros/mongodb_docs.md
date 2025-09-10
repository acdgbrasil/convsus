# Documentação do Schema e Correlação do Banco de Dados MongoDB

**Versão:** 1.0
**Data:** 2025-09-08

**Objetivo:** Este documento fornece uma visão geral da arquitetura do banco de dados MongoDB, detalhando a estrutura das coleções e como elas se relacionam. A informação é baseada nos arquivos `mongodb_schema.json` e `mongodb_correlation.json`.

---

## 1. Visão Geral da Correlação de Dados

O banco de dados é centrado na coleção `referencePerson`, que funciona como o documento principal para cada família ou indivíduo atendido. Todas as outras coleções contêm informações detalhadas e são vinculadas a um documento `referencePerson` através de um campo de ID específico.

As relações são as seguintes:

- **`referencePerson`** (Coleção Principal)
  - `fistEntryInUnityId` -> **`firstEntryInUnity`**: Detalhes sobre o primeiro ingresso na unidade de serviço.
  - `familyCompositionId` -> **`familyComposition`**: Descreve todos os membros da família.
  - `homeConditionsId` -> **`homeConditions`**: Detalha as condições de moradia.
  - `workConditionId` -> **`workCondition`**: Condições de trabalho e renda da família.
  - `familySituationViolationId` -> **`familySituationViolence`**: Histórico de situações de violência.
  - `helphyConditionId` -> **`helphyCondition`**: Detalha as condições de saúde da família.
  - `eventlyBenefitId` -> **`familyEventlyBenefits`**: Registra os benefícios eventuais recebidos.
  - `familyAndCommunityId` -> **`familyAndCommunity`**: Informações sobre os vínculos familiares e comunitários.
  - `familyHistoryOfComplianceSocialEducationalMensuresId` -> **`familyHistoryOfComplianceSocioEducationalMeasures`**: Histórico de cumprimento de medidas socioeducativas.
  - `familyHistoryInstitutionalCompletId` -> **`FamilyHistoryInstitutionalComplet`**: Histórico de acolhimento institucional.

---

## 2. Detalhamento dos Schemas das Coleções

A seguir, uma descrição dos campos para cada coleção principal, conforme definido em `mongodb_schema.json`.

### Tabela: `referencePerson`
| Campo | Tipo | Descrição |
|---|---|---|
| `fullName` | string | Nome completo. |
| `socialName` | string | Nome social. |
| `motherName` | string | Nome da mãe. |
| `nis` | string | Número de Identificação Social. |
| `cpf` | string | CPF. |
| `diagnosis` | string | Diagnóstico de saúde. |
| `rg` | object | Objeto contendo os dados do RG. |
| `isShelter` | boolean | Indica se a pessoa está em um abrigo. |
| `localLocalization` | string | Localização (URBAN/RURAL). |
| `cep` | string | CEP do endereço. |
| `adress` | string | Logradouro. |
| `...` | ... | (e outros campos de endereço e IDs de referência) |

### Tabela: `familyAndCommunity`
| Campo | Tipo | Descrição |
|---|---|---|
| `yearsInState` | number | Anos de residência no estado. |
| `awaysLivingInState` | boolean | Sempre residiu no estado. |
| `hasVictimOfThreatsOrDiscrimination` | boolean | Indica se já foi vítima de ameaças. |
| `relationshipEvaluationByTechnician` | string | Avaliação da relação familiar pelo técnico. |
| `...` | ... | (e outros campos relacionados) |

### Tabela: `familyComposition`
| Campo | Tipo | Descrição |
|---|---|---|
| `socialEspecification` | string | Especificação social (ex: Indígena). |
| `familyCompositionPerson` | array | Lista de objetos, onde cada objeto é um membro da família. |
| `observation` | array | Lista de observações. |

### Tabela: `homeConditions`
| Campo | Tipo | Descrição |
|---|---|---|
| `typeResidence` | string | Tipo de residência (Própria, Alugada, etc.). |
| `materialOfExternalWalls` | string | Material das paredes externas. |
| `hasAcessEnergy` | string | Tipo de acesso à energia elétrica. |
| `numberOfRooms` | number | Número de cômodos. |

*(Nota: As demais tabelas seguem uma estrutura similar, detalhando aspectos específicos da vida da família, como saúde, renda, histórico de violência e benefícios.)*
