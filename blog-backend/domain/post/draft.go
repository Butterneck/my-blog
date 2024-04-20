package post

type Draft struct {
	title  Title
	body   Body
	slug   Slug
	assets []string
}

func newDraft(title, body string, assets []string) (*Draft, error) {
	var err error
	var d Draft

	d.title, err = newTitle(title)
	if err != nil {
		return nil, err
	}

	d.slug, err = newSlug(d.title)
	if err != nil {
		return nil, err
	}

	d.body, err = newBody(string(body))
	if err != nil {
		return nil, err
	}

	d.assets = assets

	return &d, nil
}

func (d Draft) updateTitle(title string) (*Draft, error) {
	var err error
	d.title, err = newTitle(title)
	if err != nil {
		return nil, err
	}

	d.slug, err = newSlug(d.title)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (d Draft) updateBody(body string) (*Draft, error) {
	var err error
	d.body, err = newBody(string(body))
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (d Draft) addAssets(assets []string) (*Draft, error) {
	for _, asset := range assets {
		d.assets = append(d.assets, asset)
	}
	return &d, nil
}

func (d Draft) removeAssets(assets []string) (*Draft, error) {
	for _, asset := range assets {
		for i, a := range d.assets {
			if a == asset {
				d.assets = append(d.assets[:i], d.assets[i+1:]...)
			}
		}
	}
	return &d, nil
}

func (d Draft) Title() Title {
	return d.title
}

func (d Draft) Body() Body {
	return d.body
}

func (d Draft) Slug() Slug {
	return d.slug
}

func (d Draft) Assets() []string {
	return d.assets
}
