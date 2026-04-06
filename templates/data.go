package templates

func GetProjectBySlug(slug string) (Project, bool) {
	for _, p := range projects {
		if p.Slug == slug {
			return p, true
		}
	}
	return Project{}, false
}

var projects = []Project{
	{
		Slug:        "project-one",
		Title:       "Project One",
		Image:       "/static/images/project-one.jpg",
		Description: "Short description of project one.",
	},
	{
		Slug:        "project-two",
		Title:       "Project Two",
		Image:       "/static/images/project-two.jpg",
		Description: "Short description of project two.",
	},
}
