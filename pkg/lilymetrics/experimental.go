package lilymetrics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)
var allowOriginFunc = func(r *http.Request) bool {
	return true
}
func escapeSingleQuotes(s string) string {
	return fmt.Sprintf("%s", s)
}
func _Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}
func runCowSay(msg string) bool {
	log.Println("notice:", msg)

	// cmdstd := exec.Command("cat", "/tmp/lilypad/data/downloaded-files/QmYqdpiry4h9P39JTZ65NjhhS2QQou448LWRckE96cpkxY/stdout")

	// output, err := cmdstd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	//return true
	// }

	// fmt.Println(string(output))
	// server.BroadcastToNamespace("/", "reply", output, "data")
	// return false
	server.BroadcastToNamespace("/", "reply", msg, "data")

	cmd := exec.Command("./stack", "run-cowsay-onchain")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		r := scanner.Text()
		fmt.Println("from scanner ", r)
		// if strings.Contains(r, "Job result:") {
		// 	words := strings.Fields(r)
		// 	lastWord := words[len(words)-1]
		// 	fmt.Println("JOB " + lastWord)
		// 	cmdstd := exec.Command("cat", "/tmp/lilypad/data/downloaded-files/"+lastWord+"/stdout")

		// 	output, err := cmdstd.Output()
		// 	fmt.Println(string(output))
		// 	server.BroadcastToNamespace("/", "reply", output, "data")
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		return true
		// 	}

		// 	break
		// } else {
		server.BroadcastToNamespace("/", "reply", r, "data")
		// }

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return false
}
func (e Event) String() string {
	// convert the Event to a string
	// for example, you could marshal it to JSON:
	bytes, err := json.Marshal(e)
	if err != nil {
		return "error marshaling Event to JSON: " + err.Error()
	}
	return string(bytes)
}

	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cmd := exec.Command("bacalhau serve --node-type compute,requester --peer none --private-internal-ipfs=false --job-selection-accept-networked")

	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }
// defer Trace(context.Background(), "Migrations").End()

	// go exec.Command("bacalhau", "serve", " --node-type compute,requester --peer none --private-internal-ipfs=false --job-selection-accept-networked")
	// // cmdstd.Start()
	// // output, errr := cmdstd.Output()
	// // if errr != nil {
	// // 	fmt.Println("Error:", errr)
	// // }
	// // fmt.Println(string(output))
	// time.Sleep(5 * time.Second)
	// go exec.Command("./stack", "solver")
	// go exec.Command("./stack", "mediator")
	// go exec.Command("./stack", "resource-provider", "--offer-gpu 1")
	// go exec.Command("./stack", "jobcreator")
	
	// router.Handle("/socket.io/", func(context echo.Context) error {
	// 	server.ServeHTTP(context.Response(), context.Request())
	// 	return nil
	// })

		// http.Handle("/socket.io/", server)