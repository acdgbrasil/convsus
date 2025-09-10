# Dicionário de Dados: Mapeamento Convsus -> Prontuário SUAS

**Versão:** 2.0
**Data:** 2025-09-08
**Fonte:** Manual de Instruções para Utilização do Prontuário SUAS (2014)

**Objetivo:** Este documento detalha a correspondência entre os campos dos schemas JSON do sistema Convsus e os campos e códigos oficiais do Prontuário SUAS.

---

## Bloco: Identificação e Endereço

**Origem:** `reference_person_data.json`

| Campo Original (JSON) | Campo Destino (SUAS) | Tipo (DBF) | Descrição | Observações / Códigos Oficiais |
|---|---|---|---|---|
| `fullName` | `NM_PESSOA_REF` | Character(100) | Nome completo da pessoa de referência. | | 
| `socialName` | `NM_SOCIAL` | Character(100) | Nome social, se aplicável. | | 
| `motherName` | `NM_MAE` | Character(100) | Nome completo da mãe. | | 
| `nis` | `NU_NIS` | Character(11) | Número de Identificação Social. | | 
| `cpf` | `NU_CPF` | Character(11) | Cadastro de Pessoa Física. | Remover formatação. | 
| `birthDate` | `DT_NASC` | Date | Data de nascimento. | | 
| `biologicalGender` | `CD_SEXO` | Character(1) | Código do sexo biológico. | Mapear: 'Masculino' -> 'M', 'Feminino' -> 'F'. | 
| `localLocalization` | `CD_ZONA_RESID` | Character(1) | Código da zona de residência. | Mapear: 'URBAN' -> '1', 'RURAL' -> '2'. | 

---

## Bloco: Composição Familiar

**Origem:** `family_composition_data.json`

| Campo Original (JSON) | Campo Destino (SUAS) | Tipo (DBF) | Descrição | Observações / Códigos Oficiais |
|---|---|---|---|---|
| `familyCompositionPerson.kinship` | `CD_PARENTESCO` | Character(2) | Relação de parentesco. | **1**: Pessoa de Referência, **2**: Cônjuge, **3**: Filho(a), **4**: Enteado(a), **5**: Neto(a)/Bisneto(a), **6**: Pai/Mãe, **7**: Sogro(a), **8**: Irmão/Irmã, **9**: Genro/Nora, **10**: Outro parente, **11**: Não parente. |
| `familyCompositionPerson.educationConditionPerson.schoolShip` | `CD_ESCOLARIDADE` | Character(2) | Escolaridade (última série concluída). | **00**: Nunca frequentou, **01**: Creche, **02**: Educação Infantil, **11-19**: 1º ao 9º ano do Ens. Fundamental, **21-23**: 1º ao 3º ano do Ens. Médio, **30**: Superior Incompleto, **31**: Superior Completo, **40**: EJA Fundamental, **41**: EJA Médio, **99**: Outros. |

---

## Bloco: Condições Habitacionais

**Origem:** `home_conditions_data.json`

| Campo Original (JSON) | Campo Destino (SUAS) | Tipo (DBF) | Descrição | Observações / Códigos Oficiais |
|---|---|---|---|---|
| `typeResidence` | `CD_TIPO_DOMIC` | Character(1) | Tipo de residência. | **1**: Própria, **2**: Alugada, **3**: Cedida, **4**: Ocupada. |
| `materialOfExternalWalls` | `CD_MAT_PAREDE` | Character(1) | Material predominante nas paredes. | **1**: Alvenaria/madeira aparelhada, **2**: Madeira aproveitada/taipa/outros. |
| `hasAcessEnergy` | `CD_ENERGIA_ELET` | Character(1) | Acesso à energia elétrica. | **1**: Com medidor próprio, **2**: Com medidor comunitário, **3**: Sem medidor, **4**: Não possui. |

---

## Bloco: Condições de Trabalho e Rendimento

**Origem:** `work_condition_data.json`

| Campo Original (JSON) | Campo Destino (SUAS) | Tipo (DBF) | Descrição | Observações / Códigos Oficiais |
|---|---|---|---|---|
| `workConditionPerson.workCondition` | `CD_COND_TRAB` | Character(2) | Condição de Ocupação. | **0**: Não trabalha, **1**: C. Própria, **2**: Temp. Rural, **3**: Empregado s/ carteira, **4**: Empregado c/ carteira, **5**: Doméstico s/ carteira, **6**: Doméstico c/ carteira, **7**: Não remunerado, **8**: Militar/Serv. Público, **9**: Empregador, **10**: Estagiário, **11**: Aprendiz. |

---

## Bloco: Situações de Violência e Violação de Direitos

**Origem:** `family_situation_violence_data.json`

| Campo Original (JSON) | Campo Destino (SUAS) | Tipo (DBF) | Descrição | Observações / Códigos Oficiais |
|---|---|---|---|---|
| `childLabel.thisSituationOcurrent` | `IN_VIOL_TRAB_INF` | Character(1) | Ocorrência de Trabalho Infantil. | **1**: Sim, **2**: Não. (Mapear `true`/`false`) |
| `sexualExploitation.thisSituationOcurrent` | `IN_VIOL_EXPL_SEX` | Character(1) | Ocorrência de Exploração Sexual. | **1**: Sim, **2**: Não. |
| `physicalAbuse.thisSituationOcurrent` | `IN_VIOL_FISICA` | Character(1) | Ocorrência de Violência Física. | **1**: Sim, **2**: Não. |
| `psychologicalAbuse.thisSituationOcurrent` | `IN_VIOL_PSICO` | Character(1) | Ocorrência de Violência Psicológica. | **1**: Sim, **2**: Não. |
| `homelessSituation.thisSituationOcurrent` | `IN_SIT_RUA` | Character(1) | Vivência de Situação de Rua. | **1**: Sim, **2**: Não. |

---

## Bloco: Benefícios Eventuais

**Origem:** `family_evently_benefits_data.json`

| Campo Original (JSON) | Campo Destino (SUAS) | Tipo (DBF) | Descrição | Observações / Códigos Oficiais |
|---|---|---|---|---|
| `typeOfBenefit` | `CD_TIPO_BENEF_EVENT` | Character(1) | Tipo do benefício eventual concedido. | **1**: Natalidade, **2**: Funeral, **3**: Kit Emergência, **4**: Cesta Básica, **5**: Aluguel Social, **6**: Outros. | 
