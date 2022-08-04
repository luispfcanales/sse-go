package domain

//Tour is template to tours
type Tour struct {
	ID          string
	CompanyName string
	Description string
	Photo       string
	Logo        string
	Pictures    []Picture
	Videos      []Video
	render      string
}

//Picture is template to Tours - Picture
type Picture struct {
	ID  string
	URL string
}

//Video is template to Tours - Video
type Video struct {
	ID  string
	URL string
}

//ChangeRender change state to render
func (t *Tour) ChangeRender(state string) {
	t.render = state
}

//GetStateRender get state
func (t Tour) GetStateRender() string {
	return t.render
}
