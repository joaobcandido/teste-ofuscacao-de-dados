package main // Define o pacote principal do programa

import (
"archive/zip"   // Para descompactar arquivos ZIP
"fmt"           // Para formatação e impressão de textos
"io"            // Para operações de entrada/saída
"net/http"      // Para fazer requisições HTTP
"os"            // Para manipulação de arquivos e diretórios
"path/filepath" // Para manipulação de caminhos de arquivos
)

func main() {
// Define o dono do repositório no GitHub
owner := "joaobcandido"
// Define o nome do repositório
repo := "posto-de-gasolina"
// Define o branch a ser baixado
branch := "main"
// Monta a URL para baixar o ZIP do repositório
zipURL := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", owner, repo, branch)

// 1. Cria um diretório temporário exclusivo para armazenar arquivos temporários
tmpDir, err := os.MkdirTemp("", "repo-zip-*")
if err != nil {
	panic(err) // Encerra o programa em caso de erro
}
defer os.RemoveAll(tmpDir) // Garante que o diretório temporário será removido ao final

// Define o caminho do arquivo ZIP dentro do diretório temporário
zipPath := filepath.Join(tmpDir, "repo.zip")

// 2. Baixa o arquivo ZIP do repositório
resp, err := http.Get(zipURL)
if err != nil {
	panic(err) // Encerra o programa em caso de erro na requisição
}
defer resp.Body.Close() // Garante que o corpo da resposta HTTP será fechado

// Cria o arquivo local para salvar o ZIP baixado
out, err := os.Create(zipPath)
if err != nil {
	panic(err) // Encerra o programa em caso de erro ao criar o arquivo
}
// Copia o conteúdo do corpo da resposta HTTP para o arquivo ZIP local
_, err = io.Copy(out, resp.Body)
out.Close() // Fecha o arquivo ZIP local
if err != nil {
	panic(err) // Encerra o programa em caso de erro ao copiar os dados
}

// 3. Define o diretório onde o ZIP será descompactado
unzipDir := filepath.Join(tmpDir, "unzipped")
// Chama a função para descompactar o ZIP
err = unzip(zipPath, unzipDir)
if err != nil {
	panic(err) // Encerra o programa em caso de erro na descompactação
}

// 4. Lê o conteúdo de um arquivo específico (exemplo: README.md)
// O nome do diretório raiz é repo-branch (ex: posto-de-gasolina-main)
repoRoot := filepath.Join(unzipDir, fmt.Sprintf("%s-%s", repo, branch))
// Monta o caminho completo do arquivo README.md
readmePath := filepath.Join(repoRoot, "README.md")
// Lê o conteúdo do arquivo README.md
data, err := os.ReadFile(readmePath)
if err != nil {
	fmt.Println("README.md não encontrado:", err) // Informa se o arquivo não foi encontrado
} else {
	fmt.Println("Conteúdo do README.md:") // Imprime o conteúdo do README.md
	fmt.Println(string(data))
}

// 5. Todos os arquivos temporários serão removidos pelo defer acima
}

// Função para descompactar um arquivo ZIP para um diretório de destino
func unzip(src, dest string) error {
// Abre o arquivo ZIP para leitura
r, err := zip.OpenReader(src)
if err != nil {
	return err // Retorna erro se não conseguir abrir o ZIP
}
defer r.Close() // Garante que o arquivo ZIP será fechado ao final

// Itera sobre todos os arquivos e diretórios dentro do ZIP
for _, f := range r.File {
	// Monta o caminho completo do arquivo/diretório de destino
	fpath := filepath.Join(dest, f.Name)
	// Se for um diretório, cria o diretório
	if f.FileInfo().IsDir() {
		os.MkdirAll(fpath, os.ModePerm)
		continue // Vai para o próximo item
	}
	// Garante que o diretório do arquivo existe
	if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
		return err // Retorna erro se não conseguir criar o diretório
	}
	// Cria o arquivo de destino
	outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err // Retorna erro se não conseguir criar o arquivo
	}
	// Abre o arquivo dentro do ZIP para leitura
	rc, err := f.Open()
	if err != nil {
		outFile.Close()
		return err // Retorna erro se não conseguir abrir o arquivo do ZIP
	}
	// Copia o conteúdo do arquivo do ZIP para o arquivo de destino
	_, err = io.Copy(outFile, rc)
	outFile.Close() // Fecha o arquivo de destino
	rc.Close()      // Fecha o arquivo do ZIP
	if err != nil {
		return err // Retorna erro se não conseguir copiar os dados
	}
}
return nil // Retorna nil se tudo ocorreu bem
}
