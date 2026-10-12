# 🚀 Project M (Nome Provisório)

## 📌 Contexto
- **Protagonista:** Um robô autônomo com o objetivo supremo de dominar áreas, planetas e o mundo todo.
- **Gênero:** Automação, Tower Defense e Ação com visão superior (Top-Down).
- **Estilo Visual:** Pixel Art (32x32 pixels), com foco em legibilidade e paleta de cores coesa.

---

## 🛠️ Roadmap de Desenvolvimento e Tarefas (Em Ordem)

### 🟢 Fase 1: Prototipagem Básica (Core Loop)
*Objetivo: Movimentação, colisão e coleta manual.*
- [ ] **Configuração da Engine (Ebitengine):**
    - [ ] Definir resolução lógica virtual (ex: `480x270`) com escala *Pixel-Perfect* para a janela.
    - [ ] Configurar loop principal e sistema básico de renderização.
- [ ] **Movimentação do Robô:**
    - [ ] Implementar controles de movimento (8 direções ou livre com colisão).
    - [ ] Criar sprite placeholder (32x32) do robô.
- [ ] **Mundo e Geração Inicial:**
    - [ ] Criar sistema de Grid / Tiles (32x32).
    - [ ] Gerar um mapa procedural básico com biomas simples e recursos no chão (ex: Ferro e Cobre).
- [ ] **Coleta Manual:**
    - [ ] Permitir que o robô interaja com os nós de recursos no chão para coletá-los manualmente para o inventário.

### 🟡 Fase 2: Automação e Economia
*Objetivo: Parar de minerar na mão e construir a infraestrutura.*
- [ ] **Sistema de Construção (Base Building):**
    - [ ] Interface/atalho para selecionar e posicionar construções no grid.
    - [ ] Consumo de recursos do inventário ao construir.
- [ ] **Extratores e Processamento:**
    - [ ] Mineradoras automáticas (colocam recursos em esteiras ou armazenam localmente).
    - [ ] Transformadores (processam minério bruto em barras/placas).
- [ ] **Logística e Energia:**
    - [ ] Esteiras transportadoras (ou tubos) para mover itens entre máquinas.
    - [ ] Geradores de energia e rede de distribuição (máquinas param se faltar energia).

### 🟠 Fase 3: Defesa e Combate (Tower Defense)
*Objetivo: Proteger a base contra ameaças.*
- [ ] **Sistemas Defensivos:**
    - [ ] Torres de ataque automáticas (aliram-se e atiram nos inimigos próximos).
    - [ ] Muros reforçados e armadilhas terrestres.
- [ ] **IA Inimiga e Ameaças:**
    - [ ] Inimigos naturais (animais selvagens/criaturas) ou colônias de robôs hostis.
    - [ ] Sistema de ondas pseudo-aleatórias de ataque às bases do jogador.
- [ ] **Combate do Personagem:**
    - [ ] Armas equipáveis para o robô principal se defender ativamente.

### 🔵 Fase 4: Expansão e Progressão (Meta-Game)
*Objetivo: Conquistar o mundo e outros planetas.*
- [ ] **Estação de Pesquisa:**
    - [ ] Árvore de tecnologias para desbloquear novas torres, melhorias de máquinas e armas avançadas.
- [ ] **Planetas e Recursos Únicos:**
    - [ ] Diferenciação de planetas (materiais exclusivos e climas variados que afetam as máquinas).
    - [ ] Mecânica de viagem espacial / foguete para transitar entre mundos.