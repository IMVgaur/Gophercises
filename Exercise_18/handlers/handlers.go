package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mitchellh/go-homedir"

	"github.com/IMVgaur/Gophercises/Exercise_18/primitive"
)

//Handler ...
func Handler() http.Handler {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(imgPath))
	mux.HandleFunc("/", Welcome)
	mux.HandleFunc("/upload", Upload)
	mux.HandleFunc("/modify/", Modify)
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	return mux
}

//Welcome ... 
//Prints welcome text for home page
func Welcome(w http.ResponseWriter, r *http.Request) {
	html := `<html>
		<body>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="image"/>
			<input type="submit" value="upload"/>
		</form></body>
		</html>`
	fmt.Fprint(w, html)
}

//Upload ...
//Uploads the image and contains persitence logic
func Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
	}

	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	outFile, err := tempFile("", ext)
	if err != nil {
		http.Error(w, "New File can not be created", http.StatusBadRequest)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	http.Redirect(w, r, "/modify/"+filepath.Base(outFile.Name()), http.StatusFound)
}

//Modify ...
//This function contains the logic of transformation
func Modify(w http.ResponseWriter, r *http.Request) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	file, err := os.Open(imgPath + "/" + filepath.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()
	ext := filepath.Ext(file.Name())[1:]
	modeStr := r.FormValue("mode")
	if modeStr == "" {
		renderModeChoices(w, r, file, ext)
		return
	}
	mode, err := strconv.Atoi(modeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	n := r.FormValue("n")
	if n == "" {
		renderNumShapeChoices(w, r, file, ext, primitive.Mode(mode))
		return
	}
	http.Redirect(w, r, "/img/"+filepath.Base(file.Name()), http.StatusFound)
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	var err error
	opts := []genOpts{
		{N: 10, M: primitive.ModeCircle},
		{N: 10, M: primitive.ModeBeziers},
		{N: 10, M: primitive.ModePolygon},
		{N: 10, M: primitive.ModeCombo},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err == nil {
		html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		type dataStruct struct {
			Name string
			Mode primitive.Mode
		}
		var data []dataStruct
		for i, img := range imgs {
			data = append(data, dataStruct{
				Name: filepath.Base(img),
				Mode: opts[i].M,
			})
		}
		err = tpl.Execute(w, data)
		return
	}
	http.Error(w, "error occured in render mode choices", http.StatusInternalServerError)
}

func renderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) {
	var err error
	opts := []genOpts{
		{N: 10, M: mode},
		{N: 20, M: mode},
		{N: 30, M: mode},
		{N: 40, M: mode},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err == nil {
		html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		type dataStruct struct {
			Name      string
			Mode      primitive.Mode
			NumShapes int
		}
		var data []dataStruct
		for i, img := range imgs {
			data = append(data, dataStruct{
				Name:      filepath.Base(img),
				Mode:      opts[i].M,
				NumShapes: opts[i].N,
			})
		}
		err = tpl.Execute(w, data)
		return
	}
	http.Error(w, "Error occured in rendering number of choices", http.StatusInternalServerError)
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	var err error
	var f string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err = genImage(rs, ext, opt.N, opt.M)
		if err == nil {
			ret = append(ret, f)
		}
	}
	return ret, err
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	var outFile *os.File
	var err error
	var out io.Reader
	out, err = primitive.Transform(r, ext, numShapes, primitive.WithMode(mode))
	if err == nil {
		outFile, err = tempFile("", ext)
		if err == nil {
			defer outFile.Close()
			io.Copy(outFile, out)
			return outFile.Name(), err
		}
	}
	return "", err
}

/*
*tempFile : creates file for holding data
*input : name of file and extension : string
*return : os.File and error
**/
func tempFile(name, ext string) (*os.File, error) {
	home, err := homedir.Dir()
	imgPath := filepath.Join(home, "img")
	file, err := ioutil.TempFile(imgPath+"/", name)
	if err != nil {
		return nil, errors.New("failed to create new temp file")
	}
	defer os.Remove(file.Name())
	f, err := os.Create(fmt.Sprintf("%s.%s", file.Name(), ext))
	if err != nil {
		return nil, err
	}
	return f, nil
}
