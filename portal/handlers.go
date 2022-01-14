package portal

import (
	"context"
	"distributed/grpc/grades"
	"distributed/grpc/grades/pb"
	"distributed/grpc/registry"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	http.Handle("/", http.RedirectHandler("/students", http.StatusPermanentRedirect))

	h := new(studentsHandler)
	http.Handle("/students", h)
	http.Handle("/students/", h)
}

type studentsHandler struct{}

func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2: // /students
		sh.renderStudents(w, r)
	case 3: // /students/{:id}
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderStudent(w, r, id)
	case 4: // /students/{:id}/grades
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.ToLower(pathSegments[3]) != "grades" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderGrades(w, r, id)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (studentsHandler) renderStudents(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error retrieving students: ", err)
		}
	}()

	serviceURL, err := registry.GetProvider(registry.GRADES_SERVICE)
	conn,err := grpc.Dial(serviceURL, grpc.WithInsecure())
	fmt.Println(serviceURL)
	fmt.Println(err)

	if err != nil {
		return
	}

	client := pb.NewGradesServiceClient(conn)

	out, err := client.GetAllStudents(context.Background(), &pb.GetAllRequest{})

	if err != nil {
		return
	}

	var s grades.Students
	for _, stu := range out.Data {
		dataGrades := []grades.Grade{}
		for _, stuGrades := range stu.Grades {
			dataGrade := grades.Grade{
				Title: stuGrades.Title,
				Type:  grades.GradeType(stuGrades.Type),
				Score: stuGrades.Score,
			}
			dataGrades = append(dataGrades, dataGrade)
		}
		dataStu := grades.Student{
			ID:        int(stu.Id),
			FirstName: stu.FirstName,
			LastName:  stu.LastName,
			Grades:    dataGrades,
		}
		s = append(s, dataStu)
	}

	rootTemplate.Lookup("students.html").Execute(w, s)
}

func (studentsHandler) renderStudent(w http.ResponseWriter, r *http.Request, id int) {

	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error retrieving students: ", err)
			return
		}
	}()

	serviceURL, err := registry.GetProvider(registry.GRADES_SERVICE)
	if err != nil {
		return
	}

	conn,err := grpc.Dial(serviceURL, grpc.WithInsecure())
	if err != nil {
		return
	}

	client := pb.NewGradesServiceClient(conn)
	out, err := client.GetOneStudent(context.Background(), &pb.IdRequest{Id:int32(id)})
	if err != nil {
		return
	}
	var s grades.Student
	dataGrades := []grades.Grade{}
	for _, stuGrades := range out.Data.Grades {
		dataGrade := grades.Grade{
			Title: stuGrades.Title,
			Type:  grades.GradeType(stuGrades.Type),
			Score: stuGrades.Score,
		}
		dataGrades = append(dataGrades, dataGrade)
	}
	s = grades.Student{
		ID:        int(out.Data.Id),
		FirstName: out.Data.FirstName,
		LastName:  out.Data.LastName,
		Grades:    dataGrades,
	}

	rootTemplate.Lookup("student.html").Execute(w, s)
}

func (studentsHandler) renderGrades(w http.ResponseWriter, r *http.Request, id int) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		w.Header().Add("location", fmt.Sprintf("/students/%v", id))
		w.WriteHeader(http.StatusTemporaryRedirect)
	}()
	title := r.FormValue("Title")
	gradeType := r.FormValue("Type")
	score, err := strconv.ParseFloat(r.FormValue("Score"), 32)
	if err != nil {
		log.Println("Failed to parse score: ", err)
		return
	}
	g := &pb.Grade{
		Title: title,
		Type:  gradeType,
		Score: float32(score),
	}


	serviceURL, err := registry.GetProvider(registry.GRADES_SERVICE)
	if err != nil {
		log.Println("Failed to retrieve instance of Grading Service", err)
		return
	}
	conn,err := grpc.Dial(serviceURL, grpc.WithInsecure())
	fmt.Println(serviceURL)
	fmt.Println(err)

	if err != nil {
		return
	}

	client := pb.NewGradesServiceClient(conn)

	_, err = client.AddGrade(context.Background(), &pb.GradeRequest{
		Id: int32(id),
		Grade:g,
	})
	if err != nil {
		log.Println("Failed to save grade to Grading Service", err)
		return
	}

}
