package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var subjects map[string]map[string]float64
var students map[string]map[string]float64

type Alumnos struct {
	Name    string
	Subject string
	Grade   string
}

type AdminServer struct {
	Info []Alumnos
}

func (info *AdminServer) Add(data Alumnos) {
	info.Info = append(info.Info, data)
}

var auxStudents AdminServer

func AddGrade(data []string, reply *string) error {
	if _, err := students[data[0]][data[1]]; err {
		*reply = "Error. Ya se tiene una calificacion registrada"
		return nil
	}
	subjectAux := make(map[string]float64)
	studentAux := make(map[string]float64)
	gradeAux, _ := strconv.ParseFloat(data[2], 64)
	subjectAux[data[1]] = gradeAux
	studentAux[data[0]] = gradeAux

	if _, err := students[data[0]]; err {
		students[data[0]][data[1]] = gradeAux
		if _, err := subjects[data[1]]; err {
			subjects[data[1]][data[0]] = gradeAux
		} else {
			subjects[data[1]] = studentAux
		}
	} else {
		students[data[0]] = subjectAux
		if _, err := subjects[data[1]]; err {
			subjects[data[1]][data[0]] = gradeAux
		} else {
			subjects[data[1]] = studentAux
		}
	}
	*reply = "Se agrego correctamente la calificacion del alumno " + data[0]
	return nil
}

func AverageGradeStudent(name string, reply *string) error {
	if _, err := students[name]; err {
		var sum float64
		var subject int
		for _, grade := range students[name] {
			sum = sum + grade
			subject = subject + 1
		}
		finalGrade := sum / float64(subject)
		*reply = "La calificacion del estudiante " + name + " es: " + strconv.FormatFloat(finalGrade, 'f', 2, 64)
		return nil
	} else {
		*reply = "Error. No existe ese alumno."
		return nil
	}
}

func AverageGradeSubject(subject string, reply *string) error {
	if _, err := subjects[subject]; err {
		var sum float64
		var counter int
		for _, grade := range subjects[subject] {
			sum = sum + grade
			counter = counter + 1
		}
		finalGrade := sum / float64(counter)
		*reply = "El promedio general de la materia " + subject + " es: " + strconv.FormatFloat(finalGrade, 'f', 2, 64)
		return nil
	} else {
		*reply = "Error. No existe esa materia"
		return nil
	}
}

func AverageGradeAll(all int, reply *string) error {
	var sum float64
	var cont int
	for student := range students {
		for _, grade := range students[student] {
			sum = sum + grade
			cont = cont + 1
		}
	}
	finalGrade := sum / float64(cont)
	*reply = "El promedio general es: " + strconv.FormatFloat(finalGrade, 'f', 0, 64)
	return nil
}

func alumnos(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		if req.RequestURI == "/alumnos" {
			fmt.Println(req.PostForm)
			data := Alumnos{Name: req.FormValue("nombre"), Subject: req.FormValue("materia"), Grade: req.FormValue("calificacion")}
			var info []string
			info = append(info, data.Name)
			info = append(info, data.Subject)
			info = append(info, data.Grade)
			var result string
			err := AddGrade(info, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				if result != "Error. Ya se tiene una calificacion registrada" {
					auxStudents.Add(data)
				}
			}
			fmt.Println(auxStudents)
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				result,
			)
		}
		if req.RequestURI == "/promedioAlumno" {
			fmt.Println(req.PostForm)
			data := req.FormValue("nombre")
			var result string
			err := AverageGradeStudent(data, &result)
			if err != nil {
				fmt.Println(err)
			}
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				result,
			)
		}
		if req.RequestURI == "/promedioMateria" {
			fmt.Println(req.PostForm)
			data := req.FormValue("nombre")
			var result string
			err := AverageGradeSubject(data, &result)
			if err != nil {
				fmt.Println(err)
			}
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				result,
			)
		}
	case "GET":
		var all int
		var result string
		err := AverageGradeAll(all, &result)
		if err != nil {
			fmt.Println(err)
		}
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioGeneral.html"),
			result,
		)
	}
}

func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)
	return string(html)
}

func main() {
	subjects = make(map[string]map[string]float64)
	students = make(map[string]map[string]float64)
	http.HandleFunc("/form", form)
	http.HandleFunc("/alumnos", alumnos)
	http.HandleFunc("/promedioAlumno", alumnos)
	http.HandleFunc("/promedioMateria", alumnos)
	http.HandleFunc("/promedioGeneral", alumnos)
	fmt.Println("Corriendo servidor...")
	http.ListenAndServe(":9000", nil)
}
