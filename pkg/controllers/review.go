package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/aglide100/chicken_review_webserver/pkg/db"
	"github.com/aglide100/chicken_review_webserver/pkg/models"
	"github.com/aglide100/chicken_review_webserver/pkg/views"
)

type ReviewController struct {
	db                *db.Database
	APIKeys           *models.APIKeys
	SessionController *SessionController
}

func NewReviewController(db *db.Database, APIKeys *models.APIKeys, SessionController *SessionController) *ReviewController {
	return &ReviewController{db: db, APIKeys: APIKeys, SessionController: SessionController}
}

func findString(resp http.ResponseWriter, req *http.Request, str string) (id int, orderType string, pagenumber int) {
	var matches []string
	// pagenumber
	Pnumber := 0

	var showReviewPattern = regexp.MustCompile("^/reviews/([0-9]+$)")
	var deleteReviewPattern = regexp.MustCompile("^/delete/([0-9]+$)")
	var updateReviewPattern = regexp.MustCompile("^/update/([0-9]+$)")
	var uploadUpdateReviewPattern = regexp.MustCompile("^/reviews/update/upload/([0+9+])$")
	var listReviewPattern = regexp.MustCompile("^/reviews/([A-Z]+)-pagenumber=([0-9]+)$")

	switch str {
	case "Show":
		matches = showReviewPattern.FindStringSubmatch(req.URL.Path)
	case "Delete":
		matches = deleteReviewPattern.FindStringSubmatch(req.URL.Path)
	case "Update":
		matches = updateReviewPattern.FindStringSubmatch(req.URL.Path)
	case "UploadUpdate":
		matches = uploadUpdateReviewPattern.FindStringSubmatch(req.URL.Path)
	case "List":
		matches = listReviewPattern.FindStringSubmatch(req.URL.Path)
	}

	if len(matches) != 2 {
		//http.Error(resp, "no ID provided", http.StatusBadRequest)

		if str == "List" {
			//var err error
			/*
				orderType = matches[3]
				pagenumber, err = strconv.Atoi(matches[20]) // 20
				if err != nil {
					log.Printf("PageNumber is not numeric: %v", err)
				}
			*/

			return 0, orderType, pagenumber
		}

		return
	}

	idStr := matches[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//http.Error(resp, fmt.Sprintf("ID is not numeric: %v", err.Error()), http.StatusBadRequest)
		return
	}

	return id, "", Pnumber
}

func (hdl *ReviewController) GetScript(resp http.ResponseWriter, req *http.Request) {
	//log.Printf("[review_func]: receive request to get script")

	view := views.NewReviewGetScriptView(views.DefaultBaseHTMLContext, req.URL.Path)
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) GetAssets(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to get assets")

	view := views.NewReviewGetAssetsView(views.DefaultBaseHTMLContext, req.URL.Path)
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) GetImage(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to get image")

	view := views.NewReviewGetImageView(views.DefaultBaseHTMLContext, req.URL.Path, "ReviewImage") // subdivided to ReviewImage and Favico
	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}

}

