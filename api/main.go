import ("main.go"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"package"

)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println ("api is on :8080")	
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
}


func getUsers (w http.ResponseWriter, r *http.Request) {
	
if r.Method ≠ "POST"{
	http.error(w, http.StausText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
w.Header().Set("Content=Type", "application/json")
json.NewEncoder(w).Encode([]User{
	{
	ID:1,
	Name: "Rafael",
},

{
	ID: 2,
	Name: "Pipoca",
}})
}