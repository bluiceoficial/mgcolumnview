# MGColumnView

[![License](https://img.shields.io/badge/license-PolyForm%20Perimeter%201.0.1-5351FB)](LICENSE.md)

**MGColumnView** é um componente customizado para **Fyne (Go)** que implementa uma **visualização tabular baseada em colunas**, com suporte a:

- Cabeçalhos clicáveis (ordenação)
- Larguras fixas por coluna
- Scroll vertical
- Seleção por checkbox (linha individual ou selecionar todos)
- Manipulação dinâmica de linhas (adicionar, atualizar, remover)
- Recuperação de dados selecionados ou completos

É ideal para aplicações desktop que precisam de algo mais flexível que `widget.Table`, mantendo simplicidade e controle total.

---

## ✨ Características

- Componente customizado
- Cabeçalhos clicáveis para **ordenação por coluna**
- Suporte opcional a **checkbox por linha**
- Checkbox **Selecionar Todos**
- Largura fixa por coluna
- Scroll vertical automático
- Atualização parcial eficiente (reconstrói apenas o corpo)
- API simples e previsível
- Totalmente escrito em Go puro

---

## 📦 Instalação

```bash
go get github.com/profmugomes/mgcolumnview
```

---

## 🚀 Uso básico

```go
headers := []string{"Nome", "Email", "Status"}
widths := []float32{150, 250, 100}

cv := mgcolumnview.NewColumnView(headers, widths, true)
```

Adicionando ao layout:

```go
container.NewVBox(cv)
```

---

## ➕ Adicionando linhas

```go
cv.AddRow([]string{"João", "joao@email.com", "Inativo"})
cv.AddRow([]string{"Maria", "maria@email.com", "Ativo"})
```

Se faltar alguma coluna, o componente preenche automaticamente com string vazia.

---

## ✏️ Atualizando dados

### Atualizar uma linha inteira

```go
cv.UpdateItem(0, []string{"João", "novo@email.com", "Ativo"})
```

---

### Atualizar uma coluna específica

```go
cv.UpdateColumnItem(1, 2, "Ativo")
```

---

## ❌ Removendo linhas

### Remover linhas selecionadas

```go
cv.RemoveSelected()
```

---

### Remover todas as linhas

```go
cv.RemoveAll()
```

---

## ☑️ Seleção de linhas

### Recuperar linhas selecionadas

```go
selected := cv.ListSelected()

for _, row := range selected {
    fmt.Println(row.ID, row.Data)
}
```

---

### Recuperar todas as linhas

```go
all := cv.ListAll()
```

---

## 🔃 Ordenação por coluna

Clicar no **título da coluna** realiza ordenação ascendente com base no conteúdo textual.

A ordenação é estável e preserva a ordem relativa de valores iguais.

---

## 🧩 Compatibilidade

* Go 1.26.5+
* Fyne 2.8.0

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved.

This project is licensed under the PolyForm Perimeter License 1.0.1.

### Summary

This software is available for commercial and noncommercial use, subject to the terms of the PolyForm Perimeter License 1.0.1.

You may:

* ✔ Use the software for commercial and noncommercial purposes.
* ✔ Inspect and study the source code.
* ✔ Modify the software.
* ✔ Create derivative works based on the software.
* ✔ Redistribute the software and permitted modifications.

You may not:

* ✖ Provide a product that competes with the software.

See the full license terms at LICENSE.md.

This summary is provided for convenience only and does not replace or modify the full license terms.