func (hdl *ReviewController) Create(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to create a review")

	CheckUser := hdl.SessionController.GetUserDataInSession(req)
	view := views.NewReviewCreateView(views.DefaultBaseHTMLContext, hdl.APIKeys, CheckUser)

	resp.Header().Set("Content-Type", view.ContentType())
	err := view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) Revise(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to update a reivew")

	id, _, _ := findString(resp, req, "Update")

	review, ok, err := hdl.db.GetReview(id)
	if err != nil {
		log.Printf("finding review: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	var view views.View
	if !ok {
		view = views.NewNotFoundView(views.DefaultBaseHTMLContext)
	} else {
		CheckUser := hdl.SessionController.GetUserDataInSession(req)
		view = views.NewReviewUpdateView(views.DefaultBaseHTMLContext, review, CheckUser)
	}

	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) Delete(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to delete a review")

	id, _, _ := findString(resp, req, "Delete")
	log.Printf("Delete Review ID :%v", id)
	hdl.db.DeleteReview(id)

	http.Redirect(resp, req, "/reviews", 301)
}

const (
	KiB = 1 << 10
	MiB = 1024 * KiB

	maxImageSize = 20 * MiB
)

func SaveImage(resp http.ResponseWriter, req *http.Request, hdl *ReviewController) (string, []string, bool, error) {
	// Saving image in local directory
	var PictureURLs []string
	PictureURLs = nil
	var DefaultPictureURL string

	err := req.ParseMultipartForm(500000) // grab the multipart form
	if err != nil {
		return "", PictureURLs, false, fmt.Errorf("Can't grab file: %v", err)
	}

	formdata := req.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["image"] // grab the filenames

	log.Printf("Start reading files")
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return "", PictureURLs, false, fmt.Errorf("looking up image from form file: %v", err)
		}

		imageBytes, err := ioutil.ReadAll(io.LimitReader(file, maxImageSize))
		if err != nil {
			return "", PictureURLs, true, fmt.Errorf("can't read image data: %v", err)
		}

		fileType := http.DetectContentType(imageBytes)
		switch fileType {
		case "image/jpeg", "image/png", "image/gif", "image/webp":
			fileType = strings.Replace(fileType, "image/", ".", 1)
			// ok
		default:
			return "", PictureURLs, true, fmt.Errorf("invalid image format: %q", fileType)
		}

		basePath := "ui/img/"

		reviewid, err := hdl.db.GetLastInsertReviewID()
		if err != nil {
			log.Printf("Can't get LastInsertId!", err)
		}

		IDstr := strconv.FormatInt(reviewid, 10)
		FileNum := strconv.Itoa(i)
		//FileNum = filepath.Join(FileNum, "st")
		log.Printf("reivewID: %v, IDstr: %v", reviewid, IDstr) //
		currentReviewBasePath := filepath.Join(basePath, IDstr, "/", FileNum)
		//currentReviewBasePath := basePath
		log.Printf("Review image path:%v", currentReviewBasePath)

		err = os.MkdirAll(currentReviewBasePath, os.ModePerm)
		if err != nil {
			return "", PictureURLs, true, fmt.Errorf("creatig directory for image: %v", err)
		}

		if err != nil {
			fmt.Fprintln(resp, err)
			return "", PictureURLs, true, fmt.Errorf("can't read image data: %v", err)
		}

		imageFilename := files[i].Filename
		log.Printf("file name :%v", imageFilename)
		IDstr += fileType
		//currentReviewImagePath := filepath.Join(currentReviewBasePath, IDstr, imageFilename) // Can't create directory
		currentReviewImagePath := filepath.Join(currentReviewBasePath, IDstr)

		if err := ioutil.WriteFile(currentReviewImagePath, imageBytes, 0644); err != nil {
			return "", PictureURLs, true, fmt.Errorf("creating image file on disk: %v", err)
		}

		str := currentReviewImagePath

		log.Printf("img path: %v", currentReviewImagePath)
		PictureURLs = append(PictureURLs, str)
		//return currentReviewImagePath, true, nil
		if i == 0 {
			DefaultPictureURL = currentReviewImagePath
		}
	}

	log.Printf("img path: %v", DefaultPictureURL)
	return DefaultPictureURL, PictureURLs, true, nil
}

