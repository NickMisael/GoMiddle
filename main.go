package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "user0"
	dbname = "teste"
)

func limpaTela() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

func telaInicial() {
	fmt.Println("1-Cadastrar")
	fmt.Println("2-Consultar")
	fmt.Println("3-Atualizar")
	fmt.Println("4-Deletar")
	fmt.Println("5-Sair")
	fmt.Printf("-> ")
}

func Atualizar() {
	/*
		sqlUpdate := `
		UPDATE pessoa
		SET `*/
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/* TESTE DE CONEXÃO COM O BANCO
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado com sucesso!")*/
	limpaTela()
	for {
		var es string
		telaInicial()
		fmt.Scanf("%s", &es)

		if esc, er := strconv.Atoi(es); er != nil || esc > 5 || esc < 1 {
			fmt.Println("Digite um Número válido!!")
			time.Sleep(time.Second + 3)
			limpaTela()
		} else {
			if esc == 5 {
				fmt.Println("Obrigado por utilizar o Programa! :D")
				break
			} else {
				switch esc {
				case 1:
					limpaTela()
					var first, last, email string
					var age int

					fmt.Printf("Digite o Primeiro Nome: ")
					fmt.Scanf("%s", &first)
					fmt.Printf("Digite o Sobrenome: ")
					fmt.Scanf("%s", &last)
					fmt.Printf("Digite a Idade: ")
					fmt.Scanf("%d", &age)
					fmt.Printf("Digite o Email: ")
					fmt.Scanf("%s", &email)

					/* INSERINDO OS DADOS*/
					sqlInsert := `
					INSERT INTO pessoa (age, first_name, last_name, email)
					VALUES ($1, $2, $3, $4)
					RETURNING id`
					id := 0
					err := db.QueryRow(sqlInsert, age, first, last, email).Scan(&id)
					if err != nil {
						panic(err)
					}

					limpaTela()
					fmt.Println("Novo registro ID é:", id)
					fmt.Println("Dados adicionados com Sucesso! :D")
					time.Sleep(time.Second + 3)
					limpaTela()
				case 2:
					for {
						limpaTela()
						type User struct {
							ID        int
							Age       int
							FirstName string
							LastName  string
							Email     string
						}
						var line int
						fmt.Printf("Digite a linha a ser consultada: ")
						fmt.Scanf("%d", &line)

						/* CONSULTANDO */
						sqlQuery := `SELECT * FROM pessoa WHERE id=$1;`
						var user User
						row := db.QueryRow(sqlQuery, line)
						err := row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
						switch err {
						case sql.ErrNoRows:
							fmt.Println("Linha não encontrada")
							return
						case nil:
							fmt.Println("ID:", user.ID)
							fmt.Println("Nome:", user.FirstName, user.LastName)
							fmt.Println("Age:", user.Age)
							fmt.Println("Email:", user.Email)
						default:
							panic(err)
						}
						var sair string
						fmt.Println("Digite 0 para sair")
						fmt.Scanf("%s", &sair)
						if sair == "0" {
							time.Sleep(time.Second + 2)
							limpaTela()
							fmt.Println("Voltando")
							time.Sleep(time.Second + 2)
							limpaTela()
							break
						}
						time.Sleep(time.Second + 2)
						limpaTela()
					}
				case 3:
					Atualizar()
				case 4:
					limpaTela()
					var line int
					fmt.Printf("Digite a linha a ser deletada: ")
					fmt.Scanf("%d", &line)
					sqlDelete := `
					DELETE FROM pessoa
					WHERE id=$1;`
					res, err := db.Exec(sqlDelete, line)
					if err != nil {
						panic(err)
					}
					limpaTela()
					count, err := res.RowsAffected()
					if err != nil {
						panic(err)
					}
					fmt.Println(count)
					fmt.Println("Linha deletada com sucesso!")
					time.Sleep(time.Second + 2)
					limpaTela()
				}
			}
		}
	}
}
