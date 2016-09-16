package project

import (
	"fmt"
	"sort"

	"github.com/odlp/go-tracker"
)

//go:generate counterfeiter . ProjectClient
type ProjectClient interface {
	ProjectMemberships() ([]tracker.ProjectMembership, error)
	Stories(query tracker.StoriesQuery) ([]tracker.Story, tracker.Pagination, error)
}

type Project struct {
	Client ProjectClient
}

func NewProject(trackerAPIToken string, trackerProjectID int) Project {
	trackerClient := tracker.NewClient(trackerAPIToken)
	projectClient := trackerClient.InProject(trackerProjectID)
	return Project{Client: projectClient}
}

func (p Project) FindUserByEmail(email string) (tracker.ProjectMembership, error) {
	members, err := p.Client.ProjectMemberships()

	if err != nil {
		return tracker.ProjectMembership{}, err
	}

	for _, m := range members {
		if m.Person.Email == email {
			return m, nil
		}
	}

	return tracker.ProjectMembership{}, fmt.Errorf("Unable to find '%s' in project members", email)
}

func (p Project) FindCurrentStory(member tracker.ProjectMembership) (tracker.Story, error) {
	owner := fmt.Sprintf("owner:%s", member.Person.Initials)
	state := fmt.Sprintf("state:%s", tracker.StoryStateStarted)

	q := tracker.StoriesQuery{
		Filter: []string{owner, state},
	}

	results, _, err := p.Client.Stories(q)

	if err != nil {
		return tracker.Story{}, err
	}

	if len(results) == 0 {
		return tracker.Story{}, fmt.Errorf("Unable to find current story for user '%s'", member.Person.Email)
	}

	dateSortedStories := sortByUpdatedAt(results)
	return dateSortedStories[0], nil
}

func sortByUpdatedAt(unordered []tracker.Story) []tracker.Story {
	sorted := make(updatedAtSlice, 0, len(unordered))
	for _, s := range unordered {
		sorted = append(sorted, s)
	}

	sort.Sort(sorted)
	return sorted
}

type updatedAtSlice []tracker.Story

func (p updatedAtSlice) Len() int {
	return len(p)
}

func (p updatedAtSlice) Less(i, j int) bool {
	return p[i].UpdatedAt.After(*p[j].UpdatedAt)
}

func (p updatedAtSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
