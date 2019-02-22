package kontent

import (
	"context"
	"regexp"
	"sync"
)

// SearchKeywords will go through all of the articles in return a new subset of articles
func (a *Articles) SearchKeywords(searchTerm string) (*Articles, error) {

	// Check that articles has been populated
	if a == nil {
		return nil, nil
	}

	// New subset of articles built using search terms
	var subset Articles

	// Create wait group for all go routines
	var wg sync.WaitGroup

	// Error channel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure it's called to release resources even if no errors

	errs := make(chan error, 4) // Buffer for 4 errors, one for each go routine

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Search all github articles
		for y := range a.GHRepos {
			// Iterate through all keyworkds
			for x := range a.GHRepos[y].Keywords {
				select {
				case <-ctx.Done():
					return // Error somewhere, terminate
				default: // Default is must to avoid blocking
				}
				// Apply the RegEx to the keyword
				matched, err := regexp.MatchString(searchTerm, a.GHRepos[y].Keywords[x])
				if err != nil {
					errs <- err
					cancel()
					return
				}
				// If the regex matches then add it to the new subset
				if matched == true {
					subset.GHRepos = append(subset.GHRepos, a.GHRepos[y])
					// Pop out the loop as multiple matches will create multiple entries
					break
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Search all snipped articles
		for y := range a.Snippets {
			// Iterate through all keyworkds
			for x := range a.Snippets[y].Keywords {
				select {
				case <-ctx.Done():
					return // Error somewhere, terminate
				default: // Default is must to avoid blocking
				}
				// Apply the RegEx to the keyword
				matched, err := regexp.MatchString(searchTerm, a.Snippets[y].Keywords[x])
				if err != nil {
					errs <- err
					cancel()
					return
				}
				// If the regex matches then add it to the new subset
				if matched == true {
					subset.Snippets = append(subset.Snippets, a.Snippets[y])
					// Pop out the loop as multiple matches will create multiple entries
					break
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Search all post articles
		for y := range a.Posts {
			// Iterate through all keyworkds
			for x := range a.Posts[y].Keywords {
				select {
				case <-ctx.Done():
					return // Error somewhere, terminate
				default: // Default is must to avoid blocking
				}
				// Apply the RegEx to the keyword
				matched, err := regexp.MatchString(searchTerm, a.Posts[y].Keywords[x])
				if err != nil {
					errs <- err
					cancel()
					return
				}
				// If the regex matches then add it to the new subset
				if matched == true {
					subset.Posts = append(subset.Posts, a.Posts[y])
					// Pop out the loop as multiple matches will create multiple entries
					break
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// Search all tools articles
		for y := range a.Tools {
			// Iterate through all keyworkds
			for x := range a.Tools[y].Keywords {
				select {
				case <-ctx.Done():
					return // Error somewhere, terminate
				default: // Default is must to avoid blocking
				}
				// Apply the RegEx to the keyword
				matched, err := regexp.MatchString(searchTerm, a.Tools[y].Keywords[x])
				if err != nil {
					errs <- err
					cancel()
					return
				}
				// If the regex matches then add it to the new subset
				if matched == true {
					subset.Tools = append(subset.Tools, a.Tools[y])
					// Pop out the loop as multiple matches will create multiple entries
					break
				}
			}
		}
	}()

	// Wait for all Go routines to be completed through the Done
	wg.Wait()

	if ctx.Err() != nil {
		return nil, <-errs
	}

	return &subset, nil
}
