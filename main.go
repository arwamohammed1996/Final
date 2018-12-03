package main

import (
    "fmt"
   "html/template"
 "encoding/json"
   "log"
  
  "strconv"
   "io/ioutil"
    "net/http"
      "github.com/gorilla/mux"
   
)
 





var tpl *template.Template

func init (){
    tpl =template.Must(template.ParseGlob("templates/*.html"))
}

func main() {


	  r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/process", processor)
	http.Handle("/", r)
	
	log.Fatal(http.ListenAndServe(":9090", nil))
   
   // http.HandleFunc("/", index) // setting router rule
  // http.HandleFunc("/process",processor)
   // err := http.ListenAndServe(":9090", nil) // setting listening port
    //if err != nil {
       // log.Fatal("ListenAndServe: ", err)
    //}
}

    func index (w http.ResponseWriter, r *http.Request){
           tpl.ExecuteTemplate(w,"index.html",nil)
    }



func processor (w http.ResponseWriter, r *http.Request){
    
        if r.Method !="POST"{
            http.Redirect(w,r,"/",http.StatusSeeOther)
            return
        }
        amount:=r.FormValue("amount")
        from:=r.FormValue("from")
        to:=r.FormValue("to")
        
    
        response, err := http.Get("https://api.exchangeratesapi.io/latest?symbols="+from+","+to)
if err != nil {
fmt.Printf("The HTTP request failed with error %s\n", err)
} else {
    
data, _ := ioutil.ReadAll(response.Body)

var objmap map[string]*json.RawMessage
err := json.Unmarshal(data, &objmap)
fmt.Println(err)

var dat map[string]interface{}
if err := json.Unmarshal(data, &dat); err != nil {
panic(err)
}
fmt.Println(dat)

date := dat["date"].(string)
fmt.Println(date)

base := dat["base"].(string)
fmt.Println(base)

var rate map[string]float64
err = json.Unmarshal(*objmap["rates"], &rate)
fmt.Println(rate)

fmt.Println(from)

value:=rate[from]
t:=rate[to]
fmt.Println(value)  
fmt.Println(t) 

val, err := strconv.ParseFloat(amount, 32)
if err != nil {
    // do something sensible
}


conv :=(t*val)/value


// visualisation api



d:=struct{
    Amount string
    From string
    To string
    Json float64
   R1 float64
   R2 float64
}{
   
    Amount:amount,
    From:from,
    To:to,
    Json:conv,
    R1:t,
    R2:value,

}

tpl.ExecuteTemplate(w,"processor.html",d)
}



       
    }