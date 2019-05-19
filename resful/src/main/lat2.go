func users(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
if r.Method == "POST" {
var result, err = json.Marshal(data)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.Write(result)
return
}
http.Error(w, "", http.StatusBadRequest)
}