func SaveReview(resp http.ResponseWriter, req *http.Request, hdl *ReviewController, ReviewType string) (*models.Review, error, bool, string) {
	log.Printf("[review_func]: receive request save Review")
	var PictureURLs []string

	path, PictureURLs, ok, err := SaveImage(resp, req, hdl)
	if !ok {
		log.Printf("There are no image!")
	} else if err != nil {
		log.Printf("saving image: %v", err)
		http.Error(resp, "can't save image", http.StatusInternalServerError)
		return nil, err, false, ""
	}

	checklist := [...]string{
		"store_name",
		"date",
		"phone_number",
		"author",
		"title",
		"comment",
		"lat",
		"lng",
		"addr",
	}

	blacklist := [...]string{
		"<",
		">",
		"$",
		"<script>",
		"<style>",
		"$func",
		"<empty>",
	}

	review := &models.Review{
		StoreName:         req.PostFormValue("store_name"),
		Date:              req.PostFormValue("date"),
		PhoneNumber:       req.PostFormValue("phone_number"),
		Author:            req.PostFormValue("author"),
		Title:             req.PostFormValue("title"),
		DefaultPictureURL: path,
		PictureURLs:       PictureURLs,
		Comment:           req.PostFormValue("comment"),
		Lat:               req.PostFormValue("lat"),
		Lng:               req.PostFormValue("lng"),
		Addr:              req.PostFormValue("addr"),
		//UpdateDate:        req.PostFormValue("write_date"),
	}

	checklistnum := 9
	blacklistnum := 7

	// Check blakclist
	for i := 0; i < checklistnum; i++ {
		for k := 0; k < blacklistnum; k++ {
			// Check threat
			if (checklist[i] == "lat") || (checklist[i] == "lng") || (checklist[i] == "addr") {
				if blacklist[k] == "<empty>" {
					continue
					// skip
				}
			}
			result := strings.Replace(req.PostFormValue(checklist[i]), blacklist[k], "Alert", -1)

			if result != req.PostFormValue(checklist[i]) {
				log.Printf("When saving review, found alert!")
				return nil, err, true, result
			}
		}
	}
	switch ReviewType {
	case "Save":

	case "Update":

	}

	return review, nil, false, ""
}

func (hdl *ReviewController) Save(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to save a review")

	review, err, xss, str := SaveReview(resp, req, hdl, "Save")
	if err != nil {
		if xss {
			log.Printf("%v", str)
			http.Redirect(resp, req, "/Don't-Use-Script-or-CSS-at-review", 301)
		} else {
			log.Fatal("Can't save review %v:", err)
		}
	}

	// return to default page
	if !xss {
		hdl.db.CreateReview(review)
		http.Redirect(resp, req, "/reviews", 301)
	}
}

func (hdl *ReviewController) Update(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to update a review")

	id, _ := strconv.Atoi(req.PostFormValue("id"))

	review, err, xss, str := SaveReview(resp, req, hdl, "Save")
	if err != nil {
		if xss {

			log.Printf("%v", str)
			http.Redirect(resp, req, "/Don't-Use-Script-or-Css-at-review", 301)
		} else {
			log.Fatal("Can't save review")
		}
	}

	if !xss {
		hdl.db.UpdateReview(review, id)
		http.Redirect(resp, req, "/reviews", 301)
	}
}

func (hdl *ReviewController) Search(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to search review")

	req.ParseForm()

	var (
		name     string
		subject  string
		operator string
	)
	name = req.FormValue("name")
	subject = req.FormValue("subject")
	operator = req.FormValue("operator")

	reviews, err := hdl.db.SearchReviews(name, subject, operator)
	if err != nil {
		log.Printf("listing reviews: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	CheckUser := hdl.SessionController.GetUserDataInSession(req)
	view := views.NewReviewSearchView(views.DefaultBaseHTMLContext, reviews, CheckUser)
	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed to render : %v", err)
	}
}

func (hdl *ReviewController) Get(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to get a review")

	id, _, _ := findString(resp, req, "Show")

	review, ok, err := hdl.db.GetReview(id)
	if err != nil {
		log.Printf("finding review: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	var view views.View
	if !ok {
		view = views.NewNotFoundView(views.DefaultBaseHTMLContext)
	} else {
		CheckUser := hdl.SessionController.GetUserDataInSession(req)
		view = views.NewReviewShowView(views.DefaultBaseHTMLContext, review, hdl.APIKeys, CheckUser)
	}
	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed to render: %v", err)
	}
}

func (hdl *ReviewController) List(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[review_func]: receive request to list reviews")

	_, orederType, pagenumber := findString(resp, req, "List")

	reviews, err := hdl.db.ListReviews(orederType, pagenumber)
	if err != nil {
		log.Printf("listing reviews: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	CheckUser := hdl.SessionController.GetUserDataInSession(req)
	log.Printf("CheckUser: %v", CheckUser)
	view := views.NewReviewListView(views.DefaultBaseHTMLContext, reviews, CheckUser)
	resp.Header().Set("Content-Type", view.ContentType())
	err = view.Render(resp)
	if err != nil {
		log.Printf("failed to render : %v", err)
	}
}
