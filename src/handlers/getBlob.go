package handlers

import (
	"app/src/auth"
	"app/src/gcp"
	"fmt"
	"net/http"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	if auth.CheckAuth(r) {
		bucket := r.Header.Get("BUCKET")
		blob := r.Header.Get("BLOB")
		if bucket == "" || blob == "" {
			fmt.Fprintf(w, "Error: Invalid Headers")
			w.WriteHeader(400)
			return
		}
		data, err := gcp.GetBlobDataStorageBucket(bucket, blob)
		if err != nil {
			message := "Could Not Get Blob: " + blob + " From Bucket:" + bucket
			fmt.Fprintf(w, message+err.Error())
			return
		}
		w.WriteHeader(200)
		w.Write(data)
		return
	}
	fmt.Fprintf(w, "Error: Invalid Auth Credentials")
	w.WriteHeader(400)
}
