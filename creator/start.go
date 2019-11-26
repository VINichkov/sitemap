package creator

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"sitemap/models"
	"sitemap/models/repository"
	"sitemap/viewes"
)

const NumberOfItems int = 10000
const Path string = "public/sitemap/"


type Creator struct {
	company   *repository.DbCompanyRepo
	job       *repository.DbJobRepo
	resume    *repository.DbResumeRepo
	companywithjobs *repository.DbCompanyWithJobsRepo

}

func NewCreator(db *sqlx.DB) *Creator  {
	return &Creator{
		company:   repository.NewSQLCompanyRepo(db),
		job:       repository.NewSQLdbJobRepo(db),
		resume:    repository.NewSQLdbResumeRepo(db),
		companywithjobs: repository.NewSQLCompanyWithJobsRepo(db),
	}
}

func (r *Creator) Start(){
	var maxJobsPage, maxResumesPage, maxCompaniesPage, maxCompaniesWithJobs int
	var host string = os.Getenv("HOST")
	if host[len(host)-1] == '/' {
		 host = host[:len(host)-1]
	}
	err := r.max_pages_in_resource(
		&maxJobsPage,
		&maxResumesPage ,
		&maxCompaniesPage ,
		&maxCompaniesWithJobs,
		)

	if err != nil {
		return
	}

	numberOfFiles := maxJobsPage + maxResumesPage + maxCompaniesPage + maxCompaniesWithJobs

	index := 1

	//jobs := viewes.NewJobXML(NumberOfItems)

	//create_files(maxJobsPage, &index, jobs)


	indexXML := viewes.NewIndexXML(numberOfFiles, host)
	err = ioutil.WriteFile(Path+ "sitemap.xml", indexXML.XML(), 0644)
	if err != nil{
		log.Error().Msg("Create sitemap: " + err.Error())
		return
	}

	//jobs
	err = create_files(maxJobsPage, &index, r.job, host + "/jobs")
	if err != nil{
		log.Error().Msg("Create sitemap: " + err.Error())
		return
	}

	//resumes
	err = create_files(maxResumesPage, &index, r.resume, host + "/resumes")
	if err != nil{
		log.Error().Msg("Create sitemap: " + err.Error())
		return
	}

	//companies
	err = create_files(maxCompaniesPage, &index, r.company, host + "/companies")
	if err != nil{
		log.Error().Msg("Create sitemap: " + err.Error())
		return
	}

	//companies_with_jobs
	err = create_files(maxCompaniesWithJobs, &index, r.companywithjobs, host + "/company_jobs")
	if err != nil{
		log.Error().Msg("Create sitemap: " + err.Error())
		return
	}


}

func (r *Creator)max_pages_in_resource(maxJobsPage *int, maxResumesPage *int, maxCompaniesPage *int,
	maxCompaniesWithJobs *int) error {
	jobCount, err := r.job.Count()
	if err != nil {
		log.Error().Msg("Error Jobs.Count()")
		return  err
	}

	companyCount, err := r.company.Count()
	if err != nil {
		log.Error().Msg("Error Companies.Count()")
		return err
	}

	resumeCount, err := r.resume.Count()
	if err != nil {
		log.Error().Msg("Error Resumes.Count()")
		return err
	}

	companyWithJobsCount, err := r.companywithjobs.Count()
	if err != nil {
		log.Error().Msg("Error CompanyWithJobs.CompaniesWithJobs()")
		return err
	}

	*maxJobsPage = max_page(jobCount)
	*maxResumesPage = max_page(resumeCount)
	*maxCompaniesPage = max_page(companyCount)
	*maxCompaniesWithJobs = max_page(companyWithJobsCount)

	return nil
}


func max_page(count int) int{
	result := count / NumberOfItems
	if count % NumberOfItems > 0 {
		result +=1
	}
	return result
}

func create_files(count int, index *int, object models.ObjectRepository, url string) error {
	for i:=1; i<=count; i++{

		file, err := os.Create(Path + fmt.Sprintf("sitemaps%d.xml",*index))
		if err != nil {
			return err
		}

		file.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">")

		entities, err := object.ObjectsForSitemap(i, NumberOfItems)

		if err != nil {
			return err
		}
		for _, entity := range *entities {
			file.WriteString(*entity.XML(url))
		}

		file.WriteString("</urlset>")
		defer file.Close()
		*index +=1
	}
	return nil
}
