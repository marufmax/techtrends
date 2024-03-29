package collector

import (
	"context"
	"fmt"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/marufmax/techtrends/crawler"
	"github.com/marufmax/techtrends/db/models"
)

func Collect() {
	ctx := context.Background()
	storeJobCount(ctx)
}

func storeJobCount(ctx context.Context) {
	categories, err := models.GetAllCategories(ctx)

	if err != nil {
		fmt.Println(err)
	}

	var jobCounts []models.JobCount
	for _, category := range categories {
		job, err := crawler.GetCount(category.Name)
		if err != nil {
			continue
		}

		jobCounts = append(jobCounts, models.JobCount{
			Count:        uint64(job.Count),
			CategoryID:   category.ID,
			Vacancy:      uint64(job.Vacancy),
			Category:     &category,
			CategoryType: category.Type,
			CategoryName: category.Name,
		})
	}

	_, err = models.InsertManyJobs(ctx, jobCounts)

	if err != nil {
		bugsnag.Notify(fmt.Errorf("could not insert job count to the DB: %s", err), ctx)
	}

}
