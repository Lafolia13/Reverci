package reversi;
import (
  "io"
  "io/ioutil"
  "fmt"
  "os/exec"
  "strconv"
  "strings"

  "api/request"
)

func GetSolvers() (solvers request.Solvers) {
  files, err := ioutil.ReadDir("/solvers")
  if err != nil {
    fmt.Println(err)
  }

  solvers.Solver = make([]string, 1)
  solvers.Solver[0] = "manual"
  for _, file := range files {
    if !file.IsDir() {
      solvers.Solver = append(solvers.Solver, file.Name())
    }
  }

  return
}

func RunSolver(gameData *request.GameData) bool {
  path := "/solvers/" + gameData.Status

  cmd := exec.Command(path)
  stdin, err := cmd.StdinPipe()
  if err != nil {
    fmt.Println(err)
    return false
  }

  io.WriteString(stdin, strconv.Itoa(gameData.Turn) + "\n")
  io.WriteString(stdin, strconv.Itoa(len(gameData.Field)) + "\n")
  for _, raw := range gameData.Field {
    for _, value := range raw {
      io.WriteString(stdin, strconv.Itoa(value) + " ")
    }
    io.WriteString(stdin, "\n")
  }

  out, err := cmd.Output()
  if err != nil {
    fmt.Println(err)
    return false
  }

  r := strings.NewReader(string(out))
  fmt.Fscanf(r, "%d %d", &gameData.Request.H, &gameData.Request.W)

  return true
}