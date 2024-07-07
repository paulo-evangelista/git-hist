package main

import (
	"fmt"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
    // Abra o repositório
    repo, err := git.PlainOpen(".") // "." indica o diretório atual
    if err != nil {
        fmt.Println("Erro ao abrir o repositório:", err)
        return
    }

    // Pegue o worktree (árvore de trabalho) para adicionar e commitar arquivos
    worktree, err := repo.Worktree()
    if err != nil {
        fmt.Println("Erro ao obter worktree:", err)
        return
    }

    // Adiciona todos os arquivos modificados para staging
    err = worktree.AddWithOptions(&git.AddOptions{All: true})
    if err != nil {
        fmt.Println("Erro ao adicionar arquivos:", err)
        return
    }

	commitDate, err := time.Parse("2006-01-02 15:04:05", "2024-07-07 12:00:00")
    if err != nil {
        fmt.Println("Error getting commit date:", err)
        return
    }

    // Faça o commit
    commitMsg := ":)"
    _, err = worktree.Commit(commitMsg, &git.CommitOptions{
        Author: &object.Signature{
            Name:  "Paulo Evangelista",
            Email: "paulo.evangelista@sou.inteli.edu.br",
            When:  commitDate,
        },
    })
    if err != nil {
        fmt.Println("Erro ao commitar:", err)
        return
    }

    fmt.Println("Commit realizado com sucesso")
}