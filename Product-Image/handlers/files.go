package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/kevin/product-image/files"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
)

type Files struct {
	log   hclog.Logger
	store files.Storage
}

func NewFiles(l hclog.Logger, s files.Storage) *Files {
	return &Files{log: l, store: s}
}

func (f *Files) UploadMultipart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(w, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	id, idErr := strconv.Atoi(r.FormValue("id"))
	f.log.Info("Process form for id", "id", id)
	if idErr != nil {
		http.Error(w, "Expected integer id", http.StatusBadRequest)
		return
	}

	ff, mh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(w, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	f.saveFile(r.FormValue("id"), mh.Filename, w, ff)

}

func (f *Files) UploadRest(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle Post", "id", id, "filename", fn)

	f.saveFile(id, fn, writer, request.Body)

}

func (f *Files) saveFile(id, path string, writer http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(writer, "Unable to save file", http.StatusInternalServerError)
	}
}
